package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"real-time-forum/backend/models"
)

// GET THE request body
func (CHanlder *CommentHandler) addComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		errJson := models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)}
		WriteJsonErrors(w, errJson)
		return
	}
	err := CHanlder.service.AddComment(&comment)
	if err != nil {
		WriteJsonErrors(w, *err)
		return
	}
}

func (CHanlder *CommentHandler) getComments(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url_query", r.URL)
	postId, err := strconv.ParseInt(r.URL.Query().Get("postId"), 10, 64)
	if err != nil {
		errJson := models.ErrorJson{Status: 400, Message: "Bad Request! Post Not Found"}
		WriteJsonErrors(w, errJson)
		return
	}
	comments, err_ := CHanlder.service.GetComments(int(postId))
	if err_ != nil {
		WriteJsonErrors(w, *err_)
	}
	if err := json.NewEncoder(w).Encode(&comments); err != nil {
		errJson := models.ErrorJson{Status: 400, Message: "Bad Request!"}
		WriteJsonErrors(w, errJson)
		return
	}
}

func (CHanlder *CommentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		CHanlder.getComments(w, r)
		return
	case http.MethodPost:
		CHanlder.addComment(w, r)
		return
	default:
		errJson := models.ErrorJson{Status: 405, Message: "Method Not Allowed!!"}
		w.WriteHeader(errJson.Status)
		json.NewEncoder(w).Encode(errJson)
		return

	}
}
