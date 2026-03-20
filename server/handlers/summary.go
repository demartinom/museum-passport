package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/demartinom/museum-passport/ai"
	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/museums"
)

type SummaryHandler struct {
	summaryClient *ai.SummaryClient
	cache         *cache.Cache
	clients       map[string]museums.Client
}

func NewSummaryHandler(summaryClient *ai.SummaryClient, cache *cache.Cache, clients map[string]museums.Client) *SummaryHandler {
	return &SummaryHandler{
		summaryClient: summaryClient,
		cache:         cache,
		clients:       clients,
	}
}

func (h *SummaryHandler) GenerateSummary(w http.ResponseWriter, r *http.Request) {
	artworkID := r.URL.Query().Get("id")
	if artworkID == "" {
		http.Error(w, "artwork ID required", http.StatusBadRequest)
		return
	}

	// Check cache first
	if summary, found := h.cache.GetSummary(artworkID); found {
		json.NewEncoder(w).Encode(map[string]string{"summary": summary})
		return
	}

	// Get artwork details
	parts := strings.SplitN(artworkID, "-", 2)
	if len(parts) != 2 {
		http.Error(w, "invalid artwork ID", http.StatusBadRequest)
		return
	}

	museumName := parts[0]
	artworkIDNum, err := strconv.Atoi(parts[1])
	if err != nil {
		return
	}

	client, exists := h.clients[museumName]
	if !exists {
		http.Error(w, "unknown museum", http.StatusNotFound)
		return
	}

	artwork, err := client.ArtworkbyID(artworkIDNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate summary
	summary, err := h.summaryClient.GenerateSummary(*artwork)
	if err != nil {
		http.Error(w, "failed to generate summary", http.StatusInternalServerError)
		return
	}

	// Cache it
	h.cache.SetSummary(artworkID, summary)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"summary": summary})
}
