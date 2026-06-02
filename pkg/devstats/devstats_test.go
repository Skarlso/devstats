package devstats

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Skarlso/devstats/models"
)

func TestFetchContributeSuccess(t *testing.T) {
	var gotReq DevStatsRequest

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(body, &gotReq)
		_ = json.NewEncoder(w).Encode(DevStatsContributionsResponse{
			Contributions: 999,
			Issues:        12,
			PRs:           34,
		})
	}))
	defer srv.Close()

	client := NewDevStats(srv.URL)
	user := &models.User{Username: "skarlso"}

	if err := client.FetchContribute(user); err != nil {
		t.Fatalf("FetchContribute returned error: %v", err)
	}

	if gotReq.API != "GithubIDContributions" {
		t.Errorf("api = %q, want %q", gotReq.API, "GithubIDContributions")
	}
	if gotReq.Payload.GitHubID != "skarlso" {
		t.Errorf("github_id = %q, want %q", gotReq.Payload.GitHubID, "skarlso")
	}
	if user.Contribution != 999 || user.IssueCount != 12 || user.PRCount != 34 {
		t.Errorf("user = %+v, want contributions=999 issues=12 prs=34", user)
	}
}

func TestFetchContributeServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":"boom"}`))
	}))
	defer srv.Close()

	client := NewDevStats(srv.URL)

	if err := client.FetchContribute(&models.User{Username: "skarlso"}); err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestNewDevStatsDefaultsURL(t *testing.T) {
	ds, ok := NewDevStats("").(*DevStats)
	if !ok {
		t.Fatal("NewDevStats did not return *DevStats")
	}
	if ds.DevStatsURL != URL {
		t.Errorf("url = %q, want %q", ds.DevStatsURL, URL)
	}
}
