package routes

import (
	"net/http"

	"real-time-forum/backend/handler"
	m "real-time-forum/backend/middleware"
	s "real-time-forum/backend/service"
)

func SetRoutes(
	Phandler *handler.PostHandler,
	Chandler *handler.CommentHandler,
	Rhanlder *handler.ReactionHanlder,
	Uhandler *handler.UserHanlder,
	logout *handler.Logout,
	users *handler.Users,
	loggedin *handler.UserData,
	categories *handler.CategoriesHandler,
	chat *handler.ChatServer,
	messages *handler.MessagesHandler,
	service *s.AppService,
) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/comment", m.NewMiddleWare(Chandler, service))
	mux.Handle("/api/post", m.NewMiddleWare(Phandler, service))
	mux.Handle("/api/user/", m.NewLoginMiddleware(Uhandler, service))
	mux.Handle("/api/react/", m.NewMiddleWare(Rhanlder, service))
	mux.Handle("/api/user/logout", m.NewMiddleWare(logout, service))
	mux.HandleFunc("/api/categories", categories.GetCategories)
	mux.Handle("/api/users", m.NewMiddleWare(users, service))
	mux.Handle("/api/messages", m.NewMiddleWare(messages, service))
	mux.HandleFunc("/api/loggedin", loggedin.GetLoggedIn)
	mux.Handle("/ws/", m.NewMiddleWare(chat, service))
	handlerSPA := http.HandlerFunc(handler.ServeStaticFiles)
	mux.Handle("/", handlerSPA)
	return mux
}

