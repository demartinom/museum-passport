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
	return &museums.SearchResult{IDs: []int{24343, 213732, 4334}}, nil
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
