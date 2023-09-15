package dao

import (
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
	"github.com/cyjhuiz/netflix-clone/backend/show/netflix_clone_db/public/table"
	"github.com/go-jet/jet/v2/postgres"
)

type LikeDao struct {
	Store *Store
}

func NewLikeDao(store *Store) *LikeDao {
	return &LikeDao{
		Store: store,
	}
}

func (likeDao *LikeDao) GetLikeByShowIDAndUserID(showID int64, userID int64) (*model.Like, error) {
	statement := postgres.RawStatement(
		`
		SELECT "like".like_id AS "like.like_id", "like".show_id AS "like.show_id", "like".user_id AS "like.user_id"
		FROM public."like" 
		WHERE "like".show_id = #1 AND "like".user_id = #2`,
		postgres.RawArgs{"#1": showID, "#2": userID},
	)

	var likes []*model.Like

	err := statement.Query(likeDao.Store.db, &likes)
	if err != nil {
		return nil, err
	}

	if len(likes) == 0 {
		return nil, nil
	}
	return likes[0], nil
}

func (likeDao *LikeDao) CreateLike(like *model.Like) error {
	statement := table.Like.
		INSERT(table.Like.MutableColumns).
		MODEL(like)

	_, err := statement.Exec(likeDao.Store.db)
	if err != nil {
		return err
	}

	return nil
}

func (likeDao *LikeDao) DeleteLikeByShowIDAndUserID(showID int64, userID int64) error {
	statement := postgres.RawStatement(
		`
		DELETE
		FROM public."like" 
		WHERE "like".show_id = #1 AND "like".user_id = #2`,
		postgres.RawArgs{"#1": showID, "#2": userID},
	)

	_, err := statement.Exec(likeDao.Store.db)
	if err != nil {
		return err
	}

	return nil
}
