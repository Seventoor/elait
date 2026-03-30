package main

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type Translator struct {
	client *openai.Client
	cfg    Config
}

func NewTranslator(cfg Config) *Translator {
	return &Translator{
		client: openai.NewClient(cfg.OpenAIAPIKey),
		cfg:    cfg,
	}
}

func (t *Translator) Translate(ctx context.Context, text string, targetLang string) (string, error) {
	resp, err := t.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Du übersetzt Predigttexte. Behalte biblische Sprachformen bei. Gib nur die Übersetzung zurück, nichts anderes.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Übersetze ins %s: %s", targetLang, text),
			},
		},
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
