package main

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	clients map[*websocket.Conn]string
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients: make(map[*websocket.Conn]string),
	}
}

func (h *Hub) AddClient(conn *websocket.Conn, lang string) {
	h.mu.Lock()
	h.clients[conn] = lang
	h.mu.Unlock()
}

func (h *Hub) RemoveClient(conn *websocket.Conn) {
	h.mu.Lock()
	delete(h.clients, conn)
	h.mu.Unlock()
}

func (h *Hub) Broadcast(audio []byte, lang string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for conn, clientLang := range h.clients {
		if clientLang == lang {
			err := conn.WriteMessage(websocket.BinaryMessage, audio)
			if err != nil {
				conn.Close()
				delete(h.clients, conn)
			}
		}
	}
}

func (h *Hub) GetActiveLanguages() []string {
	h.mu.Lock()
	defer h.mu.Unlock()

	uniqueLangs := make(map[string]bool)
	var langs []string

	for _, lang := range h.clients {
		if !uniqueLangs[lang] {
			uniqueLangs[lang] = true
			langs = append(langs, lang)
		}
	}
	return langs
}
