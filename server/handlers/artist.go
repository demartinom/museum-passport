package handlers

import "github.com/demartinom/museum-passport/cache"

type ArtistHandler struct {
	ArtistClient *artist.ArtistClient
	Cache        *cache.Cache
}

}

func NewArtistHandler(c *cache.Cache) *ArtistHandler {
	return &ArtistHandler{Cache: c}
}
