package museums

import (
	"bytes"
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
	Score        float64 `json:"_score"`
}

type ArticSearchResponse struct {
	Pagination Pagination           `json:"pagination"`
	Data       []ArticSingleArtwork `json:"data"`
}

// Returns pagination information for API call to artic
type Pagination struct {
	CurrentPage int `json:"current_page"`
	TotalPages  int `json:"total_pages"`
}

// Type for receiving single artwork from API
type ArticSingleArtworkResponse struct {
	Data ArticSingleArtwork `json:"data"`
}

type ElasticPayload struct {
	Query  ElasticQuery `json:"query"`
	Fields []string     `json:"fields"`
	Limit  int          `json:"limit"`
}

type ElasticQuery struct {
	Match map[string]string `json:"match"`
}

func NewArticClient(cache *cache.Cache) *ArticClient {
	return &ArticClient{BaseURL: "https://api.artic.edu/api/v1/artworks", Cache: cache}
}

func (a *ArticClient) GetMuseumName() string {
	return "Art Institute of Chicago"
}

func (a *ArticClient) ArtworkURL(id int) string {
	return fmt.Sprintf("https://www.artic.edu/artworks/%d", id)
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
		URL:          a.ArtworkURL(receivedArt.ID),
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("artic API error: %s", resp.Status)
	}
	defer resp.Body.Close()

	var result ArticSingleArtworkResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	normalized := a.NormalizeArtwork(result.Data)
	return &normalized, nil
}

func (a *ArticClient) GeneralSearch(query string, resultsLength int, pageNumber int) (*SearchResult, error) {
	queryURL := fmt.Sprintf("%s/search?q=%s&limit=%d&page=%d&fields=id,title,artist_title,image_id,medium_display,date_start,is_public_domain", a.BaseURL, url.QueryEscape(query), resultsLength, pageNumber)

	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("artic API error: %s", resp.Status)
	}
	defer resp.Body.Close()

	var searchResult ArticSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return nil, err
	}
	filtered := []ArticSingleArtwork{}

	for _, artwork := range searchResult.Data {
		if artwork.Score >= 1.0 && artwork.ImageID != "" {
			filtered = append(filtered, artwork)
		}
	}
	var normalized []*models.SingleArtwork
	for _, artwork := range filtered {
		art := a.NormalizeArtwork(artwork)
		normalized = append(normalized, &art)
	}

	return &SearchResult{ResultsLength: len(normalized), Art: normalized, TotalPages: searchResult.Pagination.TotalPages}, nil
}

func (a *ArticClient) Search(params SearchParams, pageLength int, pagNumber int) (*SearchResult, error) {
	if params.Name != "" {
		return a.SearchByField("title", params.Name, pageLength)
	}
	if params.Artist != "" {
		return a.SearchByField("artist_title", params.Artist, pageLength)
	}

	return nil, fmt.Errorf("no search parameters provided")
}

func (a *ArticClient) SearchByField(field string, fieldValue string, length int) (*SearchResult, error) {
	body := ElasticPayload{Query: ElasticQuery{Match: map[string]string{field: fieldValue}},
		Fields: []string{"id", "title", "artist_title", "image_id", "medium_display", "date_start", "is_public_domain"},
		Limit:  length}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(fmt.Sprintf("%s/search", a.BaseURL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("artic API error: %s", resp.Status)
	}

	defer resp.Body.Close()

	var searchResult ArticSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&searchResult); err != nil {
		return nil, err
	}
	var normalized []*models.SingleArtwork
	for _, artwork := range searchResult.Data {
		if artwork.Score >= 1.0 && artwork.ImageID != "" {
			art := a.NormalizeArtwork(artwork)
			normalized = append(normalized, &art)
		}
	}

	return &SearchResult{ResultsLength: len(normalized), Art: normalized}, nil
}
