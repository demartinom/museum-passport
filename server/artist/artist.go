package artist

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/demartinom/museum-passport/cache"
	"github.com/demartinom/museum-passport/models"
)

type ArtistClient struct {
	Cache *cache.Cache
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

// Finds correctly formatted artist name and page id
func (a *ArtistClient) FindTitle(query string) (string, int, error) {
	encoded := url.QueryEscape(query)
	queryUrl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch=%s&srlimit=3&format=json&origin=*", encoded)

	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		return "", 0, err
	}
	// Headers to prevent blocking of api call
	req.Header.Set("User-Agent", "museum-passport/1.0 (https://museum-passport.vercel.app; contact@yourdomain.com)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	// Error handling
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("wikipedia search returned status %d", resp.StatusCode)
	}

	var result CandidateSearchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", 0, err
	}

	if len(result.Query.Search) == 0 {
		return "", 0, fmt.Errorf("no results found for query: %q", query)
	}

	return result.Query.Search[0].Title, result.Query.Search[0].PageID, nil
}

// Searches wikipedia API for data on artist
func (a *ArtistClient) FindArtist(query string) (*models.ArtistResult, error) {
	// Retrieve name and ID
	pageTitle, artistID, err := a.FindTitle(query)
	if err != nil {
		return nil, err
	}

	// Check if artist already exists in cache
	if cached, exists := a.Cache.GetArtist(strconv.Itoa(artistID)); exists {
		log.Println("cache hit for artist", artistID)
		return &cached, nil
	}

	log.Println("cache miss, fetching from wikipedia")

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

	var result models.Artist
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("decode failed: %w, body: %s", err, string(body))
	}

	return &models.ArtistResult{Artist: &result, ID: artistID}, nil
}
