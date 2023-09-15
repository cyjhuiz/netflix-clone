package main

import (
	"encoding/json"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/show/dao"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
	"io/ioutil"
	"log"
)

type RawShow struct {
	Adult            bool     `json:"adult"`
	BackdropPath     string   `json:"backdrop_path"`
	Id               int      `json:"id"`
	Title            string   `json:"title,omitempty"`
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    string   `json:"original_title,omitempty"`
	Overview         string   `json:"overview"`
	PosterPath       string   `json:"poster_path"`
	MediaType        string   `json:"media_type"`
	GenreIds         []int    `json:"genre_ids"`
	Popularity       float64  `json:"popularity"`
	ReleaseDate      string   `json:"release_date,omitempty"`
	Video            bool     `json:"video,omitempty"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Name             string   `json:"name,omitempty"`
	OriginalName     string   `json:"original_name,omitempty"`
	FirstAirDate     string   `json:"first_air_date,omitempty"`
	OriginCountry    []string `json:"origin_country,omitempty"`
}

func main() {
	seedData()
}

func seedData() {
	store, err := dao.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	categoryDao := dao.NewCategoryDao(store)
	showDao := dao.NewShowDao(store)
	episodeDao := dao.NewEpisodeDao(store)

	categoryNames := [...]string{"Netflix Originals", "Trending Now", "Top Rated", "Action Movies", "Comedy Movies", "Horror Movies", "Romance Movies", "Documentaries"}
	for _, categoryName := range categoryNames {
		category := model.NewCategory(categoryName)

		if err := categoryDao.CreateCategory(category); err != nil {
			log.Fatal(err)
		}
	}

	_, err = categoryDao.GetCategories()
	if err != nil {
		log.Fatal(err)
	}

	jsonFilenames := []string{"Netflix Originals", "Trending Now", "Top Rated", "Action Movies", "Comedy Movies", "Horror Movies", "Romance Movies", "Documentaries"}
	for idx, filename := range jsonFilenames {
		// Let's first read the `config.json` file
		content, err := ioutil.ReadFile("./seed/" + filename + ".json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		type ResponseObject struct {
			Page         int64     `json:"page"`
			Results      []RawShow `json:"results"`
			TotalPages   int64     `json:"total_pages"`
			TotalResults int64     `json:"total_results"`
		}
		// Now let's unmarshall the data into `payload`
		var responseObject ResponseObject
		err = json.Unmarshal(content, &responseObject)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		var shows []*model.Show

		for _, show := range responseObject.Results {
			var showType string
			mediaType := show.MediaType
			if mediaType == "movie" {
				showType = "movie"
			} else {
				showType = "series"
			}

			releaseDate := show.ReleaseDate
			if len(releaseDate) == 0 {
				if len(show.FirstAirDate) != 0 {
					releaseDate = show.FirstAirDate
				} else {
					releaseDate = "2022-11-02"
				}
			}

			showTitle := show.Title
			if len(showTitle) == 0 {
				showTitle = show.Name
			}

			show := model.NewShow(
				showTitle,
				show.Overview,
				60,
				showType,
				int64(idx+1),
				"https://image.tmdb.org/t/p/original/"+show.BackdropPath,
				releaseDate+" 15:23:45+02",
				1,
			)

			shows = append(shows, show)
		}

		fmt.Println("%+v\n", shows)

		err = showDao.CreateShows(shows)
		if err != nil {
			log.Fatal(err)
		}
	}

	showViews, err := showDao.GetShowViews()
	if err != nil {
		log.Fatal(err)
	}

	var showID int64
	for _, showView := range showViews {
		if showView.Title == "VGLY" {
			showID = showView.ShowID
			break
		}
	}

	episode1 := model.NewEpisode(
		showID,
		1,
		"The Three Brosketeers",
		"Talented artist Vgly struggles to break into the music business, but an unexpected tragedy might change everything.",
		"https://www.youtube.com/watch?v=zbEeDZWWFsg",
		"https://image.tmdb.org/t/p/original//zzMmUE7aRRtHR9GL90zrtvlqXlm.jpg",
		"2023-05-25 15:23:45+02",
	)
	episode2 := model.NewEpisode(
		showID,
		2,
		"Honor Thy Father",
		"Oliver sets his sights on taking down a criminal with ties to the Chinese triad: Martin Somers, who's being prosecuted by Laurel.",
		"https://www.youtube.com/watch?v=zbEeDZWWFsg",
		"https://image.tmdb.org/t/p/original//zzMmUE7aRRtHR9GL90zrtvlqXlm.jpg",
		"2023-05-25 15:23:45+02",
	)
	episode3 := model.NewEpisode(
		showID,
		3,
		"Yung Jo",
		"Bubble resorts to extreme measures to salvage his reputation; Vgly decides to start a rap beef.",
		"https://www.youtube.com/watch?v=zbEeDZWWFsg",
		"https://image.tmdb.org/t/p/original//zzMmUE7aRRtHR9GL90zrtvlqXlm.jpg",
		"2023-05-25 15:23:45+02",
	)

	episodes := []*model.Episode{episode1, episode2, episode3}
	for _, episode := range episodes {
		err := episodeDao.CreateEpisode(episode)
		if err != nil {
			log.Fatal(err)
		}
	}

}
