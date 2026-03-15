package websocket

import (
	"encoding/json"
	"log"
	"net/http"

	gorillaws "github.com/gorilla/websocket"
	"github.com/adamwrose/streamfusion/internal/hub"
	"github.com/adamwrose/streamfusion/internal/models"
)

var upgrader = gorillaws.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWS(h *hub.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("websocket upgrade: %v", err)
		return
	}
	defer conn.Close()

	clientCh := make(chan models.ChatMessage, 64)
	h.Register(clientCh)
	defer h.Unregister(clientCh)

	for msg := range clientCh {
		data, err := json.Marshal(msg)
		if err != nil {
			continue
		}
		if err := conn.WriteMessage(gorillaws.TextMessage, data); err != nil {
			return
		}
	}
}

// HandleDashboard serves the moderator WebSocket endpoint.
func HandleDashboard(h *hub.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serveWS(h, w, r)
	}
}

// HandleOverlay serves the OBS browser-source WebSocket endpoint.
func HandleOverlay(h *hub.Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		serveWS(h, w, r)
	}
}
