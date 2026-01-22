package handlers

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/demartinom/museum-passport/models"
	"github.com/demartinom/museum-passport/museums"
)

// Struct and methods allow for use of Client interfact
type MockMuseumClient struct {
	artwork *models.SingleArtwork
	err     error
}

func (m *MockMuseumClient) ArtworkbyID(id int) (*models.SingleArtwork, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.artwork, nil
}

func (m *MockMuseumClient) GetMuseumName() string {
	return "Mock Museum"
}

func (m *MockMuseumClient) Search(params museums.SearchParams) (*museums.SearchResult, error) {
	if m.err != nil {
		return nil, m.err
	}
	// Return mock IDs for testing
	return &museums.SearchResult{Art: []*models.SingleArtwork{{
		ID:           "met-205423",
		ArtworkTitle: "Dog kennel",
		ArtistName:   "Claude I Sené",
		DateCreated:  "ca. 1775–80",
		ArtMedium:    "Gilded beech and pine; silk and velvet",
		ImageLarge:   "https://images.metmuseum.org/CRDImages/es/original/DT241800.jpg",
		ImageSmall:   "https://images.metmuseum.org/CRDImages/es/web-large/DT241800.jpg",
		Museum:       "The Metropolitan Museum of Art",
	},
		{
			ID:           "met-544519",
			ArtworkTitle: "Mechanical Dog",
			ArtistName:   "",
			DateCreated:  "ca. 1390–1352 B.C.",
			ArtMedium:    "Ivory (elephant)",
			ImageLarge:   "https://images.metmuseum.org/CRDImages/eg/original/0227r2_SEC501K.jpg",
			ImageSmall:   "https://images.metmuseum.org/CRDImages/eg/web-large/0227r2_SEC501K.jpg",
			Museum:       "The Metropolitan Museum of Art",
		}}}, nil
}

func TestGetArtwork_Success(t *testing.T) {
	handler := &ArtworkHandler{
		Clients: map[string]museums.Client{
			"met": &MockMuseumClient{},
		},
	}

	req := httptest.NewRequest("GET", "/api/artwork/met-123", nil)
	w := httptest.NewRecorder()

	handler.GetArtwork(w, req)

	if w.Code != 200 {
		t.Errorf("got status %d, want 200", w.Code)
	}
}

func TestArtwork_UnknownMuseum(t *testing.T) {

	handler := &ArtworkHandler{
		Clients: map[string]museums.Client{
			"guggenheim": &MockMuseumClient{},
		},
	}
	req := httptest.NewRequest("GET", "/api/artwork/fake-2324", nil)
	w := httptest.NewRecorder()

	handler.GetArtwork(w, req)

	if w.Code != 404 {
		t.Errorf("got status %d, want 404", w.Code)
	}
}

func TestGetArtwork_InvalidID(t *testing.T) {
	handler := &ArtworkHandler{
		Clients: map[string]museums.Client{
			"met": &MockMuseumClient{},
		},
	}

	req := httptest.NewRequest("GET", "/api/artwork/met-abc", nil)
	w := httptest.NewRecorder()

	handler.GetArtwork(w, req)

	if w.Code != 400 {
		t.Errorf("got status %d, want 200", w.Code)
	}
}

func TestGetArtwork_NoArtworkFound(t *testing.T) {
	handler := &ArtworkHandler{
		Clients: map[string]museums.Client{
			"met": &MockMuseumClient{err: fmt.Errorf("artwork not found")},
		},
	}

	req := httptest.NewRequest("GET", "/api/artwork/met-0000", nil)
	w := httptest.NewRecorder()

	handler.GetArtwork(w, req)

	if w.Code != 404 {
		t.Errorf("got status %d, want 404", w.Code)
	}
}
