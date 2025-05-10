package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/backend/models"
)

//    Login

func (Uhandler *UserHanlder) Login(w http.ResponseWriter, r *http.Request) {
	var login = &models.Login{}
	if r.Method != http.MethodPost {
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "Method not Allowed!"})
		return
	}
	// don't need this checking anymore
    // check if the user has the session before checking the data sent to the api 
	// fmt.Println("hnaa inside login" , r.Cookie)
    cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println("err1", err)
		WriteJsonErrors(w, *models.NewErrorJson(401, "THERE IS NO COOKIE SET"))
		return
	}
	
	fmt.Println("cookie", cookie.Value)


	err = json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}
	errJson := Uhandler.service.Login(login)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}
	// in the login we don't need to rewrite the data ??? 
	// allahu a3lam
	WriteDataBack(w, login)
}
