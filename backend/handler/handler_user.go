package handler

import (
	"net/http"

	"real-time-forum/backend/models"
)

// add the endpoint of getusers
func (Uhandler *UserHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/api/user/login":
		Uhandler.Login(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/api/user/register":
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
