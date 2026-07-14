package museums

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/models"
)

// Client for handling calls to the Princeton API
type PrincetonClient struct {
	BaseURL string
	Cache   *cache.Cache
}

// Struct for receiving single artwork response from Princeton API
type PrincetonSingleArtwork struct {
	ID           int      `json:"objectid"`
	Dated        string   `json:"displaydate"`
	Medium       string   `json:"medium"`
	Artist       string   `json:"displaymaker"`
	PrimaryImage []string `json:"primaryimage"`
	Title        string   `json:"displaytitle"`
}

// SearchResponse is the top-level response.
type PrincetonSearchResponse struct {
	Hits struct {
		Total int `json:"total"`
		Hits  []struct {
			Index  string                `json:"_index"`
			ID     string                `json:"_id"`
			Score  float64               `json:"_score"`
			Source PrincetonSearchObject `json:"_source"`
			Type   string                `json:"_type"`
		} `json:"hits"`
	} `json:"hits"`
}

// ArtObject holds the actual art object fields. DisplayMaker and DisplayDate
// can be null in the source data, so they're pointers.
type PrincetonSearchObject struct {
	DisplayMaker *string  `json:"displaymaker"`
	ObjectNumber string   `json:"objectnumber"`
	DisplayDate  *string  `json:"displaydate"`
	PrimaryImage []string `json:"primaryimage"`
	Medium       string   `json:"medium"`
	DisplayTitle string   `json:"displaytitle"`
	ObjectID     int      `json:"objectid"`
}

type PrincetonMakerSearch struct {
	Hits struct {
		Hits []struct {
			Source PrincetonMakerStruct `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

type PrincetonMakerStruct struct {
	MakerID     int    `json:"makerid"`
	DisplayName string `json:"displayname"`
}

// Converts search object struct into singleartwork struct for data normalization
func (o PrincetonSearchObject) ToSingleArtwork() PrincetonSingleArtwork {
	artist := ""
	if o.DisplayMaker != nil {
		artist = *o.DisplayMaker
	}
	dated := ""
	if o.DisplayDate != nil {
		dated = *o.DisplayDate
	}

	return PrincetonSingleArtwork{
		ID:           o.ObjectID,
		Dated:        dated,
		Medium:       o.Medium,
		Artist:       artist,
		PrimaryImage: o.PrimaryImage,
		Title:        o.DisplayTitle,
	}
}

// Create new Princeton API client
func NewPrincetonClient(cache *cache.Cache) *PrincetonClient {
	return &PrincetonClient{BaseURL: "https://data.artmuseum.princeton.edu", Cache: cache}
}

// Allows for Princeton client to fall under museum interface
func (p *PrincetonClient) GetMuseumName() string {
	return "Princeton University Art Museum"
}

func (p *PrincetonClient) ArtworkURL(id int) string {
	return fmt.Sprintf("https://artmuseum.princeton.edu/art/collections/objects/%d", id)
}

// Takes Object API response store in PrincetonSingleArtwork
// Normalizes response into the models.Artwork struct and saves in cache
func (p *PrincetonClient) NormalizeArtwork(receivedArt PrincetonSingleArtwork) models.SingleArtwork {
	normalized := models.SingleArtwork{
		ID:           fmt.Sprintf("princeton-%d", receivedArt.ID),
		ArtworkTitle: receivedArt.Title,
		ArtistName:   receivedArt.Artist,
		ArtMedium:    receivedArt.Medium,
		ImageLarge:   fmt.Sprintf("%s/full/max/0/default.jpg", receivedArt.PrimaryImage[0]),
		ImageSmall:   fmt.Sprintf("%s/full/800,/0/default.jpg", receivedArt.PrimaryImage[0]),
		Museum:       p.GetMuseumName(),
		URL:          p.ArtworkURL(receivedArt.ID),
	}
	p.Cache.SetArtwork(normalized.ID, normalized)
	return normalized
}

// Makes an API call to Princeton to receive data on a single artwork based on id provided
func (p *PrincetonClient) ArtworkByID(id int) (*models.SingleArtwork, error) {
	artwork, exists := p.Cache.GetArtwork(fmt.Sprintf("princeton-%d", id))
	if exists {
		return &artwork, nil
	}

	queryURL := fmt.Sprintf("%s/objects/%d", p.BaseURL, id)

	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result PrincetonSingleArtwork
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	normalized := p.NormalizeArtwork(result)
	return &normalized, nil
}

// Searchs API based on params ("Artist, Artwork title")
func (p *PrincetonClient) Search(params SearchParams, pageLength int, pageNumber int) (*SearchResult, error) {
	if params.Artist != "" {
		artistID, err := p.FindMakerID(params.Artist)
		if err != nil {
			return nil, err
		}
		artworks, err := p.SearchArtistWorks(artistID, pageLength, pageNumber)
		if err != nil {
			return nil, err
		}
		return artworks, nil
	}
	if params.Name != "" {
		artworks, err := p.SearchByArtwork(params.Name, pageLength, pageNumber)
		if err != nil {
			return nil, err
		}
		return artworks, nil
	}
	return nil, fmt.Errorf("no search parameters provided")
}

// Returns the ID of artist searched to use API's maker search
func (p *PrincetonClient) FindMakerID(name string) (int, error) {
	queryURL := fmt.Sprintf("%s/search?q=%s&type=makers", p.BaseURL, name)
	resp, err := http.Get(queryURL)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Princeton API returned %d", resp.StatusCode)
	}

	var result PrincetonMakerSearch
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return 0, err
	}

	if len(result.Hits.Hits) == 0 {
		return 0, fmt.Errorf("No maker found matching %q", name)
	}

	return result.Hits.Hits[0].Source.MakerID, nil
}

// Takes artist ID and returns works by artist
func (p *PrincetonClient) SearchArtistWorks(id int, resultslength int, pageNumber int) (*SearchResult, error) {
	from := (pageNumber - 1) * resultslength
	queryURL := fmt.Sprintf("%s/objects?maker=%d&size=%d&from=%d", p.BaseURL, id, resultslength, from)

	resp, err := http.Get(queryURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Princeton API returned %d", resp.StatusCode)
	}
	var result PrincetonSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var normalized []*models.SingleArtwork
	for _, artwork := range result.Hits.Hits {
		art := p.NormalizeArtwork(artwork.Source.ToSingleArtwork())
		normalized = append(normalized, &art)
	}

	return &SearchResult{ResultsLength: len(normalized), Art: normalized}, nil
}

// search is the shared implementation behind GeneralSearch and SearchByArtwork.
// artType corresponds to the API's `type` query param ("all", "artobjects", etc).
func (p *PrincetonClient) search(query string, artType string, resultsLength int, pageNumber int) (*SearchResult, error) {
	from := (pageNumber - 1) * resultsLength
	queryURL := fmt.Sprintf("%s/search?q=%s&type=%s&size=%d&from=%d",
		p.BaseURL, url.QueryEscape(query), artType, resultsLength, from)

	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Princeton API returned %d", resp.StatusCode)
	}

	var result PrincetonSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var normalized []*models.SingleArtwork
	for _, artwork := range result.Hits.Hits {
		art := p.NormalizeArtwork(artwork.Source.ToSingleArtwork())
		normalized = append(normalized, &art)
	}

	return &SearchResult{ResultsLength: len(normalized), Art: normalized}, nil
}

// GeneralSearch searches across all Princeton data types.
func (p *PrincetonClient) GeneralSearch(query string, resultsLength int, pageNumber int) (*SearchResult, error) {
	return p.search(query, "all", resultsLength, pageNumber)
}

// SearchByArtwork restricts the search to art objects only.
func (p *PrincetonClient) SearchByArtwork(query string, resultsLength int, pageNumber int) (*SearchResult, error) {
	return p.search(query, "artobjects", resultsLength, pageNumber)
}
