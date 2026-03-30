package main

import (
	"context"
	"fmt"
	"strings"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

type SpeechEngine struct {
	client *texttospeech.Client
	cfg    Config
}

var BestVoices = map[string]string{
	"de-DE": "de-DE-Neural2-B",
	"de-AT": "de-AT-Neural2-A",
	"en-US": "en-US-Chirp3-HD-Orus",
	"en-GB": "en-GB-Neural2-B",
	"ru-RU": "ru-RU-Wavenet-D",
}

func (s *SpeechEngine) GetVoiceForLang(ctx context.Context, lang string) string {
	if voice, ok := BestVoices[lang]; ok {
		return voice
	}

	return getBestVoice(ctx, s.client, lang)
}

func NewSpeechEngine(cfg Config) (*SpeechEngine, error) {
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
	currentVoice := s.GetVoiceForLang(ctx, lang)

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
			SampleRateHertz: 24000,
		},
	}

	resp, err := s.client.SynthesizeSpeech(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.AudioContent, nil
}

func getBestVoice(ctx context.Context, client *texttospeech.Client, languageCode string) string {
	resp, err := client.ListVoices(ctx, &texttospeechpb.ListVoicesRequest{
		LanguageCode: languageCode,
	})
	if err != nil {
		return ""
	}

	// Neural2 > Wavenet > Standard — beste Qualität zuerst
	priority := []string{"Neural2", "Polyglot", "Wavenet", "Standard"}

	for _, tier := range priority {
		for _, voice := range resp.Voices {
			if strings.Contains(voice.Name, tier) {
				return "ru-RU-Wavenet-D"
			}
		}
	}
	return ""
}
