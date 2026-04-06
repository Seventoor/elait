package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Turbocommerce/clait/config"
	"github.com/deepgram/deepgram-go-sdk/v3/pkg/client/interfaces"
	client "github.com/deepgram/deepgram-go-sdk/v3/pkg/client/listen"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func audioIncomeHandler(w http.ResponseWriter, r *http.Request, deepgramKey string, cfg config.Config, tr *Translator, tts *SpeechEngine, hub *Hub) {
	ctx := context.Background()
	options := &interfaces.LiveTranscriptionOptions{
		Model:          "nova-3",
		Language:       cfg.SourceLanguage,
		InterimResults: true,
	}

	callback := MyCallback{
		cfg:        cfg,
		translator: tr,
		tts:        tts,
		hub:        hub,
	}

	dgClient, err := client.NewWSUsingCallback(ctx, deepgramKey, &interfaces.ClientOptions{}, options, callback)
	if err != nil {
		panic(err)
	}
	if !dgClient.Connect() {
		panic("failed to connect")
	}
	defer dgClient.Stop()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	for {
		_, audioBytes, err := conn.ReadMessage()
		if err != nil {
			break
		}
		dgClient.Write(audioBytes)
	}
}

func audioOutcomeHandler(w http.ResponseWriter, r *http.Request, hub *Hub) {
	lang := r.URL.Query().Get("lang")
	if lang == "" {
		lang = "en-US"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	hub.AddClient(conn, lang)
	defer func() {
		hub.RemoveClient(conn)
		conn.Close()
	}()

	// offen halten bis Client disconnect
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func languageListHandler(w http.ResponseWriter, r *http.Request) {
	languages := config.Languages

	w.Header().Set("Content-Type", "application/json")

	// Encode the array as JSON and write it to the response
	if err := json.NewEncoder(w).Encode(languages); err != nil {
		http.Error(w, "Failed to encode languages", http.StatusInternalServerError)
		return
	}
}
