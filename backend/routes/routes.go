package routes

import (
	"io/fs"
	"net/http"
	"os"

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
	mux.Handle("/api/comment", m.NewRateLimitMiddleWare(m.NewMiddleWare(Chandler, service)))
	mux.Handle("/api/post", m.NewRateLimitMiddleWare(m.NewMiddleWare(Phandler, service)))
	mux.Handle("/api/user/", m.NewRateLimitMiddleWare(m.NewLoginMiddleware(Uhandler, service)))
	mux.Handle("/api/react/", m.NewRateLimitMiddleWare(m.NewMiddleWare(Rhanlder, service)))
	mux.Handle("/api/user/logout", m.NewRateLimitMiddleWare(m.NewMiddleWare(logout, service)))
	mux.HandleFunc("/api/categories", categories.GetCategories)
	mux.Handle("/api/users", m.NewRateLimitMiddleWare(m.NewMiddleWare(users, service)))
	mux.Handle("/api/messages", m.NewRateLimitMiddleWare(m.NewMiddleWare(messages, service)))
	mux.HandleFunc("/api/loggedin", loggedin.GetLoggedIn)
	mux.Handle("/ws/", m.NewRateLimitMiddleWare(m.NewMiddleWare(chat, service)))
	var frontend fs.FS = os.DirFS("../frontend")
	httpFS := http.FS(frontend)
	fileServer := http.FileServer(httpFS)
	serveIndex := handler.ServeFileContents("index.html", httpFS)
	mux.Handle("/", handler.Intercept404(fileServer, serveIndex))
	return mux
}

// func handleSPA(w http.ResponseWriter, r *http.Request) {
// 	file_info, err := os.Stat(filepath.Join("../frontend/", r.URL.Path[1:]))
// 	if err != nil || file_info.IsDir() {
// 		http.ServeFile(w, r, "../frontend/index.html")
// 		return
// 	}
// 	http.FileServer(http.Dir("../frontend")).ServeHTTP(w, r)
// }
