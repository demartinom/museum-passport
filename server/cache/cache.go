package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/demartinom/museum-passport/models"
	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

var ctx = context.Background()

func NewCache(rdb *redis.Client) *Cache {
	return &Cache{client: rdb}
}

// Adds artwork to site cache
// Uses "artwork" + ID ("met-123") as key and artwork struct as value
func (c *Cache) SetArtwork(id string, artwork models.SingleArtwork) {
	key := "artwork:" + id

	data, err := json.Marshal(artwork)
	if err != nil {
		return
	}

	// Lasts for 2 weeks
	c.client.Set(ctx, key, data, 14*24*time.Hour)
}

// Search cache for artwork
// Returns false boolean if not in cache
func (c *Cache) GetArtwork(id string) (models.SingleArtwork, bool) {
	key := "artwork:" + id

	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return models.SingleArtwork{}, false
	}

	var artwork models.SingleArtwork
	if err := json.Unmarshal([]byte(val), &artwork); err != nil {
		return models.SingleArtwork{}, false
	}

	return artwork, true
}

func (c *Cache) GetSummary(artworkID string) (string, bool) {
	key := "summary:" + artworkID
	val, err := c.client.Get(ctx, key).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (c *Cache) SetSummary(artworkID, summary string) {
	key := "summary:" + artworkID
	c.client.Set(ctx, key, summary, 0)
}
