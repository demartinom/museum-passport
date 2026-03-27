package cache

import (
	"sync"

	"github.com/demartinom/museum-passport/models"
)

type Cache struct {
	Data      map[string]CachedItem
	summaries map[string]CachedSummary
	mu        sync.RWMutex
}

type CachedItem struct {
	Artwork models.SingleArtwork
}

type CachedSummary struct {
	Summary string
}

func NewCache() *Cache {
	return &Cache{Data: make(map[string]CachedItem), summaries: make(map[string]CachedSummary)}
}

// Adds artwork to site cache
// Uses ID ("met-123") as key and artwork struct as value
func (c *Cache) SetArtwork(id string, artwork models.SingleArtwork) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Data[id] = CachedItem{
		Artwork: artwork,
	}
}

// Search cache for artwork
// Returns false boolean if not in cache
func (c *Cache) GetArtwork(id string) (models.SingleArtwork, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	artwork, exists := c.Data[id]
	if !exists {
		return models.SingleArtwork{}, false
	}
	return artwork.Artwork, true
}

func (c *Cache) GetSummary(artworkID string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.summaries[artworkID]
	if !exists {
		return "", false
	}

	return item.Summary, true
}

func (c *Cache) SetSummary(artworkID, summary string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.summaries[artworkID] = CachedSummary{
		Summary: summary,
	}
}
