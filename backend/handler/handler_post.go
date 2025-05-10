package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"real-time-forum/backend/models"
)

// Rewrite the PostHandler in the proper way
// Write the added post f response again (good practice)
func (Phandler *PostHandler) addPost(w http.ResponseWriter, r *http.Request) {
	var post *models.Post
	err := json.NewDecoder(r.Body).Decode(&post); 
	if err != nil {
		if err ==io.EOF {
			WriteJsonErrors(w, models.ErrorJson{Status: 400 , Message: &models.PostError {
        
			}})
		}
		// which status code to return
		errJSon := models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)}
		WriteJsonErrors(w, errJSon)
		return
	}
	err_ := Phandler.service.AddPost(post)
	if err != nil {
		WriteJsonErrors(w, *err_)
		return
	}
	WriteDataBack(w, post)
}

func (Phandler *PostHandler) getPosts(w http.ResponseWriter) {
	posts, err := Phandler.service.GetPosts()
	if err != nil {
		WriteJsonErrors(w, *err)
		return
	}
	err_ := json.NewEncoder(w).Encode(posts)
	if err_ != nil {
		errJSon := models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
		WriteJsonErrors(w, errJSon)
		return
	}
}

func (Phandler *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		Phandler.getPosts(w)
		return
	case http.MethodPost:
		Phandler.addPost(w, r)
		return
	default:
		errJson := models.ErrorJson{Status: 405, Message: "Method Not Allowed!!"}
		w.WriteHeader(errJson.Status)
		json.NewEncoder(w).Encode(errJson)
		return
	}
}
