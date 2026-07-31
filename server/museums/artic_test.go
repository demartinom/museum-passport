package museums

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestArtworkByID(t *testing.T) {
// 	// Define what the fake server returns
// 	fakeResponse := `{
//         "data": {
//             "id": 42,
//             "title": "Nighthawks",
//             "artist_title": "Edward Hopper",
//             "medium_display": "Oil on canvas",
//             "image_id": "img-999"
//         }
//     }`

// 	// Spin up a local test server
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		fmt.Fprint(w, fakeResponse)
// 	}))
// 	defer server.Close()

// 	// Point the client at the fake server instead of the real API
// 	client := &ArticClient{BaseURL: server.URL, Cache: nil}

// 	art, err := client.ArtworkByID(42)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if art.ArtworkTitle != "Nighthawks" {
// 		t.Errorf("got %q, want %q", art.ArtworkTitle, "Nighthawks")
// 	}
// }

// func TestGeneralSearch(t *testing.T) {
// 	fakeResponse := `{
//         "data": [
//             {"id": 1, "title": "A", "image_id": "img-1", "_score": 5},
//             {"id": 2, "title": "B", "image_id": "img-2", "_score": 0}
//         ]
//     }`

// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, fakeResponse)
// 	}))
// 	defer server.Close()

// 	client := &ArticClient{BaseURL: server.URL, Cache: nil}

// 	result, err := client.GeneralSearch("monet", 10)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	// Score 0 should be filtered out, so only 1 result
// 	if result.ResultsLength != 1 {
// 		t.Errorf("expected 1 result after filtering, got %d", result.ResultsLength)
// 	}
// }
