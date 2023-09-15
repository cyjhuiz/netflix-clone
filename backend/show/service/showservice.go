package service

import (
	"encoding/json"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/show/dao"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
	"github.com/redis/go-redis/v9"
)

type ShowService struct {
	ShowDao  *dao.ShowDao
	RedisDao *dao.RedisDao
}

func NewShowService(showDao *dao.ShowDao, redisDao *dao.RedisDao) *ShowService {
	return &ShowService{
		ShowDao:  showDao,
		RedisDao: redisDao,
	}
}

func (showService *ShowService) GetShows() ([]*model.ShowViewConcise, error) {
	shows, err := showService.ShowDao.GetShowViews()
	if err != nil {
		return nil, err
	}

	return shows, nil
}

func (showService *ShowService) GetShowsByCategory(category string) ([]*model.ShowViewConcise, error) {
	var shows []*model.ShowViewConcise

	cacheIdentifier := fmt.Sprintf("GetShowsByCategory=%s", category)
	cachedShowStr, err := showService.RedisDao.Get(cacheIdentifier)

	// if isCached, return result
	isCached := err != redis.Nil && err == nil
	if isCached {
		err = json.Unmarshal([]byte(cachedShowStr), &shows)
		if err != nil {
			return nil, err
		}

		return shows, nil
	} else if err != redis.Nil && err != nil {
		return nil, err
	}

	// Get show by category
	shows, err = showService.ShowDao.GetShowViewsByCategory(category)
	if err != nil {
		return nil, err
	}

	// cache the result before returning it
	cachedShowByte, err := json.Marshal(shows)
	if err != nil {
		return nil, err
	}
	cachedShowStr = string(cachedShowByte)

	err = showService.RedisDao.Set(cacheIdentifier, cachedShowStr)
	if err != nil {
		return nil, err
	}

	return shows, nil
}

func (showService *ShowService) GetShowByShowID(showID int64) (*model.ShowView, error) {
	show, err := showService.ShowDao.GetShowViewByShowID(showID)
	if err != nil {
		return nil, err
	}

	return show, nil
}

func (showService *ShowService) CreateShow(show *model.Show) error {
	err := showService.ShowDao.CreateShow(show)
	if err != nil {
		return err
	}
	return nil
}
