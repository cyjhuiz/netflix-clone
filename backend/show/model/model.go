package model

type Category struct {
	CategoryId int64  `json:"categoryID" sql:"primary_key"`
	Name       string `json:"name"`
}

func NewCategory(name string) *Category {
	return &Category{
		Name: name,
	}
}

type Show struct {
	ShowID       int64  `json:"showID" sql:"primary_key"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Duration     int64  `json:"duration"`
	ShowType     string `json:"showType"`
	CategoryID   int64  `json:"categoryID"`
	ThumbnailURL string `json:"thumbnailURL"`
	ReleaseDate  string `json:"releaseDate"`
	UploaderID   int64  `json:"uploaderID"`
}

type ShowViewConcise struct {
	*Show
	Category Category `json:"category"`
}

type ShowView struct {
	*Show
	Category Category   `json:"category"`
	Episodes []*Episode `json:"episodes"`
}

func NewShow(title string, description string, duration int64, showType string, categoryID int64, thumbnailURL string, releaseDate string, uploaderID int64) *Show {
	return &Show{
		Title:        title,
		Description:  description,
		Duration:     duration,
		ShowType:     showType,
		CategoryID:   categoryID,
		ThumbnailURL: thumbnailURL,
		ReleaseDate:  releaseDate,
		UploaderID:   uploaderID,
	}
}

type Episode struct {
	EpisodeID    int64  `json:"episodeID" sql:"primary_key"`
	ShowID       int64  `json:"showID"`
	Number       int64  `json:"number"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	VideoURL     string `json:"videoURL"`
	ThumbnailURL string `json:"thumbnailURL"`
	ReleaseDate  string `json:"releaseDate"`
}

type EpisodeView struct {
	*Episode
	Show *Show `json:"show"`
}

func NewEpisode(showID int64, number int64, title string, description string, videoURL string, thumbnailURL string, releaseDate string) *Episode {
	return &Episode{
		ShowID:       showID,
		Number:       number,
		Title:        title,
		Description:  description,
		VideoURL:     videoURL,
		ThumbnailURL: thumbnailURL,
		ReleaseDate:  releaseDate,
	}
}

type Like struct {
	LikeID int64 `json:"likeID" sql:"primary_key"`
	ShowID int64 `json:"showID"`
	UserID int64 `json:"userID"`
}

func NewLike(showID int64, userID int64) *Like {
	return &Like{
		ShowID: showID,
		UserID: userID,
	}
}

type Favourite struct {
	FavouriteID int64 `json:"favouriteID" sql:"primary_key"`
	ShowID      int64 `json:"showID"`
	UserID      int64 `json:"userID"`
}

func NewFavourite(showID int64, userID int64) *Favourite {
	return &Favourite{
		ShowID: showID,
		UserID: userID,
	}
}
