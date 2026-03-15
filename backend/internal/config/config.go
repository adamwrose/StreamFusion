package config

import (
	"fmt"
	"os"
)

// Config holds all runtime configuration for the StreamFusion backend.
type Config struct {
	Addr        string
	SQLitePath  string
	InfluxURL   string
	InfluxToken string
	InfluxOrg   string
	InfluxBucket string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	addr := getEnv("SF_ADDR", ":8080")
	sqlitePath := getEnv("SF_SQLITE_PATH", "./streamfusion.db")
	influxURL := getEnv("SF_INFLUX_URL", "http://localhost:8086")
	influxToken := os.Getenv("SF_INFLUX_TOKEN")
	influxOrg := getEnv("SF_INFLUX_ORG", "streamfusion")
	influxBucket := getEnv("SF_INFLUX_BUCKET", "metrics")

	if influxToken == "" {
		return nil, fmt.Errorf("SF_INFLUX_TOKEN must be set")
	}

	return &Config{
		Addr:         addr,
		SQLitePath:   sqlitePath,
		InfluxURL:    influxURL,
		InfluxToken:  influxToken,
		InfluxOrg:    influxOrg,
		InfluxBucket: influxBucket,
	}, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
