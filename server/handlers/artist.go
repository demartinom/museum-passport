package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/demartinom/museum-passport/artist"
	"github.com/demartinom/museum-passport/cache"
)

type ArtistHandler struct {
	ArtistClient *artist.ArtistClient
	Cache        *cache.Cache
}

type ArtistResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Bio         string `json:"bio"`
	ImageURL    string `json:"imageUrl,omitempty"`
	Description string `json:"description,omitempty"`
}

func NewArtistHandler(a *artist.ArtistClient, c *cache.Cache) *ArtistHandler {
	return &ArtistHandler{ArtistClient: a, Cache: c}
}

func (h *ArtistHandler) GetArtist(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("artistname") // or however you're reading it

	result, err := h.ArtistClient.FindArtist(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // or map specific errors to specific codes
		return
	}

	resp := ArtistResponse{
		ID:          result.ID,
		Name:        result.Artist.Titles.Normalized,
		Bio:         result.Artist.Blurb,
		ImageURL:    result.Artist.Image.ImageURL,
		Description: result.Artist.Description,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
