package providers

import "github.com/adamwrose/streamfusion/internal/models"

// Provider is the contract all platform integrations must satisfy.
type Provider interface {
	GetName() string
	Connect() error
	Disconnect() error
	GetMessageChannel() <-chan models.ChatMessage
	ExecuteAction(action string, target string) error
}
