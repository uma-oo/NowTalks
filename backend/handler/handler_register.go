package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"real-time-forum/backend/models"
)

func (Uhandler *UserHanlder) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		if err == io.EOF {
			WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: &models.RegisterError{
				Nickname:      " empty username field",
				Age:           " empty username field",
				Gender:        " empty gender field",
				FirstName:     " empty first Name field",
				LastName:      " empty lastName field",
				Email:         " empty email field",
				Password:      " empty password  field",
				VerifPassword: " empty verification Password  field",
			}})
			return
		}
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}
	errJson := Uhandler.service.Register(&user)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}

	// before setting the session we need the actual id of the user
	userData, errJson := Uhandler.service.GetUser(&models.Login{LoginField: user.Nickname})
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}
	// var login = models.Login{LoginField: user.Nickname}
	session, err_ := Uhandler.service.SetUserSession(userData)
	if err_ != nil {
		WriteJsonErrors(w, *err_)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   session.Token,
		Expires: session.ExpDate,
		Path:    "/",
	})
	// we don't need to write back the data for the repsonse ( sentitive data ;)
	WriteDataBack(w, models.UserData{
		IsLoggedIn: true,
		Id:         userData.Id,
		Nickname:   userData.Nickname,
	})
}
