package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/demartinom/museum-passport/cache"
)

type AOTDHandler struct {
	Cache *cache.Cache
}

func NewAOTDHandler(c *cache.Cache) *AOTDHandler {
	return &AOTDHandler{Cache: c}
}

func (a AOTDHandler) UpdateAOTD(w http.ResponseWriter, r *http.Request) {
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
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully rotated AOTD to: " + winner))
}
