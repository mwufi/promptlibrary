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
	"prompt-library/backend/openai"
	"prompt-library/backend/utils"

	"github.com/gofiber/fiber/v2"
)

func HandleRunPrompt(c *fiber.Ctx) error {
	start := time.Now()
	promptID := c.Params("id")
	log.Printf("[POST /api/prompts/%s/run] Starting request", promptID)

	id, err := strconv.Atoi(promptID)
	if err != nil {
		log.Printf("[POST /api/prompts/%s/run] Invalid ID format: %v", promptID, err)
		return c.Status(400).JSON(ErrorResponse{Error: fmt.Sprintf("Invalid prompt ID: %s", promptID)})
	}

	var req RunPromptRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("[POST /api/prompts/%d/run] Invalid request body: %v", id, err)
		return c.Status(400).JSON(ErrorResponse{Error: "Invalid request body"})
	}

	// Create a new database client for this request
	client, err := database.NewClient()
	if err != nil {
		wrappedErr := utils.WrapError(err, "failed to create database client")
		log.Printf("[POST /api/prompts/%d/run] %+v", id, wrappedErr)
		return c.Status(500).JSON(ErrorResponse{Error: "Database connection error"})
	}
	defer client.Close()

	openaiService := c.Locals("openai").(*openai.Service)
	
	var result struct {
		prompt *ent.Prompt
		conversation *ent.Conversation
		response string
	}

	err = database.WithRetry(c.Context(), client, func(ctx context.Context) error {
		return database.WithTx(ctx, client, func(tx *ent.Tx) error {
			// Fetch prompt within transaction
			p, err := tx.Prompt.
				Query().
				Where(prompt.IDEQ(id)).
				Only(ctx)
			
			if err != nil {
				if ent.IsNotFound(err) {
					return fiber.NewError(fiber.StatusNotFound, "Prompt not found")
				}
				return err
			}
			result.prompt = p

			// Run OpenAI prompt
			resp, err := openaiService.RunPrompt(ctx, p.Content, req.Input)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "Failed to run prompt with OpenAI")
			}
			result.response = resp

			// Create conversation record within same transaction
			conv, err := tx.Conversation.
				Create().
				SetMessages(req.Input + "\n\n" + resp).
				SetPrompt(p).
				SetUserID("anonymous"). // TODO: Replace with actual user ID
				Save(ctx)
			
			if err != nil {
				return err
			}
			result.conversation = conv
			
			return nil
		})
	})

	if err != nil {
		if ferr, ok := err.(*fiber.Error); ok {
			return c.Status(ferr.Code).JSON(ErrorResponse{Error: ferr.Message})
		}
		wrappedErr := utils.WrapError(err, fmt.Sprintf("Failed to run prompt %d", id))
		log.Printf("[POST /api/prompts/%d/run] Error: %+v", id, wrappedErr)
		return c.Status(fiber.StatusInternalServerError).JSON(ErrorResponse{Error: "Internal server error"})
	}

	duration := time.Since(start)
	log.Printf("[POST /api/prompts/%d/run] Request completed in %v", id, duration)

	return c.JSON(RunPromptResponse{
		ConversationID: result.conversation.ID,
		Response:       result.response,
	})
}
