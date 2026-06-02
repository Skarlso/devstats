package service

import (
	"github.com/Skarlso/devstats/models"

	"github.com/Skarlso/devstats/pkg/devstats"
)

type DevStatsService struct {
	client devstats.DevStatsInterface
}

func NewDevStatsService() *DevStatsService {
	return &DevStatsService{
		client: devstats.NewDevStats(""),
	}
}

func NewDevStatsServiceWithClient(client devstats.DevStatsInterface) *DevStatsService {
	return &DevStatsService{client: client}
}

func (s *DevStatsService) GetUserStats(username string) (*models.User, error) {
	user := &models.User{
		Username: username,
	}

	if err := s.client.FetchContribute(user); err != nil {
		return nil, err
	}

	return user, nil
}
