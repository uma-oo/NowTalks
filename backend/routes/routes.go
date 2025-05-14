package routes

import (
	"net/http"

	"real-time-forum/backend/handler"
	m "real-time-forum/backend/middleware"
	s "real-time-forum/backend/service"
)

func SetRoutes(Phandler *handler.PostHandler,
	Chandler *handler.CommentHandler,
	Uhandler *handler.UserHanlder,
	logout *handler.Logout,
	service *s.AppService,
) {
	http.Handle("/api/comment", m.NewMiddleWare(Chandler, service))
	http.Handle("/api/post", m.NewMiddleWare(Phandler, service))
	http.Handle("/api/user/", m.NewLoginMiddleware(Uhandler, service))
	http.Handle("/api/user/logout", m.NewMiddleWare(logout, service))

}
