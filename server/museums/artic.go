package museums

import "github.com/demartinom/museum-passport/cache"

type ArticClient struct {
	BaseURL string
	Cache   *cache.Cache
}

func NewArticClient(key string, cache *cache.Cache) *ArticClient {
	return &ArticClient{BaseURL: "https://api.artic.edu/api/v1/artworks", Cache: cache}
}
