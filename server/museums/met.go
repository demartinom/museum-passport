package museums

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/demartinom/museum-passport/models"
)

// Client for handling calls to the Met API
type MetClient struct {
	BaseURL string
}

// Struct for receiving a single artwork from the Met API
type MetSingleArtwork struct {
	ObjectID          int    `json:"objectID"`
	ObjectName        string `json:"objectName"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ObjectDate        string `json:"objectDate"`
	Medium            string `json:"medium"`
	Repository        string `json:"repository"`
	PrimaryImage      string `json:"primaryImage"`
	PrimaryImageSmall string `json:"primaryImageSmall"`
}

// Start up new Met Client
func NewMetClient() *MetClient {
	return &MetClient{BaseURL: "https://collectionapi.metmuseum.org/public/collection/v1"}
}

// Takes Object API response store in MetSingleArtwork and normalizes it into the models.Artwork struct
func (m *MetClient) NormalizeArtwork(receivedArt MetSingleArtwork) models.SingleArtwork {
	return models.SingleArtwork{
		ID:           fmt.Sprintf("met-%d", receivedArt.ObjectID),
		ArtworkTitle: receivedArt.ObjectName,
		ArtistName:   receivedArt.ArtistDisplayName,
		DateCreated:  receivedArt.ObjectDate,
		ArtMedium:    receivedArt.Medium,
		ImageLarge:   receivedArt.PrimaryImage,
		ImageSmall:   receivedArt.PrimaryImageSmall,
		Museum:       receivedArt.Repository,
	}
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
