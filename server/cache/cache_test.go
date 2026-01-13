package cache

import (
	"testing"

	"github.com/demartinom/museum-passport/models"
)

func TestGetArtwork_Success(t *testing.T) {
	cache := Cache{Data: map[string]CachedItem{"met-1111": {models.SingleArtwork{ID: "met-1111", ArtworkTitle: "Starry Night", ArtistName: "Vincent Van Gogh"}}}}

	_, exists := cache.GetArtwork("met-1111")
	if !exists {
		t.Errorf("Artwork should exist")
	}
}

func TestSetArtwork(t *testing.T) {
	cache := Cache{}

	artwork := models.SingleArtwork{ID: "met-1111", ArtworkTitle: "Starry Night", ArtistName: "Vincent Van Gogh"}
	cache.SetArtwork("met-1111", artwork)

	retrieved, found := cache.GetArtwork("met-1111")
	if !found {
		t.Errorf("Should have found artwork")
	}

	if retrieved.ArtworkTitle != artwork.ArtworkTitle {
		t.Errorf("Wanted %s, got %s", artwork.ArtworkTitle, retrieved.ArtworkTitle)
	}
}
