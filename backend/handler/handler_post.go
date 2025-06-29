package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"real-time-forum/backend/models"
)

// Rewrite the PostHandler in the proper way
// Write the added post f response again (good practice)
func (Phandler *PostHandler) addPost(w http.ResponseWriter, r *http.Request) {
	var post *models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		if err == io.EOF {
			WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: &models.PostError{
				Title:      "empty Title field!",
				Content:    "empty Content fiedl!",
				Categories: "select at least one category!",
			}})
			return
		}
		// which status code to return
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}
	// even if the userid is given wrong we insert the correct one
	post.UserId = Phandler.service.GetUserIdFromSession(r)
	postCreated, err_ := Phandler.service.AddPost(post)
	if err_ != nil {
		WriteJsonErrors(w, *err_)
		return
	}
	WriteDataBack(w, postCreated)
}

func (Phandler *PostHandler) getPosts(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	session, _ := Phandler.service.GetSessionByTokenEnsureAuth(cookie.Value)
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	if errConvoff != nil {
		WriteJsonErrors(w, *models.NewErrorJson(400, "ERROR!! Incorrect offset"))
		return
	}

	posts, err_get := Phandler.service.GetPosts(session.UserId, offset)
	if err_get != nil {
		WriteJsonErrors(w, *err_get)
		return
	}

	err_ := json.NewEncoder(w).Encode(posts)
	if err_ != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err_)})
		return
	}
}

func (Phandler *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		Phandler.getPosts(w, r)
		return
	case http.MethodPost:
		Phandler.addPost(w, r)
		return
	default:
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!!"})
		return
	}
}
