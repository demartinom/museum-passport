package museums

import "github.com/demartinom/museum-passport/models"

// Interface for museum clients
type Client interface {
	GetMuseumName() string
	ArtworkbyID(int) (*models.SingleArtwork, error)
	Search(SearchParams) ([]int, error)
}

type SearchParams struct {
	Name string
}
