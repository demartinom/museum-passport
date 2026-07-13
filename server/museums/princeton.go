package museums

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/models"
)

// Client for handling calls to the Princeton API
type PrincetonClient struct {
	BaseURL string
	Cache   *cache.Cache
}

// Struct for receiving single artwork response from Princeton API
type PrincetonSingleArtwork struct {
	ID           int      `json:"objectid"`
	Dated        string   `json:"displaydate"`
	Medium       string   `json:"medium"`
	Artist       string   `json:"displaymaker"`
	PrimaryImage []string `json:"primaryimage"`
	Title        string   `json:"displaytitle"`
}

// Create new Princeton API client
func NewPrincetonClient(cache *cache.Cache) *PrincetonClient {
	return &PrincetonClient{BaseURL: "https://data.artmuseum.princeton.edu", Cache: cache}
}

// Allows for Princeton client to fall under museum interface
func (p *PrincetonClient) GetMuseumName() string {
	return "Princeton University Art Museum"
}

func (p *PrincetonClient) ArtworkURL(id int) string {
	return fmt.Sprintf("https://artmuseum.princeton.edu/art/collections/objects/%d", id)
}

// Takes Object API response store in PrincetonSingleArtwork
// Normalizes response into the models.Artwork struct and saves in cache
func (p *PrincetonClient) NormalizeArtwork(receivedArt PrincetonSingleArtwork) models.SingleArtwork {
	normalized := models.SingleArtwork{
		ID:           fmt.Sprintf("princeton-%d", receivedArt.ID),
		ArtworkTitle: receivedArt.Title,
		ArtistName:   receivedArt.Artist,
		ArtMedium:    receivedArt.Medium,
		ImageLarge:   fmt.Sprintf("%s/full/max/0/default.jpg", receivedArt.PrimaryImage[0]),
		ImageSmall:   fmt.Sprintf("%s/full/800,/0/default.jpg", receivedArt.PrimaryImage[0]),
		Museum:       p.GetMuseumName(),
		URL:          p.ArtworkURL(receivedArt.ID),
	}
	p.Cache.SetArtwork(normalized.ID, normalized)
	return normalized
}

// Makes an API call to Princeton to receive data on a single artwork based on id provided
func (p *PrincetonClient) ArtworkByID(id int) (*models.SingleArtwork, error) {
	artwork, exists := p.Cache.GetArtwork(fmt.Sprintf("princeton-%d", id))
	if exists {
		return &artwork, nil
	}

	queryURL := fmt.Sprintf("%s/objects/%d", p.BaseURL, id)

	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result PrincetonSingleArtwork
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	normalized := p.NormalizeArtwork(result)
	return &normalized, nil
}
