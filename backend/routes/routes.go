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
	// var frontend fs.FS = os.DirFS("../frontend")
	// httpFS := http.FS(frontend)
	// fileServer := http.FileServer(httpFS)
	// serveIndex := handler.ServeFileContents("index.html", httpFS)
	// mux.Handle("/", handler.Intercept404(fileServer, serveIndex))
	handlerSPA := http.HandlerFunc(m.ServeStaticFiles)
	mux.Handle("/", handlerSPA)
	return mux
}

// func handleSPA(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("inside the handleSPA", r.URL.Path[1:])
// 	file_info, err := os.Stat(filepath.Join("../frontend/", r.URL.Path[1:]))
// 	if err != nil || file_info.IsDir() {
// 		fmt.Printf("err: %v\n", err)
// 		http.ServeFile(w, r, "../frontend/index.html")
// 		return
// 	} else {
// 		fmt.Printf("file_info: %v\n", file_info.Name())
// 		http.ServeFile(w, r, file_info.Name())
// 		http.FileServer(http.Dir("../frontend")).ServeHTTP(w, r)
// 	}
// }
