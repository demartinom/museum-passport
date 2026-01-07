package museums

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
