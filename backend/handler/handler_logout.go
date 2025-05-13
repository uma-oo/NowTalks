package handler

import (
	"net/http"
	"time"

	"real-time-forum/backend/models"
)

// let's implement the logout
func (logout *Logout) Logout(w http.ResponseWriter, r *http.Request) {
	// delete from the database before
	cookie, _ := r.Cookie("session")
	session, errJson := logout.service.GetUserSessionByTokenEnsureAuth(cookie.Value)
	if errJson != nil {
		WriteJsonErrors(w, *models.NewErrorJson(errJson.Status, errJson.Message))
		return
	}
	if err := logout.service.DeleteSession(session); err != nil {
		WriteJsonErrors(w, *models.NewErrorJson(err.Status, err.Message))
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   "",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
		Path:    "/",
	})

	w.WriteHeader(http.StatusNoContent)
}

func (logout *Logout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.Method != http.MethodPost && r.URL.Path == "/api/user/logout":
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!"})
		return
	case r.Method == http.MethodPost:
		logout.Logout(w, r)
		return
	default:
		WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "ERROR!! Page Not Found!"})
		return

	}
}
