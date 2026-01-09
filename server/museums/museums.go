package museums

import "github.com/demartinom/museum-passport/models"

type Client interface {
	GetMuseumName() string
	ArtworkbyID(int) (*models.SingleArtwork, error)
}
