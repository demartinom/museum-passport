package museums

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/models"
	"golang.org/x/sync/errgroup"
)

// Client for handling calls to the Met API
type MetClient struct {
	BaseURL string
	Cache   *cache.Cache
}

// Struct for receiving a single artwork from the Met API
type MetSingleArtwork struct {
	ObjectID          int    `json:"objectID"`
	Title             string `json:"title"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ObjectDate        string `json:"objectDate"`
	Medium            string `json:"medium"`
	PublicDomain      bool   `json:"isPublicDomain"`
	Repository        string `json:"repository"`
	PrimaryImage      string `json:"primaryImage"`
	PrimaryImageSmall string `json:"primaryImageSmall"`
	Classification    string `json:"classification"`
}

// Struct for receiving search API response
type MetSearchResponse struct {
	Total     int   `json:"total"`
	ObjectIDs []int `json:"objectIDs"`
}

// Start up new Met Client
func NewMetClient(cache *cache.Cache) *MetClient {
	return &MetClient{BaseURL: "https://collectionapi.metmuseum.org/public/collection/v1", Cache: cache}
}

// Allows for Met Client to fall under museum interface
func (m *MetClient) GetMuseumName() string {
	return "The Metropolitan Museum of Art"
}

// Takes Object API response store in MetSingleArtwork and normalizes it into the models.Artwork struct and saves in cache
func (m *MetClient) NormalizeArtwork(receivedArt MetSingleArtwork) models.SingleArtwork {
	normalized := models.SingleArtwork{
		ID:           fmt.Sprintf("met-%d", receivedArt.ObjectID),
		ArtworkTitle: receivedArt.Title,
		ArtistName:   receivedArt.ArtistDisplayName,
		DateCreated:  receivedArt.ObjectDate,
		ArtMedium:    receivedArt.Medium,
		ImageLarge:   receivedArt.PrimaryImage,
		ImageSmall:   receivedArt.PrimaryImageSmall,
		Museum:       m.GetMuseumName(),
		PublicDomain: receivedArt.PublicDomain,
		ArtworkType:  receivedArt.Classification,
	}
	m.Cache.SetArtwork(normalized.ID, normalized)
	return normalized
}

// Makes an API call to the Met to receive data on a single artwork based on id provided
func (m *MetClient) ArtworkbyID(id int) (*models.SingleArtwork, error) {
	artwork, exists := m.Cache.GetArtwork(fmt.Sprintf("met-%d", id))
	if exists {
		return &artwork, nil
	}

	queryUrl := fmt.Sprintf("%s/objects/%d", m.BaseURL, id)
	resp, err := http.Get(queryUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result MetSingleArtwork
	json.NewDecoder(resp.Body).Decode(&result)

	normalized := m.NormalizeArtwork(result)
	return &normalized, nil
}

// Search for artwork
// Currently only uses title when searching
func (m *MetClient) Search(params SearchParams, resultsLength int) (*SearchResult, error) {
	queryURL := m.BuildURL(params)

	resp, err := http.Get(queryURL)
	if err != nil {
		return &SearchResult{}, err
	}
	defer resp.Body.Close()

	var result MetSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	// Returns full data for artwork IDs returned in API call
	artObjects, err := m.SearchRequest(result.ObjectIDs, resultsLength)
	if err != nil {
		return nil, err
	}
	return artObjects, nil
}

// Build url for making API call
// Parameter order matters
func (m *MetClient) BuildURL(params SearchParams) string {
	if params.Name != "" {
		return fmt.Sprintf("%s/search?title=true&q=%s",
			m.BaseURL, url.QueryEscape(params.Name))
	}
	if params.Artist != "" {
		return fmt.Sprintf("%s/search?hasImages=true&artistOrCulture=true&q=%s", m.BaseURL, url.QueryEscape(params.Artist))
	}
	if params.ArtworkType != "" {
		return fmt.Sprintf("%s/search?hasImages=true&medium=%s&q=*", m.BaseURL, url.QueryEscape(params.ArtworkType))
	}

	return fmt.Sprintf("%s/search", m.BaseURL)
}

// Handles full search of individual artworks
func (m *MetClient) SearchRequest(searchIDs []int, resultsLength int) (*SearchResult, error) {
	var currentSearch []int
	// Limits number of calls to resultsLength number
	if len(searchIDs) > resultsLength {
		currentSearch = searchIDs[:resultsLength]
	} else {
		currentSearch = searchIDs
	}
	artworks := make([]*models.SingleArtwork, len(currentSearch))

	// Handles concurrent calls to Met API for returning full data objects on individual artworks
	g := new(errgroup.Group)
	g.SetLimit(10)

	for i, id := range currentSearch {
		g.Go(func() error {
			artwork, err := m.ArtworkbyID(id)
			if err != nil {
				return err
			}
			artworks[i] = artwork
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Filter out public domain artworks that have no images
	filtered := make([]*models.SingleArtwork, 0, len(artworks))
	for _, artwork := range artworks {
		if artwork != nil && artwork.PublicDomain == true && artwork.ImageLarge != "" {
			filtered = append(filtered, artwork)
		} else if artwork.PublicDomain == false {
			filtered = append(filtered, artwork)
		}
	}
	return &SearchResult{ResultsLength: len(searchIDs), Art: filtered}, nil
}
