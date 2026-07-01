package museums

import "github.com/demartinom/museum-passport/cache"

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
