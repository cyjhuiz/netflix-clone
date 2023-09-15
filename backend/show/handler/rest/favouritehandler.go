package rest

import (
	"context"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
	"github.com/cyjhuiz/netflix-clone/backend/show/service"
	"github.com/gorilla/mux"

	"net/http"
	"strconv"
)

type FavouriteHandler struct {
	FavouriteService *service.FavouriteService
}

func NewFavouriteHandler(favouriteService *service.FavouriteService) *FavouriteHandler {
	return &FavouriteHandler{
		FavouriteService: favouriteService,
	}
}

func (favouriteHandler *FavouriteHandler) HandleFavourite(writer http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "GET":
		if len(request.URL.Query()) == 0 {
			return favouriteHandler.handleGetFavouritesByShowID(writer, request)
		} else {
			return favouriteHandler.handleGetFavouriteByShowIDAndUserID(writer, request)
		}
	case "POST":
		return favouriteHandler.handleCreateFavourite(writer, request)
	case "DELETE":
		return favouriteHandler.handleDeleteFavouriteByShowIDAndUserID(writer, request)
	}

	return fmt.Errorf("method not allowed %s", request.Method)
}

func (favouriteHandler *FavouriteHandler) handleGetFavouritesByShowID(writer http.ResponseWriter, request *http.Request) error {
	ctx := context.Background()

	showIDStr := mux.Vars(request)["showID"]
	showID, err := strconv.ParseInt(showIDStr, 10, 64)
	if err != nil {
		return err
	}

	favourites, err := favouriteHandler.FavouriteService.GetFavouritesByShowID(ctx, showID)
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, favourites)
}

func (favouriteHandler *FavouriteHandler) handleGetFavouriteByShowIDAndUserID(writer http.ResponseWriter, request *http.Request) error {
	ctx := context.Background()

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

	favourite, err := favouriteHandler.FavouriteService.GetFavouriteByShowIDAndUserID(ctx, showID, userID)
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, favourite)
}

func (favouriteHandler *FavouriteHandler) handleCreateFavourite(writer http.ResponseWriter, request *http.Request) error {
	ctx := context.Background()

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

	favourite := model.NewFavourite(
		showID,
		userID,
	)

	err = favouriteHandler.FavouriteService.CreateFavourite(ctx, favourite)
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, map[string]string{"message": "success"})
}

func (favouriteHandler *FavouriteHandler) handleDeleteFavouriteByShowIDAndUserID(writer http.ResponseWriter, request *http.Request) error {
	ctx := context.Background()

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

	err = favouriteHandler.FavouriteService.DeleteFavouriteByShowIDAndUserID(ctx, showID, userID)
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, map[string]string{"message": "success"})
}
