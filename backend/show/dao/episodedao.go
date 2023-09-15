package dao

import (
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
	"github.com/cyjhuiz/netflix-clone/backend/show/netflix_clone_db/public/table"
	"github.com/go-jet/jet/v2/postgres"
)

type EpisodeDao struct {
	Store *Store
}

func NewEpisodeDao(store *Store) *EpisodeDao {
	return &EpisodeDao{
		Store: store,
	}
}

func (episodeDao *EpisodeDao) GetEpisodeViewByShowIDAndNumber(showID int64, number int64) (*model.EpisodeView, error) {
	statement := postgres.
		SELECT(
			table.Episode.AllColumns,
			table.Show.AllColumns,
		).
		FROM(
			table.Episode.
				INNER_JOIN(
					table.Show,
					table.Episode.ShowID.EQ(table.Show.ShowID),
				),
		).WHERE(
		table.Episode.ShowID.EQ(postgres.Int(showID)).
			AND(table.Episode.Number.EQ(postgres.Int(number))),
	)
	var episodes []*model.EpisodeView
	err := statement.Query(episodeDao.Store.db, &episodes)
	if err != nil {
		return nil, err
	}

	if len(episodes) == 0 {
		return nil, fmt.Errorf("episode not found")
	}

	return episodes[0], nil
}

func (episodeDao *EpisodeDao) CreateEpisode(episode *model.Episode) error {
	statement := table.Episode.
		INSERT(table.Episode.MutableColumns).
		MODEL(episode)

	_, err := statement.Exec(episodeDao.Store.db)
	if err != nil {
		return err
	}

	return nil
}
