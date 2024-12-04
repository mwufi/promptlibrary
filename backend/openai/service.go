package openai

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Service struct {
	client *openai.Client
}

func NewService(apiKey string) *Service {
	client := openai.NewClient(apiKey)
	return &Service{
		client: client,
	}
}

func (s *Service) RunPrompt(ctx context.Context, promptContent string, userInput string) (string, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: promptContent,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		},
	}

	resp, err := s.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Messages:    messages,
			MaxTokens:   1000,
			Temperature: 0.7,
		},
	)

	if err != nil {
		return "", fmt.Errorf("failed to create chat completion: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response choices returned from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}
