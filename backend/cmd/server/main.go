package main

import (
	"log"
	"net/http"

	"github.com/adamwrose/streamfusion/internal/config"
	"github.com/adamwrose/streamfusion/internal/db/sqlite"
	"github.com/adamwrose/streamfusion/internal/hub"
	"github.com/adamwrose/streamfusion/internal/websocket"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sqlite.New(cfg.SQLitePath)
	if err != nil {
		log.Fatalf("failed to open sqlite: %v", err)
	}
	defer db.Close()

	h := hub.NewHub()
	go h.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/ws/dashboard", websocket.HandleDashboard(h))
	mux.HandleFunc("/ws/overlay", websocket.HandleOverlay(h))

	log.Printf("StreamFusion backend listening on %s", cfg.Addr)
	log.Fatal(http.ListenAndServe(cfg.Addr, mux))
}
