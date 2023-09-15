package rest

import (
	"encoding/json"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/show/model"
	"github.com/cyjhuiz/netflix-clone/backend/show/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ShowHandler struct {
	ShowService *service.ShowService
}

func NewShowHandler(showService *service.ShowService) *ShowHandler {
	return &ShowHandler{
		ShowService: showService,
	}
}

func (showHandler *ShowHandler) HandleShows(writer http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "GET":
		return showHandler.handleGetShows(writer, request)
	}

	return fmt.Errorf("method not allowed %s", request.Method)

}

func (showHandler *ShowHandler) handleGetShows(writer http.ResponseWriter, request *http.Request) error {
	category := request.URL.Query().Get("category")

	var shows []*model.ShowViewConcise
	var err error

	hasCategory := len(category) > 0
	if hasCategory {
		shows, err = showHandler.ShowService.GetShowsByCategory(category)
	} else {
		shows, err = showHandler.ShowService.GetShows()
	}

	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, shows)
}

func WriteJSON(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(value)
}

func (showHandler *ShowHandler) HandleShow(writer http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "GET":
		return showHandler.handleGetShowByID(writer, request)
	}

	return fmt.Errorf("method not allowed %s", request.Method)

}

func (showHandler *ShowHandler) handleGetShowByID(writer http.ResponseWriter, request *http.Request) error {
	showIDStr := mux.Vars(request)["showID"]
	showID, err := strconv.ParseInt(showIDStr, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid id given %s", showID)
	}

	shows, err := showHandler.ShowService.GetShowByShowID(showID)
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, shows)
}
