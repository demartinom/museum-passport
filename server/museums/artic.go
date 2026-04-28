package museums

import "github.com/demartinom/museum-passport/cache"

type ArticClient struct {
	BaseURL string
	Cache   *cache.Cache
}

type ArticSingleArtwork struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Date         string `json:"date_start"`
	Medium       string `json:"medium_display"`
	PublicDomain bool   `json:"is_public_domain"`
	ImageID      string `json:"image_id"`
}

type ArticSearchResponse struct {
	Data []ArticSingleArtwork `json:"data"`
}

func NewArticClient(key string, cache *cache.Cache) *ArticClient {
	return &ArticClient{BaseURL: "https://api.artic.edu/api/v1/artworks", Cache: cache}
}
// Allows for Artic client to fall under museum interface
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

func (a *ArticClient) NormalizeArtwork(receivedArt ArticSingleArtwork) models.SingleArtwork {
	normalized := models.SingleArtwork{
		ID:           fmt.Sprintf("artic-%d", receivedArt.Id),
		ArtworkTitle: receivedArt.Title,
		DateCreated:  receivedArt.Date,
		ArtMedium:    receivedArt.Medium,
		PublicDomain: receivedArt.PublicDomain,
		ImageLarge:   a.BuildImageURL(receivedArt.ImageID, 843),
		ImageSmall:   a.BuildImageURL(receivedArt.ImageID, 400),
		Museum:       a.GetMuseumName(),
	}
	a.Cache.SetArtwork(normalized.ID, normalized)

	return normalized
}

