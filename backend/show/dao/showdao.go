package dao

import (
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
	"github.com/cyjhuiz/netflix-clone/backend/show/netflix_clone_db/public/table"
	"github.com/go-jet/jet/v2/postgres"
	_ "github.com/lib/pq"
)

type ShowDao struct {
	Store *Store
}

func NewShowDao(store *Store) *ShowDao {
	return &ShowDao{
		Store: store,
	}
}

func (showDao *ShowDao) GetShowViews() ([]*model.ShowViewConcise, error) {
	statement := postgres.
		SELECT(
			table.Show.AllColumns,
			table.Category.AllColumns,
		).
		FROM(
			table.Show.
				INNER_JOIN(table.Category, table.Show.CategoryID.EQ(table.Category.CategoryID)),
		).
		ORDER_BY(
			table.Show.ReleaseDate.DESC(),
		)

	var showViews []*model.ShowViewConcise
	err := statement.Query(showDao.Store.db, &showViews)
	if err != nil {
		return nil, err
	}

	return showViews, nil
}

func (showDao *ShowDao) GetShowViewsByCategory(category string) ([]*model.ShowViewConcise, error) {
	statement := postgres.
		SELECT(
			table.Show.AllColumns,
			table.Category.AllColumns,
		).
		FROM(
			table.Show.
				INNER_JOIN(table.Category, table.Show.CategoryID.EQ(table.Category.CategoryID)),
		).
		WHERE(
			table.Category.Name.EQ(postgres.String(category)),
		).
		ORDER_BY(
			table.Show.ReleaseDate.DESC(),
		)

	var showViews []*model.ShowViewConcise
	err := statement.Query(showDao.Store.db, &showViews)
	if err != nil {
		return nil, err
	}

	return showViews, nil
}

func (showDao *ShowDao) GetShowViewByShowID(showID int64) (*model.ShowView, error) {
	statement := postgres.
		SELECT(
			table.Show.AllColumns,
			table.Category.AllColumns,
			table.Episode.AllColumns,
		).
		FROM(
			table.Show.
				INNER_JOIN(table.Category, table.Show.CategoryID.EQ(table.Category.CategoryID)).
				LEFT_JOIN(table.Episode, table.Show.ShowID.EQ(table.Episode.ShowID)),
		).
		WHERE(
			table.Show.ShowID.EQ(postgres.Int(showID)),
		).
		ORDER_BY(
			table.Show.ReleaseDate.DESC(),
			table.Episode.Number.ASC(),
		)

	var showViews []*model.ShowView

	err := statement.Query(showDao.Store.db, &showViews)
	if err != nil {
		return nil, err
	}

	if len(showViews) == 0 {
		return nil, fmt.Errorf("show not found")
	}
	
	return showViews[0], nil
}

func (showDao *ShowDao) CreateShow(show *model.Show) error {
	statement := table.Show.INSERT(table.Show.MutableColumns).MODEL(show)

	_, err := statement.Exec(showDao.Store.db)
	if err != nil {
		return err
	}

	return nil
}

func (showDao *ShowDao) CreateShows(shows []*model.Show) error {
	statement := table.Show.INSERT(table.Show.MutableColumns).MODELS(shows)

	_, err := statement.Exec(showDao.Store.db)
	if err != nil {
		return err
	}

	return nil
}
