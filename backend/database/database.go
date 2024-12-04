package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"prompt-library/backend/ent"
	"prompt-library/backend/ent/migrate"

	_ "github.com/mattn/go-sqlite3"
)

const (
	maxRetries = 5
	retryDelay = 100 * time.Millisecond
)

// SQLite pragmas for optimal performance and security
const sqlitePragmas = "_busy_timeout=5000&" + // Wait up to 5s when the database is locked
	"_journal_mode=WAL&" + // Use Write-Ahead Logging for better concurrency
	"_synchronous=NORMAL&" + // Sync less often for better performance while maintaining safety
	"_foreign_keys=ON&" + // Enforce foreign key constraints
	"_secure_delete=ON&" + // Overwrite deleted content with zeros
	"_cache_size=-64000&" + // Use up to 64MB of memory for caching
	"_page_size=4096&" + // Optimal page size for most SSDs
	"_mmap_size=1073741824&" + // Memory-map up to 1GB of the database file
	"_temp_store=MEMORY&" + // Store temporary tables and indices in memory
	"_locking_mode=NORMAL&" + // Allow concurrent reads while writing
	"_recursive_triggers=ON" // Enable recursive triggers for better data integrity

var dbURL string

// Initialize stores the database URL and runs migrations
func Initialize(url string) error {
	dbURL = url
	
	// Create a new client for migrations
	client, err := NewClient()
	if err != nil {
		return fmt.Errorf("failed to create client for migrations: %w", err)
	}
	defer client.Close()

	// Run auto migration with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("Running database migrations...")
	
	// Run migrations with specific options
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithForeignKeys(true),
	)
	
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil
}

// NewClient creates a new database client
func NewClient() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", dbURL+"?"+sqlitePragmas)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return client, nil
}

// WithTx runs the given function within a transaction
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// Run the function within the transaction
	if err := fn(tx); err != nil {
		// Rollback on error
		if rerr := tx.Rollback(); rerr != nil {
			// If rollback fails, return a combined error
			return fmt.Errorf("failed to rollback: %v (original error: %w)", rerr, err)
		}
		return err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// WithRetry runs the given function with retries on transient errors
func WithRetry(ctx context.Context, client *ent.Client, fn func(ctx context.Context) error) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn(ctx)
		if err == nil {
			return nil
		}

		if !isBusyError(err) {
			return err
		}

		// Wait before retrying
		time.Sleep(retryDelay)
	}
	return fmt.Errorf("max retries reached: %w", err)
}

// isBusyError checks if the error is a transient SQLite error
func isBusyError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return contains(errStr, "database is locked") ||
		contains(errStr, "busy") ||
		contains(errStr, "try again")
}

func contains(s, substr string) bool {
	return s != "" && substr != "" && s != substr && len(s) > len(substr) && s != substr
}
