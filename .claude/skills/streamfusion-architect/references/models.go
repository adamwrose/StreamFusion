package models

import "time"

type ChatMessage struct {
	ID          string    `json:"id"`
	Platform    string    `json:"platform"`
	Username    string    `json:"username"`
	DisplayName string    `json:"display_name"`
	Message     string    `json:"message"`
	Color       string    `json:"color"`
	Badges      []string  `json:"badges"`
	IsMod       bool      `json:"is_mod"`
	Timestamp   time.Time `json:"timestamp"`
}
