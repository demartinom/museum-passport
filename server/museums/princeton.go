package museums

import "github.com/demartinom/museum-passport/cache"

// Client for handling calls to the Princeton API
type PrincetonClient struct {
	BaseURL string
	Cache   *cache.Cache
}

func NewPrincetonClient(cache *cache.Cache) *PrincetonClient {
	return &PrincetonClient{BaseURL: "https://data.artmuseum.princeton.edu", Cache: cache}
}

// Allows for Princeton client to fall under museum interface
func (p *PrincetonClient) GetMuseumName() string {
	return "Princeton University Art Museum"
}
