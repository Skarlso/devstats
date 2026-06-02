package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Skarlso/devstats/models"
	"github.com/Skarlso/devstats/service"
)

type fakeClient struct {
	err           error
	contributions int
	prs           int
	issues        int
}

func (f *fakeClient) FetchContribute(user *models.User) error {
	if f.err != nil {
		return f.err
	}
	user.Contribution = f.contributions
	user.PRCount = f.prs
	user.IssueCount = f.issues
	return nil
}

func newTestMux(client *fakeClient) *http.ServeMux {
	mux := http.NewServeMux()
	SetupRoutes(mux, service.NewDevStatsServiceWithClient(client))
	return mux
}

func TestScoreHandler(t *testing.T) {
	tests := []struct {
		name       string
		target     string
		client     *fakeClient
		wantStatus int
		wantType   string
		wantBody   []string
	}{
		{
			name:       "renders card for a valid username",
			target:     "/?username=skarlso",
			client:     &fakeClient{contributions: 1234567, prs: 42, issues: 7},
			wantStatus: http.StatusOK,
			wantType:   "image/svg+xml",
			wantBody:   []string{"CNCF DevStats", "1,234,567", "42", "7"},
		},
		{
			name:       "missing username is a bad request",
			target:     "/",
			client:     &fakeClient{},
			wantStatus: http.StatusBadRequest,
			wantBody:   []string{"Missing username"},
		},
		{
			name:       "upstream error is reported",
			target:     "/?username=ghost",
			client:     &fakeClient{err: errors.New("not found")},
			wantStatus: http.StatusBadRequest,
			wantBody:   []string{"Username not found."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tt.target, nil)

			newTestMux(tt.client).ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
			if tt.wantType != "" {
				if got := rec.Header().Get("Content-Type"); got != tt.wantType {
					t.Errorf("Content-Type = %q, want %q", got, tt.wantType)
				}
			}
			body := rec.Body.String()
			for _, want := range tt.wantBody {
				if !strings.Contains(body, want) {
					t.Errorf("body missing %q\n%s", want, body)
				}
			}
		})
	}
}

func TestScoreHandlerSetsCacheHeaders(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/?username=skarlso", nil)

	newTestMux(&fakeClient{}).ServeHTTP(rec, req)

	if got := rec.Header().Get("Cache-Control"); got != "public, max-age=7200" {
		t.Errorf("Cache-Control = %q", got)
	}
	if rec.Header().Get("Expires") == "" {
		t.Error("Expires header not set")
	}
}

func TestHealthHandler(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)

	newTestMux(&fakeClient{}).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if got := rec.Body.String(); got != "OK" {
		t.Errorf("body = %q, want %q", got, "OK")
	}
}

func TestUnknownPathIsNotFound(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/somethingelse", nil)

	newTestMux(&fakeClient{}).ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}
