package routes

import (
	"net/http"

	"real-time-forum/backend/handler"
)

func SetRoutes(handler *handler.AppHandler) {
	http.HandleFunc("/posts", handler.GetPostsHandler)
}
