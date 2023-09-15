package main

import (
	"flag"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/notification/api"
	"github.com/cyjhuiz/netflix-clone/backend/notification/client"
	"github.com/cyjhuiz/netflix-clone/backend/notification/dao"
	"github.com/cyjhuiz/netflix-clone/backend/notification/listener"
	"github.com/cyjhuiz/netflix-clone/backend/notification/service"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	store, err := dao.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()
	if *seed {
		fmt.Println("seeding data")
	}

	notificationDao := dao.NewNotificationDao(store)

	userAPIGRPCClient, err := client.NewUserAPIGRPCClient()
	if err != nil {
		log.Fatal(err)
	}

	showAPIGRPCClient, err := client.NewShowAPIGRPCClient()
	if err != nil {
		log.Fatal(err)
	}

	notificationService := service.NewNotificationService(
		notificationDao,
		showAPIGRPCClient,
		userAPIGRPCClient,
	)

	fmt.Println("Consumer Application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rabbitmqChannel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	_, err = rabbitmqChannel.QueueDeclare(
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

	notificationListener, err := listener.NewNotificationListener(
		rabbitmqChannel,
		notificationService,
	)
	if err != nil {
		log.Fatal(err)
	}

	keepAsyncListenersRunning := make(chan bool)

	go notificationListener.ListenToNewEpisodeNotificationQueue()

	fmt.Println("Successfully connected to rabbitMQ instance")
	fmt.Println("waiting for messages")

	restApiServer := api.NewRESTAPIServer(":3003", notificationService)
	restApiServer.Run()

	<-keepAsyncListenersRunning
}
