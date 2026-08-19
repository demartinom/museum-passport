package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/museums"
)

type AOTDHandler struct {
	Cache   *cache.Cache
	Clients map[string]museums.Client
}

func NewAOTDHandler(c *cache.Cache, clients map[string]museums.Client) *AOTDHandler {
	return &AOTDHandler{Cache: c, Clients: clients}
}

func (a *AOTDHandler) UpdateAOTD(w http.ResponseWriter, r *http.Request) {
	secretToken := r.Header.Get("Authorization")
	if secretToken != os.Getenv("AOTD_PASS") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := a.Cache.RemoveOldAOTD(); err != nil {
		log.Printf("ERROR: failed to prune old AOTD history: %v", err)
	}

	winner, err := a.Cache.SetAOTD()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully rotated AOTD to: " + winner))
}

func (a *AOTDHandler) GetAOTD(w http.ResponseWriter, r *http.Request) {
	artwork, err := a.Cache.GetCurrentAOTD()
	if err != nil {
		http.Error(w, "failed to fetch AOTD", http.StatusInternalServerError)
		return
	}

	if artwork == nil {
		http.Error(w, "No artwork of the day selected yet", http.StatusNotFound)
		return
	}

	if artwork.ArtworkTitle == "" {
		artworkInfo := strings.SplitN(artwork.ID, "-", 2)
		// Remove "artwork:" from string
		artworkInfo[0] = strings.ReplaceAll(artworkInfo[0], "artwork:", "")
		client, valid := a.Clients[artworkInfo[0]]
		if !valid {
			http.Error(w, "Unknown museum", http.StatusNotFound)
			return
		}
		IDNum, err := strconv.Atoi(artworkInfo[1])
		if err != nil {
			http.Error(w, "Invalid artwork ID", http.StatusInternalServerError)
			return
		}
		artwork, err = client.ArtworkByID(IDNum)
		if err != nil {
			http.Error(w, "Failed to re-fetch artwork", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artwork)
}
