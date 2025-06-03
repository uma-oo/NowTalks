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
				Categories: "ERROR!! Incorrect Format of category ID or There is No category affected!",
			}})
			return
		}
		// which status code to return
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return
	}
	postCreated, err_ := Phandler.service.AddPost(post)
	if err_ != nil {
		fmt.Println("error adding post")
		fmt.Printf("err_: %v\n", err_)
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
	if ok && categories != nil {
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
	err_ := json.NewEncoder(w).Encode(posts)
	if err_ != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err_)})
		return
	}
}





func (Phandler *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL", r.URL.Path)
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
