package rest

import (
	"context"
	"fmt"
	"github.com/cyjhuiz/netflix-clone/backend/show/service"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type EpisodeHandler struct {
	EpisodeService *service.EpisodeService
}

func NewEpisodeHandler(episodeService *service.EpisodeService) *EpisodeHandler {
	return &EpisodeHandler{
		EpisodeService: episodeService,
	}
}

func (episodeHandler *EpisodeHandler) HandleEpisode(writer http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "GET":
		return episodeHandler.handleGetEpisodeByShowIDAndNumber(writer, request)
	}

	return fmt.Errorf("method not allowed %s", request.Method)
}

func (episodeHandler *EpisodeHandler) handleGetEpisodeByShowIDAndNumber(writer http.ResponseWriter, request *http.Request) error {
	ctx := context.Background()

	showIDStr := mux.Vars(request)["showID"]
	showID, err := strconv.ParseInt(showIDStr, 10, 64)
	if err != nil {
		return err
	}

	numberStr := mux.Vars(request)["number"]
	number, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		return err
	}

	episode, err := episodeHandler.EpisodeService.GetEpisodeByShowIDAndNumber(ctx, showID, number)
	if err != nil {
		return err
	}

	return WriteJSON(writer, http.StatusOK, episode)
}
