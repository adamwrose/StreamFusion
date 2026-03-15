package influx

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// Client wraps the InfluxDB client for writing time-series metrics.
type Client struct {
	client   influxdb2.Client
	writeAPI api.WriteAPIBlocking
	org      string
	bucket   string
}

func NewClient(url, token, org, bucket string) *Client {
	c := influxdb2.NewClient(url, token)
	return &Client{
		client:   c,
		writeAPI: c.WriteAPIBlocking(org, bucket),
		org:      org,
		bucket:   bucket,
	}
}

// WriteViewerCount records a viewer count data point.
func (c *Client) WriteViewerCount(ctx context.Context, platform string, count int64) error {
	p := influxdb2.NewPoint(
		"viewer_count",
		map[string]string{"platform": platform},
		map[string]interface{}{"count": count},
		time.Now(),
	)
	if err := c.writeAPI.WritePoint(ctx, p); err != nil {
		return fmt.Errorf("influx: write viewer_count: %w", err)
	}
	return nil
}

// WriteChatVelocity records messages-per-minute for a platform.
func (c *Client) WriteChatVelocity(ctx context.Context, platform string, mpm float64) error {
	p := influxdb2.NewPoint(
		"chat_velocity",
		map[string]string{"platform": platform},
		map[string]interface{}{"mpm": mpm},
		time.Now(),
	)
	if err := c.writeAPI.WritePoint(ctx, p); err != nil {
		return fmt.Errorf("influx: write chat_velocity: %w", err)
	}
	return nil
}

func (c *Client) Close() {
	c.client.Close()
}
