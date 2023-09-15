package model

import (
	"time"
)

type User struct {
	UserID         int64  `json:"userID"`
	Email          string `json:"email"`
	SubscriptionID int64  `json:"subscriptionID"`
}

type Show struct {
	ShowID       int64  `json:"showID";`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Duration     int64  `json:"duration"`
	ShowType     string `json:"showType"`
	CategoryID   int64  `json:"categoryID"`
	ThumbnailURL string `json:"thumbnailURL"`
	ReleaseDate  string `json:"releaseDate"`
	UploaderID   int64  `json:"uploaderID"`
}

type Episode struct {
	EpisodeID    int64  `json:"episodeID"`
	ShowID       int64  `json:"showID"`
	Number       int64  `json:"number"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	VideoURL     string `json:"videoURL"`
	ThumbnailURL string `json:"thumbnailURL"`
	ReleaseDate  string `json:"releaseDate"`
	Show         *Show  `json:"show"`
}

type Favourite struct {
	FavouriteID int64 `json:"favouriteID"`
	ShowID      int64 `json:"showID"`
	UserID      int64 `json:"userID"`
}

type Notification struct {
	NotificationID int64     `json:"notificationID" sql:"primary_key"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	ThumbnailURL   string    `json:"thumbnailURL"`
	CreatedAt      time.Time `json:"createdAt"`
}

type UserNotification struct {
	UserNotificationID int64 `json:"user_notification_id" sql:"primary_key"`
	NotificationID     int64 `json:"notification_id"`
	UserID             int64 `json:"notification_id"`
}

type UserNotificationView struct {
	*UserNotification
	*Notification
}

func NewNotification(title string, description string, thumbnailURL string) *Notification {
	return &Notification{
		Title:        title,
		Description:  description,
		ThumbnailURL: thumbnailURL,
	}
}

func NewUserNotification(userID int64, notificationID int64) *UserNotification {
	return &UserNotification{
		UserID:         userID,
		NotificationID: notificationID,
	}
}
