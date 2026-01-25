package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	museum := r.URL.Query().Get("museum")
	name := r.URL.Query().Get("name")
	pageLength := r.URL.Query().Get("length")

	resultsLength, err := strconv.Atoi(pageLength)
	if err != nil {
		return
	}

	artwork, err := s.Clients[museum].Search(museums.SearchParams{Name: name}, resultsLength)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "Invalid search", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artwork)
}
