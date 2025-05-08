package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/backend/models"
)

func (Uhandler *UserHanlder) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Method != http.MethodPost {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Method not Allowed!"})
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
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
		fmt.Println("errror", session)
		WriteJsonErrors(w, *err_)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   session.Token,
		Expires: session.ExpDate,
	})
	// we don't need to write back the data for the repsonse ( sentitive data ;)
	// WriteDataBack(w, user)
}
