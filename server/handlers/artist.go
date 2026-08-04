package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/demartinom/museum-passport/artist"
	"github.com/demartinom/museum-passport/cache"
)

type ArtistHandler struct {
	ArtistClient *artist.ArtistClient
	Cache        *cache.Cache
}

func NewArtistHandler(a *artist.ArtistClient, c *cache.Cache) *ArtistHandler {
	return &ArtistHandler{ArtistClient: a, Cache: c}
}

func (a *ArtistHandler) GetArtist(w http.ResponseWriter, r *http.Request) {
	// "Monet" placeholder test
	artist, err := a.ArtistClient.FindArtist("Monet")
	if err != nil {
		log.Printf("FindArtist error: %v", err) // TEMP debug
		http.Error(w, "No artist found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artist)
}
