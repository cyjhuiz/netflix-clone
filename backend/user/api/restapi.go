package api

import (
	"github.com/cyjhuiz/netflix-clone/backend/user/handler/rest"
	"github.com/cyjhuiz/netflix-clone/backend/user/service"
	"github.com/cyjhuiz/netflix-clone/backend/user/util/authutil"
	"github.com/cyjhuiz/netflix-clone/backend/user/util/generalutil"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

type RESTAPIServer struct {
	ListenAddr  string
	UserHandler *rest.UserHandler
}

func NewRESTAPIServer(listenAddr string, userService *service.UserService) *RESTAPIServer {
	return &RESTAPIServer{
		ListenAddr:  listenAddr,
		UserHandler: rest.NewUserHandler(userService),
	}
}

func (restAPIServer *RESTAPIServer) Run() {
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	router.HandleFunc("/user", addTopLevelErrorHandler(restAPIServer.UserHandler.HandleUsers))

	router.HandleFunc("/user/login", addTopLevelErrorHandler(restAPIServer.UserHandler.HandleUserLogin))
	router.HandleFunc("/user/{userID}", addJWTAuthHandler(addTopLevelErrorHandler(restAPIServer.UserHandler.HandleUser)))

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
			generalutil.WriteJSON(writer, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}

	return customApiHandlerFunc
}

func addJWTAuthHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	protectedHandlerFunc := func(writer http.ResponseWriter, request *http.Request) {

		authHeader := request.Header.Get("Authorization")
		tokenString := strings.Split(authHeader, " ") // tokenString format: "Bearer tokenContent123"
		isValidTokenFormat := len(tokenString) == 2
		if !isValidTokenFormat {
			generalutil.WriteJSON(writer, http.StatusForbidden, ApiError{Error: "access denied. token not found"})
			return
		}

		tokenContent := tokenString[1]
		token, err := authutil.ValidateJWT(tokenContent)
		if err != nil || !token.Valid {
			generalutil.WriteJSON(writer, http.StatusForbidden, ApiError{Error: "invalid token"})
			return
		}

		userIDStr := mux.Vars(request)["userID"]
		isUserRelatedRoute := len(userIDStr) == 0
		if isUserRelatedRoute && !authutil.IsActualUser(token, userIDStr) {
			generalutil.WriteJSON(writer, http.StatusForbidden, ApiError{Error: "unauthorized user"})
			return
		}

		handlerFunc(writer, request)
	}

	return protectedHandlerFunc
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
