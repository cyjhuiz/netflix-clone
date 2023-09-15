package service

import (
	"github.com/cyjhuiz/netflix-clone/backend/show/dao"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
)

type LikeService struct {
	LikeDao *dao.LikeDao
}

func NewLikeService(likeDao *dao.LikeDao) *LikeService {
	return &LikeService{
		LikeDao: likeDao,
	}
}

func (likeService *LikeService) GetLikeByShowIDAndUserID(showID int64, userID int64) (*model.Like, error) {
	like, err := likeService.LikeDao.GetLikeByShowIDAndUserID(showID, userID)
	if err != nil {
		return nil, err
	}

	return like, nil
}

func (likeService *LikeService) CreateLike(like *model.Like) error {
	err := likeService.LikeDao.CreateLike(like)
	if err != nil {
		return err
	}

	return nil
}

func (likeService *LikeService) DeleteLikeByShowIDAndUserID(showID int64, userID int64) error {
	err := likeService.LikeDao.DeleteLikeByShowIDAndUserID(showID, userID)
	if err != nil {
		return err
	}

	return nil
}
