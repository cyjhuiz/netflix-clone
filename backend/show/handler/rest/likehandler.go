package rest

import (
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
	"github.com/cyjhuiz/netflix-clone/backend/show/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type LikeHandler struct {
	LikeService *service.LikeService
}

func NewLikeHandler(likeService *service.LikeService) *LikeHandler {
	return &LikeHandler{
		LikeService: likeService,
	}
}

func (likeHandler *LikeHandler) HandleLike(writer http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "GET":
		return likeHandler.handleGetLikeByShowIDAndUserID(writer, request)
	case "POST":
		return likeHandler.handleCreateLike(writer, request)
	case "DELETE":
		return likeHandler.handleDeleteLikeByShowIDAndUserID(writer, request)
	}

	return fmt.Errorf("method not allowed %s", request.Method)
}

func (likeHandler *LikeHandler) handleGetLikeByShowIDAndUserID(writer http.ResponseWriter, request *http.Request) error {
	showIDStr := mux.Vars(request)["showID"]
	showID, err := strconv.ParseInt(showIDStr, 10, 64)
	if err != nil {
		return err
	}

	userIDStr := request.URL.Query().Get("userID")
	if len(userIDStr) == 0 {
		return fmt.Errorf("please key in a userID")
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return err
	}

	like, err := likeHandler.LikeService.GetLikeByShowIDAndUserID(showID, userID)
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, like)
}

func (likeHandler *LikeHandler) handleCreateLike(writer http.ResponseWriter, request *http.Request) error {
	showIDStr := mux.Vars(request)["showID"]
	showID, err := strconv.ParseInt(showIDStr, 10, 64)
	if err != nil {
		return err
	}

	userIDStr := request.URL.Query().Get("userID")
	if len(userIDStr) == 0 {
		return fmt.Errorf("please key in a userID")
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return err
	}

	like := model.NewLike(
		showID,
		userID,
	)

	err = likeHandler.LikeService.CreateLike(like)
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, map[string]string{"message": "success"})
}

func (likeHandler *LikeHandler) handleDeleteLikeByShowIDAndUserID(writer http.ResponseWriter, request *http.Request) error {
	showIDStr := mux.Vars(request)["showID"]
	showID, err := strconv.ParseInt(showIDStr, 10, 64)
	if err != nil {
		return err
	}

	userIDStr := request.URL.Query().Get("userID")
	if len(userIDStr) == 0 {
		return fmt.Errorf("please key in a userID")
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return err
	}

	err = likeHandler.LikeService.DeleteLikeByShowIDAndUserID(showID, userID)
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, map[string]string{"message": "success"})
}
