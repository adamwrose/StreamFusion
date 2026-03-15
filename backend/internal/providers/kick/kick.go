package kick

import (
	"fmt"

	"github.com/adamwrose/streamfusion/internal/models"
)

// Provider connects to Kick chat. Kick has no official Go SDK;
// this stub uses a WebSocket connection to their Pusher-based API.
type Provider struct {
	channel string
	msgCh   chan models.ChatMessage
}

func NewProvider(channel string) *Provider {
	return &Provider{
		channel: channel,
		msgCh:   make(chan models.ChatMessage, 256),
	}
}

func (p *Provider) GetName() string { return "kick" }

func (p *Provider) Connect() error {
	// TODO: implement Pusher WebSocket connection to Kick chat.
	return fmt.Errorf("kick: provider not yet implemented")
}

func (p *Provider) Disconnect() error { return nil }

func (p *Provider) GetMessageChannel() <-chan models.ChatMessage {
	return p.msgCh
}

func (p *Provider) ExecuteAction(action string, target string) error {
	return fmt.Errorf("kick: action %q not yet implemented", action)
}
