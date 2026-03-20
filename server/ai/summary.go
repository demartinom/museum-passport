package ai

import (
	"context"
	"fmt"

	"github.com/demartinom/museum-passport/models"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type SummaryClient struct {
	client openai.Client
}

func NewSummaryClient(APIKey string) *SummaryClient {
	return &SummaryClient{
		client: openai.NewClient(option.WithAPIKey(APIKey)),
	}
}

func (s *SummaryClient) GenerateSummary(artwork models.SingleArtwork) (string, error) {
	prompt := fmt.Sprintf(
		"Provide a brief, educational summary (2-3 paragraphs) about this artwork:\n\n"+
			"Title: %s\n"+
			"Artist: %s\n"+
			"Date: %s\n"+
			"Medium: %s\n"+
			"Museum: %s\n\n"+
			"Include historical context, artistic significance, and interesting details about the work or artist. Do not include name of museum where housed. If double sided painting, only describe current side.",
		artwork.ArtworkTitle,
		artwork.ArtistName,
		artwork.DateCreated,
		artwork.ArtMedium,
		artwork.Museum,
	)

	params := openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4oMini,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
	}
	params.MaxTokens = openai.Int(300)

	completion, err := s.client.Chat.Completions.New(context.Background(), params)
	if err != nil {
		return "", err
	}

	return completion.Choices[0].Message.Content, nil
}
