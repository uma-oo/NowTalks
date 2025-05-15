package handler

import (
	"fmt"
	"net/http"

	"real-time-forum/backend/models"
)

func (d *UserData) GetLoggedIn(w http.ResponseWriter, r *http.Request) {
	user_data := &models.UserData{}
	if r.Method != http.MethodGet {
		WriteJsonErrors(w, *models.NewErrorJson(405, "ERROR!! Method Not Allowed!!"))
		return
	}

	cookie, err := r.Cookie("session")
	if err != nil {
		user_data.IsLoggedIn = false
		WriteDataBack(w, user_data)
		return
	}

	user_data, errJson := d.service.IsLoggedInUser(cookie.Value)
	fmt.Println("user data", user_data)
	if errJson != nil {
		user_data.IsLoggedIn = false
		WriteDataBack(w, user_data)
		return
	}
	user_data.IsLoggedIn = true
	WriteDataBack(w, user_data)
}
