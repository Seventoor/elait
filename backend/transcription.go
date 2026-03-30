package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	api "github.com/deepgram/deepgram-go-sdk/v3/pkg/api/listen/v1/websocket/interfaces"
)

type MyCallback struct {
	cfg        Config
	translator *Translator
	tts        *SpeechEngine
	hub        *Hub
}

func (c MyCallback) Message(mr *api.MessageResponse) error {
	if len(mr.Channel.Alternatives) == 0 {
		return nil
	}

	transcript := strings.TrimSpace(mr.Channel.Alternatives[0].Transcript)

	if len(transcript) > 0 {

		if mr.IsFinal {
			activeLangs := c.hub.GetActiveLanguages()

			for _, lang := range activeLangs {
				targetLang := lang

				go func() {
					ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
					defer cancel()

					translated, err := c.translator.Translate(ctx, transcript, targetLang)
					if err != nil {
						return
					}
					fmt.Printf("[Final] %s \n", translated)

					audio, _ := c.tts.TextToSpeech(ctx, translated, targetLang)
					if err == nil {
						c.hub.Broadcast(audio, targetLang)
					}
				}()
			}
		} else {
			fmt.Printf("[Interim] %s \n", transcript)
		}

	}
	return nil
}

func (c MyCallback) Open(*api.OpenResponse) error {
	fmt.Printf("Stream started...\n")
	return nil
}

func (c MyCallback) Close(*api.CloseResponse) error {
	fmt.Printf("Stream stopped...\n")
	return nil
}

func (c MyCallback) Metadata(*api.MetadataResponse) error           { return nil }
func (c MyCallback) Error(*api.ErrorResponse) error                 { return nil }
func (c MyCallback) UnhandledEvent([]byte) error                    { return nil }
func (c MyCallback) SpeechStarted(*api.SpeechStartedResponse) error { return nil }
func (c MyCallback) UtteranceEnd(*api.UtteranceEndResponse) error   { return nil }
