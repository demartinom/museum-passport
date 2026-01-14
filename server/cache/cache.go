package cache

import (
	"sync"

	"github.com/demartinom/museum-passport/models"
)

type Cache struct {
	Data map[string]CachedItem
	mu   sync.RWMutex
}

type CachedItem struct {
	Artwork models.SingleArtwork
}

func NewCache() *Cache {
	return &Cache{Data: make(map[string]CachedItem)}
}

func (c *Cache) SetArtwork(id string, artwork models.SingleArtwork) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Data[id] = CachedItem{
		Artwork: artwork,
	}
}

func (c *Cache) GetArtwork(id string) (models.SingleArtwork, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	artwork, exists := c.Data[id]
	if !exists {
		return models.SingleArtwork{}, false
	}
	return artwork.Artwork, true
}
