package museums

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/models"
)

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
	Score        int     `json:"_score"`
}

type ArticSearchResponse struct {
	Data []ArticSingleArtwork `json:"data"`
}

type ArticSingleArtworkResponse struct {
	Data ArticSingleArtwork `json:"data"`
}

func NewArticClient(cache *cache.Cache) *ArticClient {
	return &ArticClient{BaseURL: "https://api.artic.edu/api/v1/artworks", Cache: cache}
}

func (a *ArticClient) GetMuseumName() string {
	return "Art Institute of Chicago"
}

func (a *ArticClient) NormalizeArtwork(receivedArt ArticSingleArtwork) models.SingleArtwork {
	artistName := "Unknown Artist"

	if receivedArt.Artist != nil {
		artistName = *receivedArt.Artist
	}

	normalized := models.SingleArtwork{
		ID:           fmt.Sprintf("artic-%d", receivedArt.ID),
		ArtworkTitle: receivedArt.Title,
		ArtistName:   artistName,
		ArtMedium:    receivedArt.Medium,
		ImageLarge:   a.BuildImageURL(receivedArt.ImageID, 843),
		ImageSmall:   a.BuildImageURL(receivedArt.ImageID, 400),
		Museum:       a.GetMuseumName(),
	}

	if a.Cache != nil {
		a.Cache.SetArtwork(normalized.ID, normalized)
	}
	return normalized
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

func (a *ArticClient) ArtworkByID(id int) (*models.SingleArtwork, error) {
	if a.Cache != nil {
		artwork, exists := a.Cache.GetArtwork(fmt.Sprintf("artic-%d", id))

		if exists {
			return &artwork, nil
		}
	}

	queryUrl := fmt.Sprintf("%s/%d", a.BaseURL, id)

	resp, err := http.Get(queryUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result ArticSingleArtworkResponse
	json.NewDecoder(resp.Body).Decode(&result)
	normalized := a.NormalizeArtwork(result.Data)

	return &normalized, nil
}

func (a *ArticClient) GeneralSearch(query string, resultsLength int) (*SearchResult, error) {
	queryURL := fmt.Sprintf("%s/search?q=%s&limit=%d&fields=id,title,image_id,medium_display,date_start,is_public_domain", a.BaseURL, url.QueryEscape(query), resultsLength)

	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var searchResult ArticSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return nil, err
	}
	filtered := []ArticSingleArtwork{}

	for _, artwork := range searchResult.Data {
		if artwork.Score >= 1 {
			filtered = append(filtered, artwork)
		}
	}
	var normalized []*models.SingleArtwork
	for _, artwork := range filtered {
		art := a.NormalizeArtwork(artwork)
		normalized = append(normalized, &art)
	}

	return &SearchResult{ResultsLength: len(normalized), Art: normalized}, nil
}
