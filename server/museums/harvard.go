package museums

type HarvardClient struct {
	BaseURL string
	APIKey  string
}

type HarvardSingleArtwork struct {
}

func NewHarvardClient(key string) *HarvardClient {
	return &HarvardClient{BaseURL: "https://api.harvardartmuseums.org", APIKey: key}
}
