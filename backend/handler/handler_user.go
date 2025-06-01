package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"real-time-forum/backend/models"
)

func (Uhandler *UserHanlder) GetUsers(w http.ResponseWriter, r *http.Request) {
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	if errConvoff != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "Bad Request!! Post Not Found Or Incorrect offset!"})
		return
	}
	users, errJson := Uhandler.service.GetUsers(offset)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}
	if err := json.NewEncoder(w).Encode(&users); err != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
		return
	}
}

// add the endpoint of getusers
func (Uhandler *UserHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	switch {
	case r.Method == http.MethodPost && r.URL.Path == "/api/user/login":
		Uhandler.Login(w, r)
		return
	case r.Method == http.MethodPost && r.URL.Path == "/api/user/register":
		Uhandler.Register(w, r)
		return
	case r.Method == http.MethodGet && r.URL.Path == "/api/user/users":
		Uhandler.GetUsers(w, r)
		return
	case r.Method != http.MethodPost && r.Method != http.MethodGet:
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not allowed!!"})
		return
	default:
		WriteJsonErrors(w, *models.NewErrorJson(404, "ERROR!! Page Not Found!!"))
		return
	}
}
