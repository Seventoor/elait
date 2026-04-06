package main

import (
	"context"
	"fmt"

	"github.com/Turbocommerce/clait/config"
	"github.com/sashabaranov/go-openai"
)

type Translator struct {
	client *openai.Client
	cfg    config.Config
}

func NewTranslator(cfg config.Config) *Translator {
	return &Translator{
		client: openai.NewClient(cfg.OpenAIAPIKey),
		cfg:    cfg,
	}
}

func (t *Translator) Translate(ctx context.Context, text string, targetLang string) (string, error) {
	resp, err := t.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:               openai.GPT4oMini,
		Temperature:         0,
		MaxCompletionTokens: 200,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a translation engine. Output ONLY the translated text. No explanations, no notes, no quotes, no formatting. Just the raw translation.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("Translate the following sermon text to %s:\n\n%s", targetLang, text),
			},
		},
	})
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
