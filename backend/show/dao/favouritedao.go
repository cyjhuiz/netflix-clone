package dao

import (
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
	"github.com/cyjhuiz/netflix-clone/backend/show/netflix_clone_db/public/table"
	"github.com/go-jet/jet/v2/postgres"
)

type FavouriteDao struct {
	Store *Store
}

func NewFavouriteDao(store *Store) *FavouriteDao {
	return &FavouriteDao{
		Store: store,
	}
}

func (favouriteDao *FavouriteDao) GetFavouritesByShowID(showID int64) ([]*model.Favourite, error) {
	statement := postgres.
		SELECT(
			table.Favourite.AllColumns,
		).
		FROM(
			table.Favourite,
		).
		WHERE(
			table.Favourite.ShowID.EQ(postgres.Int(showID)),
		)

	var favourites []*model.Favourite
	err := statement.Query(favouriteDao.Store.db, &favourites)
	if err != nil {
		return nil, err
	}

	return favourites, nil
}

func (favouriteDao *FavouriteDao) GetFavouriteByShowIDAndUserID(showID int64, userID int64) (*model.Favourite, error) {
	statement := postgres.
		SELECT(
			table.Favourite.AllColumns,
		).
		FROM(
			table.Favourite,
		).
		WHERE(
			table.Favourite.ShowID.EQ(postgres.Int(showID)).
				AND(table.Favourite.UserID.EQ(postgres.Int(userID))),
		)

	var favourites []*model.Favourite
	err := statement.Query(favouriteDao.Store.db, &favourites)
	if err != nil {
		return nil, err
	}

	if len(favourites) == 0 {
		return nil, nil
	}

	return favourites[0], nil
}

func (favouriteDao *FavouriteDao) CreateFavourite(favourite *model.Favourite) error {
	statement := table.Favourite.
		INSERT(table.Favourite.MutableColumns).
		MODEL(favourite)

	_, err := statement.Exec(favouriteDao.Store.db)
	if err != nil {
		return err
	}

	return nil
}

func (favouriteDao *FavouriteDao) DeleteFavouriteByShowIDAndUserID(showID int64, userID int64) error {
	statement := table.Favourite.
		DELETE().
		WHERE(
			table.Favourite.ShowID.EQ(postgres.Int(showID)).
				AND(table.Favourite.UserID.EQ(postgres.Int(userID))),
		)

	_, err := statement.Exec(favouriteDao.Store.db)
	if err != nil {
		return err
	}

	return nil
}
