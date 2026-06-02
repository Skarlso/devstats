package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/Skarlso/devstats/models"
	"github.com/Skarlso/devstats/service"
	"github.com/Skarlso/devstats/svg"
)

func SetupRoutes(mux *http.ServeMux, devStatsService *service.DevStatsService) {
	mux.HandleFunc("GET /{$}", ScoreHandler(devStatsService))
	mux.HandleFunc("GET /health", HealthHandler)
}

func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func ScoreHandler(devStatsService *service.DevStatsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		githubID := r.URL.Query().Get("username")
		if githubID == "" {
			http.Error(w, "Missing username", http.StatusBadRequest)
			return
		}

		githubID = strings.ToLower(githubID)

		user, err := devStatsService.GetUserStats(githubID)
		if err != nil {
			http.Error(w, "Username not found.", http.StatusBadRequest)
			return
		}

		card := svg.GenerateSVG(models.CardData{
			Score:      user.Contribution,
			PRs:        user.PRCount,
			Issues:     user.IssueCount,
			TitleColor: "#0086FF",
			TextColor:  "#555555",
			Radius:     10,
		})

		w.Header().Set("Cache-Control", "public, max-age=7200")
		w.Header().Set("Expires", time.Now().Add(2*time.Hour).Format(time.RFC1123))
		w.Header().Set("Content-Type", "image/svg+xml")
		_, _ = w.Write([]byte(card))
	}
}
