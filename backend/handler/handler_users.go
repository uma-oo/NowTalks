package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"real-time-forum/backend/models"
)

func (Users *Users) GetUsers(w http.ResponseWriter, r *http.Request) {
	cookie , _ := r.Cookie("session")
	session ,_ := Users.service.GetSessionByTokenEnsureAuth(cookie.Value)
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	if errConvoff != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Bad Request!! Incorrect offset!"})
		return
	}
	users, errJson := Users.service.GetUsers(offset, session.UserId)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}
	if err := json.NewEncoder(w).Encode(&users); err != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
		return
	}
}

func (users *Users) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not ALlowed!"})
		return
	}

	users.GetUsers(w, r)
}
