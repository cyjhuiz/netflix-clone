package api

import (
	"encoding/json"
	"github.com/cyjhuiz/netflix-clone/backend/show/handler/rest"
	"github.com/cyjhuiz/netflix-clone/backend/show/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type RESTAPIServer struct {
	ListenAddr       string
	ShowHandler      *rest.ShowHandler
	EpisodeHandler   *rest.EpisodeHandler
	LikeHandler      *rest.LikeHandler
	FavouriteHandler *rest.FavouriteHandler
}

func NewRESTAPIServer(listenAddr string, showService *service.ShowService, episodeService *service.EpisodeService, likeService *service.LikeService, favouriteService *service.FavouriteService) *RESTAPIServer {
	return &RESTAPIServer{
		ListenAddr:       listenAddr,
		ShowHandler:      rest.NewShowHandler(showService),
		EpisodeHandler:   rest.NewEpisodeHandler(episodeService),
		LikeHandler:      rest.NewLikeHandler(likeService),
		FavouriteHandler: rest.NewFavouriteHandler(favouriteService),
	}
}

func (restAPIServer *RESTAPIServer) Run() {
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.HandleFunc("/show", addTopLevelErrorHandler(restAPIServer.ShowHandler.HandleShows))
	router.HandleFunc("/show/{showID}", addTopLevelErrorHandler(restAPIServer.ShowHandler.HandleShow))

	router.HandleFunc("/show/{showID}/episode/{number}", addTopLevelErrorHandler(restAPIServer.EpisodeHandler.HandleEpisode))

	router.HandleFunc("/show/{showID}/like", addTopLevelErrorHandler(restAPIServer.LikeHandler.HandleLike))

	router.HandleFunc("/show/{showID}/favourite", addTopLevelErrorHandler(restAPIServer.FavouriteHandler.HandleFavourite))

	log.Println("REST API server running on port", restAPIServer.ListenAddr)

	http.ListenAndServe(restAPIServer.ListenAddr, router)
}

type apiHandlerFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func addTopLevelErrorHandler(inputApiHandlerFunc apiHandlerFunc) http.HandlerFunc {
	customApiHandlerFunc := func(writer http.ResponseWriter, request *http.Request) {
		if err := inputApiHandlerFunc(writer, request); err != nil {
			WriteJSON(writer, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}

	return customApiHandlerFunc
}

func WriteJSON(writer http.ResponseWriter, status int, value any) error {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(status)
	return json.NewEncoder(writer).Encode(value)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		writer.Header().Add("Access-Control-Allow-Credentials", "true")
		writer.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		writer.Header().Set("content-type", "application/json;charset=UTF-8")
		if request.Method == "OPTIONS" {
			writer.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(writer, request)
	})
}
