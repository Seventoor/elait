package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	cfg := loadConfig()

	myTranslator := NewTranslator(cfg)

	myTTS, err := NewSpeechEngine(cfg)
	if err != nil {
		log.Fatal(err)
	}

	hub := &Hub{clients: make(map[*websocket.Conn]string)}

	router := http.NewServeMux()

	// Sender Stream
	router.HandleFunc("GET /ws", func(w http.ResponseWriter, r *http.Request) {
		audioIncomeHandler(w, r, cfg.DeepgramAPIKey, cfg, myTranslator, myTTS, hub)
	})

	// Receiver Stream
	router.HandleFunc("GET /audio-ws", func(w http.ResponseWriter, r *http.Request) {
		audioOutcomeHandler(w, r, hub)
	})

	fmt.Println("Server läuft auf http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
