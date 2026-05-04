package museums

import "github.com/demartinom/museum-passport/cache"

type ArticClient struct {
	BaseURL string
	Cache   *cache.Cache
}
type ArticSingleArtwork struct {
	ID        int    `json:"id"`
	DateStart int    `json:"date_start"`
	Medium    string `json:"medium_display"`
	// Pointer to account for null value
	Artist       *string `json:"artist_title"`
	ImageID      string  `json:"image_id"`
	Title        string  `json:"title"`
	PublicDomain bool    `json:"is_public_domain"`
}

func NewArticClient(cache *cache.Cache) *ArticClient {
	return &ArticClient{BaseURL: "https://api.artic.edu/api/v1/artworks", Cache: cache}
}

func (a *ArticClient) GetMuseumName() string {
	return "Art Institute of Chicago"
}

// Takes imageID from api call and creates image URL
// Input width for different sized images
func (a *ArticClient) BuildImageURL(imageID string, width int) string {
	if imageID == "" {
		return ""
	}
	return fmt.Sprintf(
		"https://www.artic.edu/iiif/2/%s/full/%d,/0/default.jpg",
		imageID, width,
	)
}
