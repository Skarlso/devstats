package devstats

type DevStatsRequest struct {
	API     string          `json:"api"`
	Payload DevStatsPayload `json:"payload"`
}

type DevStatsPayload struct {
	GitHubID string `json:"github_id"`
}

type DevStatsContributionsResponse struct {
	Contributions int `json:"contributions"`
	Issues        int `json:"issues"`
	PRs           int `json:"prs"`
}
