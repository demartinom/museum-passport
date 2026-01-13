package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/demartinom/museum-passport/museums"
)

// Handler for artwork related routes
type ArtworkHandler struct {
	Clients map[string]museums.Client
}

func NewArtworkHandler(clients map[string]museums.Client) *ArtworkHandler {
	return &ArtworkHandler{Clients: clients}
}

// Returns a single artwork from a singular museum API in the normalized Artwork struct
func (a *ArtworkHandler) GetArtwork(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/api/artwork/")

	// Should be [museum name, artwork id]
	artworkInfo := strings.SplitN(id, "-", 2)

	museum := artworkInfo[0]
	artworkID := artworkInfo[1]

	client, valid := a.Clients[museum]
	if !valid {
		http.Error(w, "Unknown Museum", http.StatusNotFound)
		return
	}

	IDNum, err := strconv.Atoi(artworkID)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	artwork, err := client.ArtworkbyID(IDNum)
	if err != nil {
		http.Error(w, "No artwork found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artwork)
}
