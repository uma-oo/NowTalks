package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"real-time-forum/backend/models"
)

// do we need to check the id of the receiver (if it exists in the database )
func (messages *MessagesHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	session, _ := messages.service.GetSessionByTokenEnsureAuth(cookie.Value)
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	if errConvoff != nil {
		WriteJsonErrors(w, *models.NewErrorJson(400, "ERROR!! Incorrect Format offset"))
		return
	}
	receiver_id, errConvrec := strconv.Atoi(r.URL.Query().Get("receiver_id"))
	if errConvrec != nil {
		WriteJsonErrors(w, *models.NewErrorJson(400, "ERROR!! Incorrect Format of the receiver_id"))
		return
	}
	type_ := r.URL.Query().Get("type")
	if type_ != "old" && type_ != "new" {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "ERROR!! type is not specified"})
		return
	}
	exists, errJson := messages.service.UserExists(receiver_id)
	if errJson != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	// check if the user 2 exists
	if !exists {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "ERROR!! receiver_id Incorrect"})
		return
	}
	mesages, errJson := messages.service.GetMessages(session.UserId, receiver_id, offset, type_)
	if errJson != nil {
		WriteJsonErrors(w, *models.NewErrorJson(errJson.Status, errJson.Message))
		return
	}
	err := json.NewEncoder(w).Encode(mesages)
	if err != nil {
		WriteJsonErrors(w, *models.NewErrorJson(500, fmt.Sprintf("%v", err)))
		return
	}
}

func (messages *MessagesHandler) UpdataReadStatus(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	session, _ := messages.service.GetSessionByTokenEnsureAuth(cookie.Value)
	receiver_id, errConvrec := strconv.Atoi(r.URL.Query().Get("receiver_id"))
	if errConvrec != nil {
		WriteJsonErrors(w, *models.NewErrorJson(400, "ERROR!! Incorrect Format of the receiver_id"))
		return
	}
	exists, errJson := messages.service.UserExists(receiver_id)
	if errJson != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
	// check if the user 2 exists
	if !exists {
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: "ERROR!! receiver_id Incorrect"})
		return
	}

	errJson = messages.service.EditReadStatus(session.UserId, receiver_id)
	if errJson != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: errJson.Status, Message: errJson.Message})
		return
	}
}

func (messages *MessagesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		messages.GetMessages(w, r)
	case http.MethodPatch:
		messages.UpdataReadStatus(w, r)
	default:
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!!"})
		return
	}
}
