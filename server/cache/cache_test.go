package cache

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/demartinom/museum-passport/models"
	"github.com/redis/go-redis/v9"
)

func newTestCache(t *testing.T) *Cache {
	mr := miniredis.RunT(t)
	return NewCache(redis.NewClient(&redis.Options{Addr: mr.Addr()}))
}

func TestSetAndGetArtwork(t *testing.T) {
	c := newTestCache(t)
	artwork := models.SingleArtwork{ID: "met-1111", ArtworkTitle: "Starry Night", ArtistName: "Vincent Van Gogh"}
	c.SetArtwork("met-1111", artwork)

	got, ok := c.GetArtwork("met-1111")
	if !ok || got.ArtworkTitle != artwork.ArtworkTitle {
		t.Fail()
	}
}

func TestGetArtwork_Missing(t *testing.T) {
	c := newTestCache(t)
	_, ok := c.GetArtwork("does-not-exist")
	if ok {
		t.Fail()
	}
}
