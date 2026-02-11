package museums

import "github.com/demartinom/museum-passport/models"

// Interface for museum clients
type Client interface {
	GetMuseumName() string
	ArtworkbyID(int) (*models.SingleArtwork, error)
	Search(SearchParams, int) (*SearchResult, error)
}

// Used for translating search parameters from url for APIs
type SearchParams struct {
	Name        string
	Artist      string
	ArtworkType string
}

// Struct for organizing returned data
type SearchResult struct {
	ResultsLength int
	Art           []*models.SingleArtwork
}
