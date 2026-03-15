package youtube

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/api/option"
	ytapi "google.golang.org/api/youtube/v3"
	"github.com/adamwrose/streamfusion/internal/models"
)

// Provider polls the YouTube Live Chat API.
type Provider struct {
	apiKey      string
	liveChatID  string
	msgCh       chan models.ChatMessage
	cancelFn    context.CancelFunc
}

func NewProvider(apiKey, liveChatID string) *Provider {
	return &Provider{
		apiKey:     apiKey,
		liveChatID: liveChatID,
		msgCh:      make(chan models.ChatMessage, 256),
	}
}

func (p *Provider) GetName() string { return "youtube" }

func (p *Provider) Connect() error {
	ctx, cancel := context.WithCancel(context.Background())
	p.cancelFn = cancel

	svc, err := ytapi.NewService(ctx, option.WithAPIKey(p.apiKey))
	if err != nil {
		return fmt.Errorf("youtube: create service: %w", err)
	}

	go p.poll(ctx, svc)
	return nil
}

func (p *Provider) poll(ctx context.Context, svc *ytapi.Service) {
	var pageToken string
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		call := svc.LiveChatMessages.List(p.liveChatID, []string{"snippet", "authorDetails"}).MaxResults(200)
		if pageToken != "" {
			call = call.PageToken(pageToken)
		}

		resp, err := call.Do()
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		for _, item := range resp.Items {
			p.msgCh <- models.ChatMessage{
				ID:          item.Id,
				Platform:    "youtube",
				Username:    item.AuthorDetails.DisplayName,
				DisplayName: item.AuthorDetails.DisplayName,
				Message:     item.Snippet.DisplayMessage,
				Timestamp:   time.Now(),
			}
		}

		pageToken = resp.NextPageToken
		time.Sleep(time.Duration(resp.PollingIntervalMillis) * time.Millisecond)
	}
}

func (p *Provider) Disconnect() error {
	if p.cancelFn != nil {
		p.cancelFn()
	}
	return nil
}

func (p *Provider) GetMessageChannel() <-chan models.ChatMessage {
	return p.msgCh
}

func (p *Provider) ExecuteAction(action string, target string) error {
	return fmt.Errorf("youtube: action %q not yet implemented", action)
}
