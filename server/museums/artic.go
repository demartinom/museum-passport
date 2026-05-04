package museums

import "github.com/demartinom/museum-passport/cache"

type ArticClient struct {
	BaseURL string
	Cache   *cache.Cache
}

func NewArticClient(cache *cache.Cache) *ArticClient {
	return &ArticClient{BaseURL: "https://api.artic.edu/api/v1/artworks", Cache: cache}
}

type ArticSingleArtwork struct {
	ID           int    `json:"id"`
	DateStart    int    `json:"date_start"`
	Medium       string `json:"medium_display"`
	Artist       string `json:"artist_title"`
	ImageID      string `json:"image_id"`
	Title        string `json:"title"`
	PublicDomain bool   `json:"is_public_domain"`
}
