package service

import (
	"context"
	"github.com/cyjhuiz/netflix-clone/backend/show/dao"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
)

type EpisodeService struct {
	EpisodeDao *dao.EpisodeDao
}

func NewEpisodeService(episodeDao *dao.EpisodeDao) *EpisodeService {
	return &EpisodeService{
		EpisodeDao: episodeDao,
	}
}

func (episodeService *EpisodeService) GetEpisodeByShowIDAndNumber(ctx context.Context, showID int64, number int64) (*model.EpisodeView, error) {
	episode, err := episodeService.EpisodeDao.GetEpisodeViewByShowIDAndNumber(showID, number)
	if err != nil {
		return nil, err
	}

	return episode, nil
}
