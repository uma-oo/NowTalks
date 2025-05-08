package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/backend/models"
)

//    Login

func (Uhandler *UserHanlder) Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login
	if r.Method != http.MethodPost {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Method not Allowed!"})
		return
	}
    // check if the user has the session before checking the data sent to the api 
    cookie, err := r.Cookie("session")
	if err != nil {
		WriteJsonErrors(w, *models.NewErrorJson(r.Response.StatusCode, "THERE IS NO COOKIE SET"))
	}

	fmt.Println("cookie", cookie.Value)


	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}
	errJson := Uhandler.service.Login(&login)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}
	WriteDataBack(w, login)
}
