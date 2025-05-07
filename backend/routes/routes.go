package routes

import (
	"net/http"

	"real-time-forum/backend/handler"
)

func SetRoutes(Phandler *handler.PostHandler, Chandler *handler.CommentHandler, Uhandler *handler.UserHanlder) {
	http.Handle("/api/comment", Chandler)
	http.Handle("/api/post", Phandler)
	http.HandleFunc("/api/register", Uhandler.Register)
	fileserver := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fileserver)
}
