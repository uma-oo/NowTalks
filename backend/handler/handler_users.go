package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"real-time-forum/backend/models"
)

func (Users *Users) GetUsers(w http.ResponseWriter, r *http.Request) {
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	if errConvoff != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Bad Request!! Post Not Found Or Incorrect offset!"})
		return
	}
	users, errJson := Users.service.GetUsers(offset)
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
	if r.Method != http.MethodGet {
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not ALlowed!"})
		return
	}
	if r.URL.Path != "/api/users" {
		WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "ERROR!! Page Not Found!"})
		return
	}

	users.GetUsers(w, r)
}
