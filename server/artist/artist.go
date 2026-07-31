package artist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ArtistClient struct {
}

type Artist struct {
	Titles []struct {
		Name string `json:"normalized"`
	} `json:"titles"`
	Blurb string `json:"extract"`
	Image []struct {
		ImageURL string `json:"source"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
	} `json:"originalimage"`
	Description string `json:"description"`
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

func (a *ArtistClient) NewArtistClient() *ArtistClient {
	return &ArtistClient{}
}

func (a *ArtistClient) FindTitle(query string) (string, error) {
	encoded := url.PathEscape(query)
	queryUrl := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&list=search&srsearch=%s&srlimit=3&format=json&origin=*", encoded)

	resp, err := http.Get(queryUrl)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	var result CandidateSearchResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
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

	resp, err := http.Get(queryUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result Artist

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
