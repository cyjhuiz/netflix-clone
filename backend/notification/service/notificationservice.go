package service

import (
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/notification/client"
	"github.com/cyjhuiz/netflix-clone/backend/notification/dao"
	"github.com/cyjhuiz/netflix-clone/backend/notification/model"
	"log"
	"sync"
	"time"
)

type NotificationService struct {
	NotificationDao   *dao.NotificationDao
	ShowAPIGRPCClient *client.ShowAPIGRPCClient
	UserAPIGRPCClient *client.UserAPIGRPCClient
}

func NewNotificationService(
	notificationDao *dao.NotificationDao,
	showAPIGRPCClient *client.ShowAPIGRPCClient,
	userAPIGRPCClient *client.UserAPIGRPCClient) *NotificationService {
	return &NotificationService{
		NotificationDao:   notificationDao,
		ShowAPIGRPCClient: showAPIGRPCClient,
		UserAPIGRPCClient: userAPIGRPCClient,
	}
}

func (notificationService *NotificationService) GetUserNotificationsByUserID(userID int64) ([]*model.UserNotificationView, error) {
	userNotifications, err := notificationService.NotificationDao.GetUserNotificationsByUserID(userID)
	if err != nil {
		return nil, err
	}

	return userNotifications, nil
}
func (notificationService *NotificationService) SendNewEpisodeNotification(showID int64, number int64) {
	start := time.Now()

	var err error

	// wait group for concurrent calls
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Get users information on users who favourited the show, where Favourite(FavouriteID, ShowID, UserID)
	var userFavourites []*model.Favourite
	go func() {
		userFavourites, err = notificationService.ShowAPIGRPCClient.
			GetUserFavouritesByShowID(showID)
		if err != nil {
			log.Fatal(err)

		}
		wg.Done()
	}()

	// Get episode information
	var episode *model.Episode
	go func() {
		episode, err = notificationService.ShowAPIGRPCClient.
			GetEpisodeByShowIDAndNumber(showID, number)
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()

	// Collate user IDs to gather user information to send app notification
	var userIDs []int64
	for _, userFavourite := range userFavourites {
		userIDs = append(userIDs, userFavourite.UserID)
		fmt.Printf("%+v\n", userFavourite)
	}

	if len(userIDs) == 0 {
		fmt.Println("No userIDs applicable to send new episode notification")
		return
	}

	notificationDescription := fmt.Sprintf("Episode %d of %s just got released", episode.Number, episode.Show.Title)
	notification := model.NewNotification(
		"New Arrival",
		notificationDescription,
		episode.ThumbnailURL,
	)

	err = notificationService.NotificationDao.CreateAppNotification(userIDs, notification)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}

func (notificationService *NotificationService) SendSuccessfulUploadNotification(showID int64, number int64) {
	start := time.Now()

	var err error

	// wait group for concurrent calls
	wg := sync.WaitGroup{}
	wg.Add(2)

	// Get episode information
	var episode *model.Episode
	episode, err = notificationService.ShowAPIGRPCClient.
		GetEpisodeByShowIDAndNumber(showID, number)
	if err != nil {
		log.Fatal(err)
	}

	uploaderUserIDs := []int64{episode.Show.UploaderID}
	fmt.Println("sent app notifications")
	notificationTitle := fmt.Sprintf("Episode %d of %s successfully uploaded", episode.Number, episode.Show.Title)
	notification := model.NewNotification(
		"Successful New Upload",
		notificationTitle,
		episode.ThumbnailURL,
	)

	err = notificationService.NotificationDao.CreateAppNotification(uploaderUserIDs, notification)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(time.Since(start))
}
