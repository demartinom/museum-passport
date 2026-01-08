package museums

import (
	"testing"
)

func TestNormalizeMetArtwork(t *testing.T) {
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
func TestNormalizeHarvardArtwork(t *testing.T) {
	raw := HarvardSingleArtwork{
		ID:     230120,
		Dated:  "c. 1884-1885",
		Medium: "Pot metal, opalescent, and uncolored glass with vitreous paint",
		People: struct {
			DisplayName string "json:\"displayname\""
		}{
			DisplayName: "John La Farge",
		},
		Primaryimageurl: "https://nrs.harvard.edu/urn-3:HUAM:38659_dynmc",
		Title:           "When the Morning Stars Sang Together and All the Sons of God Shouted for Joy",
	}
	client := NewHarvardClient("test")
	normalized := client.NormalizeArtwork(raw)

	if normalized.ID != 230120 {
		t.Errorf("Expected ID 230120, returned ID %d", normalized.ID)
	}

	if normalized.ArtworkTitle != "When the Morning Stars Sang Together and All the Sons of God Shouted for Joy" {
		t.Errorf("Expected title When the Morning Stars Sang Together and All the Sons of God Shouted for Joy, returned %s", normalized.ArtworkTitle)
	}
	if normalized.ArtistName != "John La Farge" {
		t.Errorf("Expected artist name John La Farge, returned %s", normalized.ArtistName)
	}
}
