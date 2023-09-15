package main

import (
	"encoding/json"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/notification/model"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	fmt.Println("starting app")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("success")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	_, err = ch.QueueDeclare(
		"newEpisodeNotificationQueue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
	}

	newEpisodeNotificationMessage1, err := json.Marshal(model.NewEpisodeNotificationMessage{
		ShowID: 1,
		Number: 1,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = ch.Publish(
		"",
		"newEpisodeNotificationQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        newEpisodeNotificationMessage1,
		},
	)

	newEpisodeNotificationMessage2, err := json.Marshal(model.NewEpisodeNotificationMessage{
		ShowID: 1,
		Number: 2,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = ch.Publish(
		"",
		"newEpisodeNotificationQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        newEpisodeNotificationMessage2,
		},
	)

	newEpisodeNotificationMessage3, err := json.Marshal(model.NewEpisodeNotificationMessage{
		ShowID: 1,
		Number: 3,
	})

	if err != nil {
		log.Fatal(err)
	}

	err = ch.Publish(
		"",
		"newEpisodeNotificationQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        newEpisodeNotificationMessage3,
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully published message to queue")
}
