package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/demartinom/museum-passport/models"
	"github.com/demartinom/museum-passport/museums"
)

type SearchHandler struct {
	Clients map[string]museums.Client
}

func NewSearchHandler(clients map[string]museums.Client) *SearchHandler {
	return &SearchHandler{Clients: clients}
}

// API endpoint for searching for artwork
// Uses search function specific to museum specified
func (s *SearchHandler) SearchArtwork(w http.ResponseWriter, r *http.Request) {
	// museum := r.URL.Query().Get("museum")
	name := r.URL.Query().Get("name")
	artist := r.URL.Query().Get("artist")
	artworktype := r.URL.Query().Get("type")
	pageLength := r.URL.Query().Get("length")
	general := r.URL.Query().Get("general")

	resultsLength, err := strconv.Atoi(pageLength)
	if err != nil {
		return
	}
	var artwork []*models.SingleArtwork

	for _, museum := range s.Clients {
		var foundArtwork *museums.SearchResult
		// general decides whether or not to search using specific criteria
		if general != "" {
			foundArtwork, err = museum.GeneralSearch(general, 80/len(s.Clients))
			if err != nil {
				fmt.Println("Error:", err)
				continue // Skip this museum
			}
			artwork = append(artwork, foundArtwork.Art...)
		} else {
			foundArtwork, err = museum.Search(museums.SearchParams{Name: name, Artist: artist, ArtworkType: artworktype}, resultsLength)
			if err != nil {
				fmt.Println("Error:", err)
				continue // Skip this museum
			}
			artwork = append(artwork, foundArtwork.Art...)
		}

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artwork)
}
