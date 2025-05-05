package routes

import (
	"net/http"

	"real-time-forum/backend/handler"
)

func SetRoutes(handler *handler.AppHandler) {
	http.HandleFunc("/api/post", handler.PostHandler)
	http.Handle("/api/comment", handler)


	fileserver := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fileserver)
	http.HandleFunc("/posts", handler.GetPostsHandler)
}
