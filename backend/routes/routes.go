package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"real-time-forum/backend/handler"
	m "real-time-forum/backend/middleware"
	s "real-time-forum/backend/service"
)

func SetRoutes(Phandler *handler.PostHandler,
	Chandler *handler.CommentHandler,
	Uhandler *handler.UserHanlder,
	service *s.AppService,
	logout *handler.Logout,
) {
	http.Handle("/api/comment", m.NewMiddleWare(Chandler, service))
	http.Handle("/api/post", m.NewMiddleWare(Phandler, service))
	http.Handle("/api/user/", m.NewLoginMiddleware(Uhandler, service))
	http.Handle("/api/user/logout", m.NewMiddleWare(logout, service))
	http.HandleFunc("/", handleSPA)

}

func handleSPA(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	file_info, err := os.Stat(filepath.Join("../frontend/", r.URL.Path[1:]))
	if err != nil || file_info.IsDir() {
		http.ServeFile(w, r, "../frontend/index.html")
		return
	}
	http.FileServer(http.Dir("../frontend")).ServeHTTP(w, r)
}
