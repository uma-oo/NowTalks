package handler

import (
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

	userData, errJson := d.service.IsLoggedInUser(cookie.Value)

	if errJson != nil {
		user_data.IsLoggedIn = false
		WriteDataBack(w, user_data)
		return
	}
	user_data.IsLoggedIn = true
	user_data.Id = userData.Id
	user_data.Nickname = userData.Nickname
	WriteDataBack(w, user_data)
}
