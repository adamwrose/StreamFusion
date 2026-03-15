package hub

import (
	"sync"

	"github.com/adamwrose/streamfusion/internal/models"
)

// Hub broadcasts ChatMessages to all registered WebSocket clients.
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
		Broadcast:  make(chan models.ChatMessage, 256),
		register:   make(chan chan models.ChatMessage),
		unregister: make(chan chan models.ChatMessage),
	}
}

// Register adds a client channel to the broadcast pool.
func (h *Hub) Register(ch chan models.ChatMessage) {
	h.register <- ch
}

// Unregister removes a client channel from the broadcast pool.
func (h *Hub) Unregister(ch chan models.ChatMessage) {
	h.unregister <- ch
}

// Run is the main select loop; must be started in a goroutine.
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
				select {
				case client <- msg:
				default:
					// Drop message if client buffer is full.
				}
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
