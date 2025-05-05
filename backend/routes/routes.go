package routes

import (
	"net/http"

	"real-time-forum/backend/handler"
)

func SetRoutes(handler *handler.AppHandler) {
	fileserver := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fileserver)
	http.HandleFunc("/posts", handler.GetPostsHandler)
}
