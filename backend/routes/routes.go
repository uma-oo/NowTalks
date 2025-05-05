package routes

import (
	"net/http"

	"real-time-forum/backend/handler"
)

func SetRoutes(Phandler handler.PostHandler, Chandler *handler.CommentHandler) {
	http.Handle("/api/comment", Chandler)
	http.HandleFunc("/api/posts", Phandler.PostHandler)
	fileserver := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fileserver)
}
