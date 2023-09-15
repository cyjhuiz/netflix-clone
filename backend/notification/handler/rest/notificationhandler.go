package rest

import (
	"encoding/json"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/notification/service"
	"net/http"
	"strconv"
)

type NotificationHandler struct {
	NotificationService *service.NotificationService
}

func NewNotificationHandler(notificationService *service.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		NotificationService: notificationService,
	}
}

func (notificationHandler *NotificationHandler) HandleNotifications(writer http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "GET":
		return notificationHandler.handleGetNotifications(writer, request)
	}

	return fmt.Errorf("method not allowed %s", request.Method)

}

func (notificationHandler *NotificationHandler) handleGetNotifications(writer http.ResponseWriter, request *http.Request) error {
	userIDStr := request.URL.Query().Get("userID")
	if len(userIDStr) == 0 {
		return fmt.Errorf("please key in a userID")
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return err
	}

	userNotifications, err := notificationHandler.NotificationService.GetUserNotificationsByUserID(userID)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, userNotifications)
}

func WriteJSON(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(value)
}
