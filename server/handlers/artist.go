package handlers

import "github.com/demartinom/museum-passport/cache"

type ArtistHandler struct {
	ArtistClient *artist.ArtistClient
	Cache        *cache.Cache
}

func NewArtistHandler(a *artist.ArtistClient, c *cache.Cache) *ArtistHandler {
	return &ArtistHandler{ArtistClient: a, Cache: c}
}

}
