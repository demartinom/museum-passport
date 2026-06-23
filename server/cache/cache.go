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
	c.client.Set(ctx, key, summary, 60*24*time.Hour)
}

func (c *Cache) RecordView(id string) {
	redisID := "artwork:" + id

	go c.client.ZIncrBy(ctx, "artwork_popularity", 1, redisID)
}

func (c *Cache) GetScore(id string) (float64, error) {
	redisID := "artwork:" + id

	score, err := c.client.ZScore(ctx, "artwork_popularity", redisID).Result()
	if err == redis.Nil {
		return 0, nil
	}

	return score, nil
}
// Marks artwork found by SetAOTD as having been an AOTD
func (c *Cache) MarkAsChosen(id string) error {
	// 1. Get the current time as a Unix timestamp (seconds)
	now := float64(time.Now().Unix())

	// 2. Add it to a Sorted Set
	err := c.client.ZAdd(ctx, "aotd:history", redis.Z{
		Score:  now,
		Member: id,
	}).Err()

	return err
}
