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

	// we don't need to write back the data for the repsonse ( sentitive data ;)
	// WriteDataBack(w, user)
}



func (Uhandler *UserHanlder) Login(w http.ResponseWriter, r *http.Request) {
	var user models.Login
	if r.Method!= http.MethodPost {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Method not Allowed!"})
		return
	}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}
	errJson := Uhandler.service.Login(&user)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return 
	}
	WriteDataBack(w, user)

}
