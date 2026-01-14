package museums

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/models"
)

// Client for handling calls to the Met API
type MetClient struct {
	BaseURL string
	Cache   *cache.Cache
}

// Struct for receiving a single artwork from the Met API
type MetSingleArtwork struct {
	ObjectID          int    `json:"objectID"`
	Title             string `json:"title"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ObjectDate        string `json:"objectDate"`
	Medium            string `json:"medium"`
	Repository        string `json:"repository"`
	PrimaryImage      string `json:"primaryImage"`
	PrimaryImageSmall string `json:"primaryImageSmall"`
}

// Start up new Met Client
func NewMetClient(cache *cache.Cache) *MetClient {
	return &MetClient{BaseURL: "https://collectionapi.metmuseum.org/public/collection/v1", Cache: cache}
}

// Allows for Met Client to fall under museum interface
func (m *MetClient) GetMuseumName() string {
	return "The Metropolitan Museum of Art"
}

// Takes Object API response store in MetSingleArtwork and normalizes it into the models.Artwork struct and saves in cache
func (m *MetClient) NormalizeArtwork(receivedArt MetSingleArtwork) models.SingleArtwork {
	normalized := models.SingleArtwork{
		ID:           fmt.Sprintf("met-%d", receivedArt.ObjectID),
		ArtworkTitle: receivedArt.Title,
		ArtistName:   receivedArt.ArtistDisplayName,
		DateCreated:  receivedArt.ObjectDate,
		ArtMedium:    receivedArt.Medium,
		ImageLarge:   receivedArt.PrimaryImage,
		ImageSmall:   receivedArt.PrimaryImageSmall,
		Museum:       m.GetMuseumName(),
	}
	m.Cache.SetArtwork(normalized.ID, normalized)
	return normalized
}

// Makes an API call to the Met to receive data on a single artwork based on id provided
func (m *MetClient) ArtworkbyID(id int) (*models.SingleArtwork, error) {
	queryUrl := fmt.Sprintf("%s/objects/%d", m.BaseURL, id)
	resp, err := http.Get(queryUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result MetSingleArtwork
	json.NewDecoder(resp.Body).Decode(&result)

	normalized := m.NormalizeArtwork(result)
	return &normalized, nil
}
