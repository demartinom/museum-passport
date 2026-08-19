package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/demartinom/museum-passport/cache"
)

type AOTDHandler struct {
	Cache   *cache.Cache
	Clients map[string]museums.Client
}

func NewAOTDHandler(c *cache.Cache) *AOTDHandler {
	return &AOTDHandler{Cache: c}
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artwork)
}
