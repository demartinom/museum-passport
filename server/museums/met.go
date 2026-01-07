package museums

import "github.com/demartinom/museum-passport/models"

type MetClient struct {
	BaseURL string
}

type MetSingleArtwork struct {
	ObjectID          int    `json:"objectID"`
	ObjectName        string `json:"objectName"`
	ArtistDisplayName string `json:"artistDisplayName"`
	ObjectDate        string `json:"objectDate"`
	Medium            string `json:"medium"`
	Repository        string `json:"repository"`
	PrimaryImage      string `json:"primaryImage"`
	PrimaryImageSmall string `json:"primaryImageSmall"`
}

func NewMetClient() *MetClient {
	return &MetClient{BaseURL: "https://collectionapi.metmuseum.org/public/collection/v1"}
}

// Takes Object API response store in MetSingleArtwork and normalizes it into the models.Artwork struct
func (m *MetClient) NormalizeArtwork(receivedArt MetSingleArtwork) models.SingleArtwork {
	return models.SingleArtwork{
		ID:           receivedArt.ObjectID,
		ArtworkTitle: receivedArt.ObjectName,
		ArtistName:   receivedArt.ArtistDisplayName,
		DateCreated:  receivedArt.ObjectDate,
		ArtMedium:    receivedArt.Medium,
		ImageLarge:   receivedArt.PrimaryImage,
		ImageSmall:   receivedArt.PrimaryImageSmall,
		Museum:       receivedArt.Repository,
	}
}
