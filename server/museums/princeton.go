package museums

import (
	"fmt"

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
	ID           int    `json:"objectid"`
	Dated        string `json:"displaydate"`
	Medium       string `json:"medium"`
	Artist       string `json:"displaymaker"`
	PrimaryImage string `json:"primaryimage"`
	Title        string `json:"displaytitle"`
}

// Create new Princeton API client
func NewPrincetonClient(cache *cache.Cache) *PrincetonClient {
	return &PrincetonClient{BaseURL: "https://data.artmuseum.princeton.edu", Cache: cache}
}

// Allows for Princeton client to fall under museum interface
func (p *PrincetonClient) GetMuseumName() string {
	return "Princeton University Art Museum"
}

// Takes Object API response store in PrincetonSingleArtwork
// Normalizes response into the models.Artwork struct and saves in cache
func (p *PrincetonClient) NormalizeArtwork(receivedArt PrincetonSingleArtwork) models.SingleArtwork {
	normalized := models.SingleArtwork{
		ID:           fmt.Sprintf("princeton-%d", receivedArt.ID),
		ArtworkTitle: receivedArt.Title,
		ArtistName:   receivedArt.Artist,
		ArtMedium:    receivedArt.Medium,
		ImageLarge:   fmt.Sprintf("%s/full/max/0/default.jpg", receivedArt.PrimaryImage),
		ImageSmall:   fmt.Sprintf("%s/full/800,/0/default.jpg", receivedArt.PrimaryImage),
		Museum:       p.GetMuseumName(),
	}
	p.Cache.SetArtwork(normalized.ID, normalized)
	return normalized
}
