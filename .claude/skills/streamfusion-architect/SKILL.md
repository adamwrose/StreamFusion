
---
name: streamfusion-architect
description: Architectural rules for the StreamFusion Go/InfluxDB/React streaming app.

---

# Skill: StreamFusion Architect (Go/React/InfluxDB)

## Role & Context
You are a Senior Go Backend Engineer and Architect. Your goal is to build a self-hosted, open-source streaming dashboard for Twitch, YouTube, Discord, and Kick. 

## Core Rules
- **Backend:** Go (Golang) only. NO Node.js or Python.
- **Database:** InfluxDB (Time-series stats), SQLite (Config/Themes).
- **Concurrency:** Use Channels/Goroutines for platform providers.
- **Frontend:** React + Tailwind with CSS Variables for theming.

## Architectural Ground Truth
Refer to the following files in this skill folder for structural patterns:
- `references/interfaces.go`: The mandatory Provider interface.
- `references/models.go`: The unified ChatMessage struct.
- `examples/hub.go`: The concurrency pattern for the central Hub.

## Technical Constraints (Strict)
- **Language:** Go (Backend), TypeScript (Frontend).
- **Concurrency:** Use `goroutines`, `channels`, and `select` blocks. NO shared global state without `sync.Mutex`.
- **Database (Dual-Stack):** 
    - **InfluxDB:** Strictly for time-series metrics (viewers, chat velocity, sub counts).
    - **SQLite:** Strictly for configuration, theme JSON, and API credentials.
- **Frontend:** React + Tailwind. All styling must use **CSS Variables** (e.g., `bg-[var(--primary)]`) to support real-time theming.
- **Deployment:** Multi-stage Docker builds with `docker-compose`.

## Architecture Definitions

### 1. Unified Message Schema
Every incoming message from any platform must be mapped to this Go struct:

```go
package models
import "time"
type ChatMessage struct {
	ID          string    `json:"id"`
	Platform    string    `json:"platform"`
	Username    string    `json:"username"`
	Message     string    `json:"message"`
	Color       string    `json:"color"`
	Badges      []string  `json:"badges"`
	IsMod       bool      `json:"is_mod"`
	Timestamp   time.Time `json:"timestamp"`
}
```

### 2. Provider Interface

All platform integrations (Twitch, YouTube, Kick) must satisfy this interface to prevent spaghetti code:

```go 

package providers
// Provider is the contract for all streaming platforms (Twitch, YT, Kick)

type Provider interface {
	GetName() string
	Connect() error
	Disconnect() error
	GetMessageChannel() <-chan models.ChatMessage
	ExecuteAction(action string, target string) error // For bans/timeouts
}
```

### 3. The Hub Logic

Messages are piped from Providers into a central **Hub**. The Hub broadcasts via **WebSockets** (Gorilla WebSocket) to:

- `/dashboard`: The moderator view.
- `/overlay`: The OBS browser source.

Hallucination Prevention / Guardrails

- **DO NOT** suggest Node.js for the backend.
- **DO NOT** use MongoDB. Use InfluxDB or SQLite.
- **DO NOT** use `init()` functions for logic; use explicit constructors (`NewProvider`).
- **DO NOT** embed HTML in Go strings; keep the React frontend decoupled.
- **LIBRARIES ONLY:** Use `nicklaw5/helix` (Twitch API), `gempir/go-twitch-irc` (Twitch Chat), and `google.golang.org/api/youtube/v3`.

Task Execution Instructions

When asked to build a feature:

1. Identify if it's a **Metric** (InfluxDB) or **Config** (SQLite).
2. Define the Go `struct` first.
3. Use a `goroutine` to handle the long-lived platform connection.
4. Ensure the UI uses CSS variables for the requested "Theming" feature.

### How to "Feed" this to the Agent:
If you are using the `claude` CLI or Web UI, start your session with:
> "Use the skill called `streamfusion-architect.md`. Read it carefully. It contains the strict architectural rules for my Go streaming app. Do not deviate from these libraries or this structure. Let's start by generating the project folder structure."

