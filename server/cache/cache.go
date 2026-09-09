package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

// When a artwork page is visited, increment view count by 1
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

// Sets the next Artwork of the Day.
// Chosen by artwork with highest view count that hasn't been selected before
func (c *Cache) SetAOTD() (string, error) {
	// Returns viewed artwork by view count
	topArtworks, err := c.client.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:   "artwork_popularity",
		Start: 0,
		Stop:  49,
		Rev:   true,
	}).Result()
	if err != nil {
		return "", err
	}

	var winnerID string

	// Loop through strings to see if artwork has already been AOTD
	for _, id := range topArtworks {
		err := c.client.ZScore(ctx, "aotd:history", id).Err()

		if err == redis.Nil {
			winnerID = id
			break
		} else if err != nil {
			return "", err
		}
	}

	// Fallback
	if winnerID == "" {
		return "", fmt.Errorf("could not find an unchosen artwork in the top 50")
	}

	// Adds winner to AOTD history
	err = c.MarkAsChosen(winnerID)
	if err != nil {
		return "", err
	}

	return winnerID, nil
}

// Marks artwork found by SetAOTD as having been an AOTD
func (c *Cache) MarkAsChosen(id string) error {
	// 1. Get the current time as a Unix timestamp (seconds)
	now := float64(time.Now().Unix())

	// 2. Add it to a Sorted Set
	if err := c.client.ZAdd(ctx, "aotd:history", redis.Z{
		Score:  now,
		Member: id,
	}).Err(); err != nil {
		return err
	}
	return c.client.Set(ctx, "aotd:current", id, 0).Err()
}

// Returns information on current AOTD
func (c *Cache) GetCurrentAOTD() (*models.SingleArtwork, error) {
	// Get current AOTD ID
	currentID, err := c.client.Get(ctx, "aotd:current").Result()
	if err == redis.Nil {
		return nil, nil // Means no AOTD has been selected yet
	} else if err != nil {
		return nil, err
	}

	// Get artwork information for AOTD
	val, err := c.client.Get(ctx, currentID).Result()
	if err == redis.Nil {
		// Artwork expired from cache, return just the ID so caller can re-fetch
		return &models.SingleArtwork{ID: currentID}, nil
	} else if err != nil {
		return nil, err
	}

	// Return artwork
	var artwork models.SingleArtwork
	if err := json.Unmarshal([]byte(val), &artwork); err != nil {
		return nil, err
	}

	return &artwork, nil
}

// If AOTD was selected 30 days ago, put it back in the running for AOTD
func (c *Cache) RemoveOldAOTD() error {
	// Calculate 30 days as hour
	thirtyDaysAgo := time.Now().Add(-30 * 24 * time.Hour).Unix()

	// Convert time to string
	maxScore := fmt.Sprintf("%d", thirtyDaysAgo)

	// Delete anything with a score from 0 up to 30 days ago
	err := c.client.ZRemRangeByScore(ctx, "aotd:history", "-inf", maxScore).Err()

	return err
}

// Saves artist information to cache
func (c *Cache) SetArtist(id string, artist models.ArtistResult) {
	key := "artist:" + id

	data, err := json.Marshal(artist)
	if err != nil {
		return
	}

	// Lasts for 2 weeks
	c.client.Set(ctx, key, data, 14*24*time.Hour)
}

// If artist exists in cache, retrieve info
func (c *Cache) GetArtist(id string) (models.ArtistResult, bool) {
	key := "artist:" + id

	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return models.ArtistResult{}, false // genuine cache miss
	} else if err != nil {
		log.Printf("redis get error: %v", err) // connection/other issue, worth knowing about
		return models.ArtistResult{}, false
	}

	var result models.ArtistResult
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		log.Printf("cache unmarshal failed for artist %s: %v", id, err)
		return models.ArtistResult{}, false
	}

	if result.Artist == nil {
		log.Printf("cache entry for artist %s has nil Artist, treating as miss", id)
		return models.ArtistResult{}, false
	}

	return result, true
}
