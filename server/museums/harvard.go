package museums

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/demartinom/museum-passport/models"
)

// Client for handling calls to the Harvard API
type HarvardClient struct {
	BaseURL string
	APIKey  string
}

// Struct for receiving single artwork response from Harvard API
// Receives AAPI key from .env file
type HarvardSingleArtwork struct {
	ID     int    `json:"id"`
	Dated  string `json:"dated"`
	Medium string `json:"medium"`
	People struct {
		DisplayName string `json:"displayname"`
	}
	Primaryimageurl string `json:"primaryimageurl"`
	Title           string `json:"title"`
}

// Create new Harvard API client
func NewHarvardClient(key string) *HarvardClient {
	return &HarvardClient{BaseURL: "https://api.harvardartmuseums.org", APIKey: key}
}

// Allows for Harvard Client to fall under museum interface
func (h *HarvardClient) GetMuseumName() string {
	return "Harvard Art Museums"
}

// Takes Object API response store in HarvardSingleArtwork and normalizes it into the models.Artwork struct
func (m *HarvardClient) NormalizeArtwork(receivedArt HarvardSingleArtwork) models.SingleArtwork {
	return models.SingleArtwork{
		ID:           fmt.Sprintf("harvard-%d", receivedArt.ID),
		ArtworkTitle: receivedArt.Title,
		ArtistName:   receivedArt.People.DisplayName,
		DateCreated:  receivedArt.Dated,
		ArtMedium:    receivedArt.Medium,
		ImageLarge:   receivedArt.Primaryimageurl,
		ImageSmall:   "",
		Museum:       "Harvard Art Museums",
	}
}

// Makes an API call to Harvard to receive data on a single artwork based on id provided
func (h *HarvardClient) ArtworkbyID(id int) (*models.SingleArtwork, error) {
	queryUrl := fmt.Sprintf("%s/object/%d?apikey=%s", h.BaseURL, id, h.APIKey)

	resp, err := http.Get(queryUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result HarvardSingleArtwork
	json.NewDecoder(resp.Body).Decode(&result)

	normalized := h.NormalizeArtwork(result)
	return &normalized, nil
}
