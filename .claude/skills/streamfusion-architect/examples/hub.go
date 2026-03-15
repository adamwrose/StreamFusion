package hub

import (
	"sync"
)

type Hub struct {
	clients    map[chan models.ChatMessage]bool
	Broadcast  chan models.ChatMessage
	register   chan chan models.ChatMessage
	unregister chan chan models.ChatMessage
	mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[chan models.ChatMessage]bool),
		Broadcast:  make(chan models.ChatMessage),
		register:   make(chan chan models.ChatMessage),
		unregister: make(chan chan models.ChatMessage),
	}
}

// Run starts the main loop for message distribution
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
		case msg := <-h.Broadcast:
			h.mu.Lock()
			for client := range h.clients {
				client <- msg
			}
			h.mu.Unlock()
		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client)
			}
			h.mu.Unlock()
		}
	}
}
