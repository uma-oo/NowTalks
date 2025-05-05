package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"real-time-forum/backend/models"
)

func (Phandler *PostHandler) PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodPost:
		var post models.Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			// which status code to return
			errJSon := models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)}
			w.WriteHeader(errJSon.Status)
			json.NewEncoder(w).Encode(errJSon)
			return
		}
		err := Phandler.service.AddPost(post)
		if err != nil {
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(err)
			return
		}
	case http.MethodGet:
		posts, err := Phandler.service.GetPosts()
		if err != nil {
			w.WriteHeader(err.Status)
			json.NewEncoder(w).Encode(err)
			return
		}
		err_ := json.NewEncoder(w).Encode(posts)
		if err_ != nil {
			errJSon := models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)}
			w.WriteHeader(errJSon.Status)
			json.NewEncoder(w).Encode(errJSon)
			return
		}

	default:
		errJSon := models.ErrorJson{Status: 405, Message: "Method not Allowed!!"}
		WriteJsonErrors(w, errJSon)
		return
	}
}
