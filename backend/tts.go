package main

import (
	"context"
	"fmt"
	"strings"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/Turbocommerce/clait/config"
)

type SpeechEngine struct {
	client *texttospeech.Client
	cfg    config.Config
}

func NewSpeechEngine(cfg config.Config) (*SpeechEngine, error) {
	ctx := context.Background()
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("Google TTS Client konnte nicht erstellt werden: %v", err)
	}

	return &SpeechEngine{
		client: client,
		cfg:    cfg,
	}, nil
}

func (s *SpeechEngine) TextToSpeech(ctx context.Context, text string, lang string) ([]byte, error) {
	currentVoice := s.getBestVoice(ctx, s.client, lang)

	req := &texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: lang,
			Name:         currentVoice,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding:   texttospeechpb.AudioEncoding_MP3,
			SampleRateHertz: 48000,
		},
	}

	resp, err := s.client.SynthesizeSpeech(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.AudioContent, nil
}

func (s *SpeechEngine) getBestVoice(ctx context.Context, client *texttospeech.Client, languageCode string) string {
	resp, err := client.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{
		LanguageCode: languageCode,
	})
	if err != nil {
		return ""
	}

	priority := []string{"Chirp3-HD-Orus", "Neural2", "Polyglot", "Wavenet", "Standard"}

	bestVoice := ""
	bestPriorityIndex := len(priority)

	for _, voice := range resp.Voices {
		// Only male voices for now, later as a setting in admin panel
		if voice.SsmlGender != texttospeechpb.SsmlVoiceGender_MALE {
			continue
		}

		for i, tier := range priority {
			if strings.Contains(voice.Name, tier) {
				if i < bestPriorityIndex {
					bestPriorityIndex = i
					bestVoice = voice.Name
				}
				break
			}
		}
	}

	return bestVoice
}
