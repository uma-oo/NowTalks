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
	var login = &models.Login{}
	if r.Method != http.MethodPost {
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not Allowed!"})
		return
	}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		if err==io.EOF {
			// case if the body sent if empty
			WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.Login {
					LoginField: "ERROR!! Empty Login field!!",
					Password: "ERROR!! Emty Password field!!",
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

	// before setting the session we need the actual id of the user
	userData, errJson := Uhandler.service.GetUser(login)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}
	// var login = models.Login{LoginField: user.Nickname}
	ss , e := Uhandler.service.GetSessionByUserId(userData.Id)
	fmt.Println("ss", ss.ExpDate,ss.IsExpired(), e)
	session, err_ := Uhandler.service.SetUserSession(userData)
	if err_ != nil {
		fmt.Println("errror", err_)
		WriteJsonErrors(w, *err_)
		return
	}
	// Path knt nassyaha dakshi 3lash makantsh tl3
	fmt.Println(session.IsExpired())
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   session.Token,
		Expires: session.ExpDate,
		Path: "/",
	})
    
	// in the login we don't need to rewrite the data ??? 
	// allahu a3lam
	WriteDataBack(w, login)
}
