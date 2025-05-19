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
				Title:   "ERROR!! Empty Title field!",
				Content: "ERROR!! Empty Content fiedl!",
			}})
			return
		}
		// which status code to return
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}
	postCreated, err_ := Phandler.service.AddPost(post)
	if err_ != nil {
		WriteJsonErrors(w, *err_)
		return
	}
	WriteDataBack(w, postCreated)
}

func (Phandler *PostHandler) getPosts(w http.ResponseWriter, r *http.Request) {
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	if errConvoff != nil {
		WriteJsonErrors(w, *models.NewErrorJson(400, "ERROR!! Incorrect offset"))
		return
	}

	categories, ok := r.URL.Query()["category"]
	var posts []models.Post
	err_get := &models.ErrorJson{}
	if ok && categories!= nil {
		posts, err_get = Phandler.service.GetPostsByCategory(offset, categories...)
		if err_get != nil {
			WriteJsonErrors(w, *err_get)
			return
		}
	} else {
		posts, err_get = Phandler.service.GetPosts(offset)
		if err_get != nil {
			WriteJsonErrors(w, *err_get)
			return
		}

	}

	// read the body of the request

	err_ := json.NewEncoder(w).Encode(posts)
	if err_ != nil {
		errJSon := models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err_)}
		WriteJsonErrors(w, errJSon)
		return
	}
}


func (Phandler *PostHandler) ReactToPost(w http.ResponseWriter, r *http.Request){
 

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
		errJson := models.ErrorJson{Status: 405, Message: "Method Not Allowed!!"}
		w.WriteHeader(errJson.Status)
		json.NewEncoder(w).Encode(errJson)
		return
	}
}



