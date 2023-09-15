package service

import (
	"context"
	"github.com/cyjhuiz/netflix-clone/backend/show/dao"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
)

type FavouriteService struct {
	FavouriteDao *dao.FavouriteDao
}

func NewFavouriteService(favouriteDao *dao.FavouriteDao) *FavouriteService {
	return &FavouriteService{
		FavouriteDao: favouriteDao,
	}
}

func (favouriteService *FavouriteService) GetFavouritesByShowID(ctx context.Context, showID int64) ([]*model.Favourite, error) {
	favourite, err := favouriteService.FavouriteDao.GetFavouritesByShowID(showID)
	if err != nil {
		return nil, err
	}

	return favourite, nil
}

func (favouriteService *FavouriteService) GetFavouriteByShowIDAndUserID(ctx context.Context, showID int64, userID int64) (*model.Favourite, error) {
	favourite, err := favouriteService.FavouriteDao.GetFavouriteByShowIDAndUserID(showID, userID)
	if err != nil {
		return nil, err
	}

	return favourite, nil
}

func (favouriteService *FavouriteService) CreateFavourite(ctx context.Context, favourite *model.Favourite) error {
	err := favouriteService.FavouriteDao.CreateFavourite(favourite)
	if err != nil {
		return err
	}

	return nil
}

func (favouriteService *FavouriteService) DeleteFavouriteByShowIDAndUserID(ctx context.Context, showID int64, userID int64) error {
	err := favouriteService.FavouriteDao.DeleteFavouriteByShowIDAndUserID(showID, userID)
	if err != nil {
		return err
	}

	return nil
}
