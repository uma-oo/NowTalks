package routes

import (
	"net/http"

	"real-time-forum/backend/handler"
)

func SetRoutes(handler *handler.AppHandler) {
	http.HandleFunc("/api/post", handler.PostHandler)
	http.Handle("/api/comment", handler)


}
