package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"prompt-library/backend/database"
	"prompt-library/backend/ent"
	"prompt-library/backend/ent/prompt"
	"prompt-library/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleGetPrompts(c *fiber.Ctx) error {
    start := time.Now()
    log.Printf("[GET /api/prompts] Starting request")
    
    // Create a new database client for this request
    client, err := database.NewClient()
    if err != nil {
        wrappedErr := utils.WrapError(err, "failed to create database client")
        log.Printf("[GET /api/prompts] %+v", wrappedErr)
        return c.Status(500).JSON(ErrorResponse{Error: "Database connection error"})
    }
    defer client.Close()
    
    prompts, err := client.Prompt.
        Query().
        Order(ent.Desc(prompt.FieldCreatedAt)).
        All(context.Background())
    
    if err != nil {
        wrappedErr := utils.WrapError(err, "failed to fetch prompts")
        log.Printf("[GET /api/prompts] %+v", wrappedErr)
        return c.Status(500).JSON(ErrorResponse{Error: "Failed to fetch prompts"})
    }

    log.Printf("[GET /api/prompts] Successfully fetched %d prompts", len(prompts))
    var response []PromptResponse
    for _, p := range prompts {
        log.Printf("[GET /api/prompts] Processing prompt: ID=%d, Title='%s', Category='%s'", 
            p.ID, p.Title, p.Category)
        response = append(response, PromptResponse{
            ID:        p.ID,
            Title:     p.Title,
            Content:   p.Content,
            Category:  p.Category,
            Votes:     p.Votes,
            CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
            CreatedBy: p.CreatedBy,
        })
    }

    duration := time.Since(start)
    log.Printf("[GET /api/prompts] Request completed in %v", duration)
    
    return c.JSON(response)
}

func HandleCreatePrompt(c *fiber.Ctx) error {
    start := time.Now()
    log.Printf("[POST /api/prompts] Starting request")

    var req CreatePromptRequest
    if err := c.BodyParser(&req); err != nil {
        log.Printf("[POST /api/prompts] Invalid request body: %v", err)
        return c.Status(400).JSON(ErrorResponse{Error: "Invalid request body"})
    }

    // Create a new database client for this request
    client, err := database.NewClient()
    if err != nil {
        wrappedErr := utils.WrapError(err, "failed to create database client")
        log.Printf("[POST /api/prompts] %+v", wrappedErr)
        return c.Status(500).JSON(ErrorResponse{Error: "Database connection error"})
    }
    defer client.Close()

    var result *ent.Prompt
    err = database.WithRetry(c.Context(), client, func(ctx context.Context) error {
        return database.WithTx(ctx, client, func(tx *ent.Tx) error {
            p, err := tx.Prompt.
                Create().
                SetTitle(req.Title).
                SetContent(req.Content).
                SetCategory(req.Category).
                SetCreatedBy("anonymous"). // TODO: Replace with actual user
                Save(ctx)
            
            if err != nil {
                return fmt.Errorf("failed to create prompt: %w", err)
            }
            result = p
            return nil
        })
    })

    if err != nil {
        wrappedErr := utils.WrapError(err, "failed to create prompt")
        log.Printf("[POST /api/prompts] %+v", wrappedErr)
        return c.Status(500).JSON(ErrorResponse{Error: "Failed to create prompt"})
    }

    duration := time.Since(start)
    log.Printf("[POST /api/prompts] Created prompt ID=%d in %v", result.ID, duration)

    return c.Status(201).JSON(PromptResponse{
        ID:        result.ID,
        Title:     result.Title,
        Content:   result.Content,
        Category:  result.Category,
        Votes:     result.Votes,
        CreatedAt: result.CreatedAt.Format("2006-01-02 15:04:05"),
        CreatedBy: result.CreatedBy,
    })
}

func HandleGetPrompt(c *fiber.Ctx) error {
    start := time.Now()
    promptID := c.Params("id")
    log.Printf("[GET /api/prompts/%s] Starting request", promptID)

    id, err := strconv.Atoi(promptID)
    if err != nil {
        log.Printf("[GET /api/prompts/%s] Invalid ID format: %v", promptID, err)
        return c.Status(400).JSON(ErrorResponse{Error: fmt.Sprintf("Invalid prompt ID: %s", promptID)})
    }

    // Create a new database client for this request
    client, err := database.NewClient()
    if err != nil {
        wrappedErr := utils.WrapError(err, "failed to create database client")
        log.Printf("[GET /api/prompts/%d] %+v", id, wrappedErr)
        return c.Status(500).JSON(ErrorResponse{Error: "Database connection error"})
    }
    defer client.Close()

    p, err := client.Prompt.
        Query().
        Where(prompt.IDEQ(id)).
        Only(c.Context())

    if err != nil {
        if ent.IsNotFound(err) {
            log.Printf("[GET /api/prompts/%d] Prompt not found", id)
            return c.Status(404).JSON(ErrorResponse{Error: "Prompt not found"})
        }
        wrappedErr := utils.WrapError(err, fmt.Sprintf("failed to fetch prompt %d", id))
        log.Printf("[GET /api/prompts/%d] %+v", id, wrappedErr)
        return c.Status(500).JSON(ErrorResponse{Error: "Failed to fetch prompt"})
    }

    duration := time.Since(start)
    log.Printf("[GET /api/prompts/%d] Request completed in %v", id, duration)

    return c.JSON(PromptResponse{
        ID:        p.ID,
        Title:     p.Title,
        Content:   p.Content,
        Category:  p.Category,
        Votes:     p.Votes,
        CreatedAt: p.CreatedAt.Format("2006-01-02 15:04:05"),
        CreatedBy: p.CreatedBy,
    })
}

func HandleVotePrompt(c *fiber.Ctx) error {
    start := time.Now()
    promptID := c.Params("id")
    log.Printf("[POST /api/prompts/%s/vote] Starting request", promptID)

    id, err := strconv.Atoi(promptID)
    if err != nil {
        log.Printf("[POST /api/prompts/%s/vote] Invalid ID format: %v", promptID, err)
        return c.Status(400).JSON(ErrorResponse{Error: fmt.Sprintf("Invalid prompt ID: %s", promptID)})
    }

    // Create a new database client for this request
    client, err := database.NewClient()
    if err != nil {
        wrappedErr := utils.WrapError(err, "failed to create database client")
        log.Printf("[POST /api/prompts/%d/vote] %+v", id, wrappedErr)
        return c.Status(500).JSON(ErrorResponse{Error: "Database connection error"})
    }
    defer client.Close()

    var result *ent.Prompt
    err = database.WithRetry(c.Context(), client, func(ctx context.Context) error {
        return database.WithTx(ctx, client, func(tx *ent.Tx) error {
            p, err := tx.Prompt.
                UpdateOneID(id).
                AddVotes(1).
                Save(ctx)
            
            if err != nil {
                if ent.IsNotFound(err) {
                    return fiber.NewError(fiber.StatusNotFound, "Prompt not found")
                }
                return fmt.Errorf("failed to update prompt: %w", err)
            }
            result = p
            return nil
        })
    })

    if err != nil {
        if ferr, ok := err.(*fiber.Error); ok {
            return c.Status(ferr.Code).JSON(ErrorResponse{Error: ferr.Message})
        }
        wrappedErr := utils.WrapError(err, fmt.Sprintf("failed to vote for prompt %d", id))
        log.Printf("[POST /api/prompts/%d/vote] %+v", id, wrappedErr)
        return c.Status(500).JSON(ErrorResponse{Error: "Failed to vote for prompt"})
    }

    duration := time.Since(start)
    log.Printf("[POST /api/prompts/%d/vote] Request completed in %v", id, duration)

    return c.JSON(PromptResponse{
        ID:        result.ID,
        Title:     result.Title,
        Content:   result.Content,
        Category:  result.Category,
        Votes:     result.Votes,
        CreatedAt: result.CreatedAt.Format("2006-01-02 15:04:05"),
        CreatedBy: result.CreatedBy,
    })
}
