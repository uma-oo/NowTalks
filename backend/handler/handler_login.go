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
	if r.Method != http.MethodPost {
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not Allowed!"})
		return
	}
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
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}
	errJson := Uhandler.service.Login(login)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}

	// We are kinda sure that if the user has a token he cannot be here
	// we need now

	// before setting the session we need the actual id of the user

	// we don't need to handle the error here (7itash deja login katkhdm nfss l function u error hadnling is there)
	userData, _ := Uhandler.service.GetUser(login)
	// if there is a session update it
	session, errJson := Uhandler.service.GetSessionByUserId(userData.Id)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}
	if session != (&models.Session{}) {
		new_session, errUpdate := Uhandler.service.UpdateUserSession(session)
		if errUpdate != nil {
			WriteJsonErrors(w, *errUpdate)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "session",
			Value:   new_session.Token,
			Expires: new_session.ExpDate,
			Path:    "/",
		})
		return
	}

	session, err_ := Uhandler.service.SetUserSession(userData)
	if err_ != nil {
		WriteJsonErrors(w, *err_)
		return
	}
	// Path knt nassyaha dakshi 3lash makantsh tl3
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   session.Token,
		Expires: session.ExpDate,
		Path:    "/",
	})

	WriteDataBack(w, login)
}
