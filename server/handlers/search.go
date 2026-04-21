package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/models"
	"github.com/demartinom/museum-passport/museums"
)

type SearchHandler struct {
	Clients map[string]museums.Client
	Cache   *cache.Cache
}

func NewSearchHandler(clients map[string]museums.Client, c *cache.Cache) *SearchHandler {
	return &SearchHandler{Clients: clients, Cache: c}
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
		resultsLength = 40
	}
	var artwork []*models.SingleArtwork

	for _, museum := range s.Clients {
		var foundArtwork *museums.SearchResult
		// general decides whether or not to search using specific criteria
		if general != "" {
			foundArtwork, err = museum.GeneralSearch(general, resultsLength/len(s.Clients))
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

	artwork = s.SortArtwork(artwork)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artwork)
}

func (s *SearchHandler) SortArtwork(artResults []*models.SingleArtwork) []*models.SingleArtwork {
	// Fetch all scores in one pass
	scores := make(map[string]float64, len(artResults))
	for _, art := range artResults {
		score, _ := s.Cache.GetScore(art.ID)
		scores[art.ID] = score
	}

	sort.Slice(artResults, func(i, j int) bool {
		return scores[artResults[i].ID] > scores[artResults[j].ID]
	})
	return artResults
}
