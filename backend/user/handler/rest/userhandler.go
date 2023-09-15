package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/user/model"
	"github.com/cyjhuiz/netflix-clone/backend/user/service"
	"github.com/cyjhuiz/netflix-clone/backend/user/util/generalutil"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
func (userHandler *UserHandler) HandleUsers(writer http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "GET":
		return userHandler.handleGetUsers(writer, request)
	case "POST":
		return userHandler.handleCreateUser(writer, request)
	}

	return fmt.Errorf("method not allowed %s", request.Method)
}

func (userHandler *UserHandler) handleGetUsers(writer http.ResponseWriter, request *http.Request) error {
	ctx := context.Background()

	var users []*model.UserView
	var err error

	email := request.URL.Query().Get("email")
	userIDStrs := request.URL.Query().Get("userIDs")

	if len(email) != 0 {
		users, err = userHandler.UserService.GetUsersByEmail(ctx, email)
		if err != nil {
			return err
		}
	} else if len(userIDStrs) != 0 {
		userIDs, err := getUserIDsFromString(userIDStrs)
		if err != nil {
			return err
		}

		users, err = userHandler.UserService.GetUsersByUserIDs(ctx, userIDs)
		if err != nil {
			return err
		}
	} else {
		users, err = userHandler.UserService.GetUsers(ctx)
	}

	return generalutil.WriteJSON(writer, http.StatusOK, users)
}

func (userHandler *UserHandler) handleCreateUser(writer http.ResponseWriter, request *http.Request) error {
	createUserRequest := new(model.CreateUserRequest)

	err := json.NewDecoder(request.Body).Decode(createUserRequest)
	if err != nil {
		return err
	}

	user := model.NewUser(
		createUserRequest.Email,
		createUserRequest.Password,
	)
	err = userHandler.UserService.CreateUser(context.Background(), user)
	if err != nil {
		return err
	}

	return generalutil.WriteJSON(writer, http.StatusOK, user)
}

func (userHandler *UserHandler) HandleUser(writer http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "GET":
		return userHandler.handleGetUserByUserID(writer, request)
	}

	return fmt.Errorf("method not allowed %s", request.Method)
}

func (userHandler *UserHandler) handleGetUserByUserID(writer http.ResponseWriter, request *http.Request) error {
	userIDStr := mux.Vars(request)["userID"]
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return err
	}

	user, err := userHandler.UserService.GetUserByUserID(context.Background(), userID)
	if err != nil {
		return err
	}

	return generalutil.WriteJSON(writer, http.StatusOK, user)
}

func (userHandler *UserHandler) HandleUserLogin(writer http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "POST":
		return userHandler.handleUserLogin(writer, request)
	}

	return fmt.Errorf("method not allowed %s", request.Method)
}

func (userHandler *UserHandler) handleUserLogin(writer http.ResponseWriter, request *http.Request) error {
	loginUserRequest := new(model.LoginUserRequest)

	err := json.NewDecoder(request.Body).Decode(loginUserRequest)
	if err != nil {
		return err
	}

	loginDetails, err := userHandler.UserService.LoginUser(
		context.Background(),
		loginUserRequest.Email,
		loginUserRequest.Password,
	)
	if err != nil {
		return err
	}

	return generalutil.WriteJSON(writer, http.StatusOK, loginDetails)
}

// util functions
func getUserIDsFromString(userIDStrs string) ([]int64, error) {
	var userIDs []int64
	for _, userIDStr := range strings.Split(userIDStrs, ",") {
		userID, err := strconv.ParseInt(userIDStr, 10, 64)

		if err != nil {
			return nil, err
		}

		userIDs = append(userIDs, userID)
	}

	return userIDs, nil
}
