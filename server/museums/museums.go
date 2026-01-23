package museums

import "github.com/demartinom/museum-passport/models"

// Interface for museum clients
type Client interface {
	GetMuseumName() string
	ArtworkbyID(int) (*models.SingleArtwork, error)
	Search(SearchParams, int) (*SearchResult, error)
}

type SearchParams struct {
	Name string
}

// Temporary struct to keep interface
type SearchResult struct {
	ResultsLength int
	Art           []*models.SingleArtwork
}
