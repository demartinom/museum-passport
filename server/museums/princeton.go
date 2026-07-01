package museums

import "github.com/demartinom/museum-passport/cache"

type PrincetonClient struct {
	BaseURL string
	Cache   *cache.Cache
}

func NewPrincetonClient(cache *cache.Cache) *PrincetonClient {
	return &PrincetonClient{BaseURL: "https://data.artmuseum.princeton.edu/objects/32221", Cache: cache}
}
