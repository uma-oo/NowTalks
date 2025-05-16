package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"real-time-forum/backend/models"
)

// GET THE request body
func (CHanlder *CommentHandler) addComment(w http.ResponseWriter, r *http.Request) {
	var comment *models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		if err == io.EOF {
			WriteJsonErrors(w, models.ErrorJson{
				Status: 400,
				Message: models.CommentError{
					Content: "ERROR!! Empty Content Field!",
				},
			})
			return
		}
		WriteJsonErrors(w, models.ErrorJson{Status: 400, Message: fmt.Sprintf("%v", err)})
		return

	}
	comment_created, err_ := CHanlder.service.AddComment(comment)
	if err_ != nil {
		WriteJsonErrors(w, *err_)
		return
	}
	WriteDataBack(w, comment_created)
}

func (CHanlder *CommentHandler) getComments(w http.ResponseWriter, r *http.Request) {
	// get the comments of a specific ID
	// FOR NOW let's just get them from the query
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
