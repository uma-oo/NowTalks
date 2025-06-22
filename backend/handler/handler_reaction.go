package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"real-time-forum/backend/models"
)

// to avoid duplication we need to have the type of the entity and then send it with the entity_id

func (Rhanlder *ReactionHanlder) LikeEntity(w http.ResponseWriter, r *http.Request) {
	liked := models.Reaction{}
	if err := json.NewDecoder(r.Body).Decode(&liked); err != nil {
		if err == io.EOF {
			WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: &models.ReactionErr{
				EntityId:   " emptyEntityID field!",
				EntityType: " emptyEntityType field!",
			}})
			return
		}
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: " Bad Request!"})
		return
	}
	liked.UserId = Rhanlder.service.GetUserIdFromSession(r)
	entity_type_id := Rhanlder.service.GetTypeIdByName(liked.EntityType)
	if entity_type_id == 0 {
		errJson := models.ErrorJson{Status: 500, Message: " Internal Server Error"}
		WriteJsonErrors(w, errJson)
		return
	}
	liked.EntityTypeId = entity_type_id
	reaction, errJson := Rhanlder.service.HanldeReaction(&liked, 1)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}

	WriteDataBack(w, reaction)
}

func (Rhandler *ReactionHanlder) DislikeEnity(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	session, _ := Rhandler.service.GetSessionByTokenEnsureAuth(cookie.Value)
	disliked := models.Reaction{}

	if err := json.NewDecoder(r.Body).Decode(&disliked); err != nil {
		if err == io.EOF {
			WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: &models.ReactionErr{
				EntityId:   " emptyEntityID field!",
				EntityType: " emptyEntityType field!",
			}})
			return
		}
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: " Bad Request!"})
		return
	}

	disliked.UserId = session.UserId
	entity_type_id := Rhandler.service.GetTypeIdByName(disliked.EntityType)
	if entity_type_id == 0 {
		// to be verified if the status code is 500 or 400
		errJson := models.ErrorJson{Status: 400, Message: " The entity requested Does not Exist!"}
		WriteJsonErrors(w, errJson)
		return
	}
	disliked.EntityTypeId = entity_type_id
	reaction, errJson := Rhandler.service.HanldeReaction(&disliked, -1)
	if errJson != nil {
		WriteJsonErrors(w, *errJson)
		return
	}
	WriteDataBack(w, reaction)
}

func (RHanlder *ReactionHanlder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: " Method Not allowed!!"})
		return
	}
	switch r.URL.Path {
	case "/api/react/like":
		RHanlder.LikeEntity(w, r)
		return
	case "/api/react/dislike":
		RHanlder.DislikeEnity(w, r)
		return
	default:
		WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: " Not Found!"})
		return
	}
}
