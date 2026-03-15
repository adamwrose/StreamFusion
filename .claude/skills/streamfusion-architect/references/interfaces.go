package providers

// Provider is the contract for all streaming platforms (Twitch, YT, Kick)
type Provider interface {
	GetName() string
	Connect() error
	Disconnect() error
	GetMessageChannel() <-chan models.ChatMessage
	ExecuteAction(action string, target string) error // For bans/timeouts
}
