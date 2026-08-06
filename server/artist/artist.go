package artist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/demartinom/museum-passport/cache"
)

type ArtistClient struct {
	Cache *cache.Cache
}

type Artist struct {
	Titles struct {
		Normalized string `json:"normalized"`
	} `json:"titles"`
	Blurb string `json:"extract"`
	Image struct {
		ImageURL string `json:"source"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
	} `json:"originalimage"`
	Description string `json:"description"`
}

func NewArtistClient(c *cache.Cache) *ArtistClient {
	return &ArtistClient{Cache: c}
}

type SearchCandidate struct {
	Title  string `json:"title"`
	PageID int    `json:"pageid"`
}

type CandidateSearchResponse struct {
	Query struct {
		Search []SearchCandidate `json:"search"`
	} `json:"query"`
}

func (a *ArtistClient) FindTitle(query string) (string, error) {
	encoded := url.QueryEscape(query)
	queryUrl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch=%s&srlimit=3&format=json&origin=*", encoded)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return "", err
	}
	// Headers to prevent blocking of api call
	req.Header.Set("User-Agent", "museum-passport/1.0 (https://museum-passport.vercel.app; contact@yourdomain.com)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Error handling
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("wikipedia search returned status %d", resp.StatusCode)
	}

	var result CandidateSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Query.Search) == 0 {
		return "", fmt.Errorf("no results found for query: %q", query)
	}

	return result.Query.Search[0].Title, nil
}

func (a *ArtistClient) FindArtist(query string) (*Artist, error) {
	pageTitle, err := a.FindTitle(query)
	if err != nil {
		return nil, err
	}
	encoded := url.PathEscape(strings.ReplaceAll(pageTitle, " ", "_"))
	queryUrl := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", encoded)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return nil, err
	}
	// Headers to prevent blocking of api call
	req.Header.Set("User-Agent", "museum-passport/1.0 (https://museum-passport.vercel.app; contact@yourdomain.com)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Error handling
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wikipedia summary returned status %d: %s", resp.StatusCode, string(body))
	}

	var result Artist
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decode failed: %w, body: %s", err, string(body))
	}

	return &result, nil
}
