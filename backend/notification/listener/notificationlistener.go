package listener

import (
	"encoding/json"
	"github.com/cyjhuiz/netflix-clone/backend/notification/model"
	"github.com/cyjhuiz/netflix-clone/backend/notification/service"
	"github.com/streadway/amqp"
	"log"
)

type NotificationListener struct {
	NewEpisodeNotificationMessages <-chan amqp.Delivery
	NotificationService            *service.NotificationService
}

func NewNotificationListener(rabbitmqChannel *amqp.Channel, notificationService *service.NotificationService) (*NotificationListener, error) {
	newEpisodeNotificationMessages, err := rabbitmqChannel.Consume(
		"newEpisodeNotificationQueue",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &NotificationListener{
		NewEpisodeNotificationMessages: newEpisodeNotificationMessages,
		NotificationService:            notificationService,
	}, nil
}

func (notificationListener *NotificationListener) ListenToNewEpisodeNotificationQueue() {

	for message := range notificationListener.NewEpisodeNotificationMessages {
		newEpisodeNotificationMessage := model.NewEpisodeNotificationMessage{}
		err := json.Unmarshal(message.Body, &newEpisodeNotificationMessage)
		if err != nil {
			log.Fatal(err)
		}

		showID := newEpisodeNotificationMessage.ShowID
		number := newEpisodeNotificationMessage.Number

		notificationListener.NotificationService.
			SendNewEpisodeNotification(showID, number)

		notificationListener.NotificationService.
			SendSuccessfulUploadNotification(showID, number)
		message.Ack(true)
	}

}
