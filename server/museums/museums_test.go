package museums

import (
	"testing"
)

func TestNormalizeArtwork(t *testing.T) {
	raw := MetSingleArtwork{
		ObjectID:          436105,
		ObjectName:        "The Death of Socrates",
		ArtistDisplayName: "Jacques-Louis David",
		ObjectDate:        "1787",
		Medium:            "Oil on canvas",
		Repository:        "The Metropolitan Museum of Art",
		PrimaryImage:      "https://images.metmuseum.org/CRDImages/ep/original/DP-13139-001.jpg",
		PrimaryImageSmall: "https://images.metmuseum.org/CRDImages/ep/web-large/DP-13139-001.jpg",
	}
	client := NewMetClient()
	normalized := client.NormalizeArtwork(raw)

	if normalized.ID != 436105 {
		t.Errorf("Expected ID 536105, returned ID %d", normalized.ID)
	}

	if normalized.ArtworkTitle != "The Death of Socrates" {
		t.Errorf("Expected title The Death of Socrates, returned %s", normalized.ArtworkTitle)
	}
}
