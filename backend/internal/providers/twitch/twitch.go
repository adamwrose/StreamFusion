package twitch

import (
	"fmt"

	twitchirc "github.com/gempir/go-twitch-irc/v4"
	"github.com/adamwrose/streamfusion/internal/models"
)

// Provider connects to Twitch chat via IRC.
type Provider struct {
	channel string
	client  *twitchirc.Client
	msgCh   chan models.ChatMessage
}

func NewProvider(username, oauth, channel string) *Provider {
	return &Provider{
		channel: channel,
		client:  twitchirc.NewClient(username, oauth),
		msgCh:   make(chan models.ChatMessage, 256),
	}
}

func (p *Provider) GetName() string { return "twitch" }

func (p *Provider) Connect() error {
	p.client.OnPrivateMessage(func(msg twitchirc.PrivateMessage) {
		p.msgCh <- models.ChatMessage{
			ID:          msg.ID,
			Platform:    "twitch",
			Username:    msg.User.Name,
			DisplayName: msg.User.DisplayName,
			Message:     msg.Message,
			Color:       msg.User.Color,
			IsMod:       msg.User.Badges["moderator"] == 1,
			Timestamp:   msg.Time,
		}
	})
	p.client.Join(p.channel)
	return p.client.Connect()
}

func (p *Provider) Disconnect() error {
	return p.client.Disconnect()
}

func (p *Provider) GetMessageChannel() <-chan models.ChatMessage {
	return p.msgCh
}

func (p *Provider) ExecuteAction(action string, target string) error {
	switch action {
	case "ban":
		p.client.Say(p.channel, fmt.Sprintf("/ban %s", target))
	case "timeout":
		p.client.Say(p.channel, fmt.Sprintf("/timeout %s 600", target))
	default:
		return fmt.Errorf("unknown action: %s", action)
	}
	return nil
}
