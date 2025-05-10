package handler

import (
	"fmt"
	"net/http"

	"real-time-forum/backend/models"
)

func (Uhandler *UserHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path url", r.URL.Path)
	r.Header.Set("Content-Type", "application/json")
	switch {
	case r.Method == http.MethodPost && r.URL.Path[1:] == "api/user/login":
		Uhandler.Login(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path[1:] == "api/user/register":
		Uhandler.Register(w, r)
		return
	case r.Method != http.MethodPost:
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not allowed!!"})
		return
	default:
		WriteJsonErrors(w, *models.NewErrorJson(404, "ERROR!! Page Not Found!!"))
		return
	}
}
