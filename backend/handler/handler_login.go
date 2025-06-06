package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"real-time-forum/backend/models"
)

//    Login

// Middlware = checks if there is a session
// but if not we need check the body of the request o
// reset the sesssion

func (Uhandler *UserHanlder) Login(w http.ResponseWriter, r *http.Request) {
	login := &models.Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		if err == io.EOF {
			// case if the body sent if empty
			WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.Login{
					LoginField: "ERROR!! Empty Login field!!",
					Password:   "ERROR!! Emty Password field!!",
				},
			})
			return
		}
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v 1", err)})
		return
	}
	user, errJson := Uhandler.service.Login(login)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}

	// We are kinda sure that if the user has a token he cannot be here
	// we need now
	// before setting the session we need the actual id of the user
	// if there is a session update it

	UserData, session, errJSON := Uhandler.service.CreateOrUpdateSession(user)
	if errJSON != nil {
		WriteJsonErrors(w, models.ErrorJson{
			Status:  errJSON.Status,
			Message: errJSON.Message,
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   session.Token,
		Expires: session.ExpDate,
		Path:    "/",
	})
	WriteDataBack(w, UserData)
}


