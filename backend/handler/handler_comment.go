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
	limit, errConvlim := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	postId, err := strconv.Atoi(r.URL.Query().Get("post"))
	if err != nil || errConvlim != nil || errConvoff != nil {
		errJson := models.ErrorJson{Status: 400, Message: "Bad Request!! Post Not Found Or Incorrect offset or limit!"}
		WriteJsonErrors(w, errJson)
		return
	}
	comments, err_ := CHanlder.service.GetComments(postId, limit, offset)
	if err_ != nil {
		WriteJsonErrors(w, *err_)
	}
	if err := json.NewEncoder(w).Encode(&comments); err != nil {
		errJson := models.ErrorJson{Status: 400, Message: "Bad Request!"}
		WriteJsonErrors(w, errJson)
		return
	}
}

func (CHanlder *CommentHandler) LikeComment(w http.ResponseWriter, r *http.Request) {
}

func (CHanlder *CommentHandler) DislikeComment(w http.ResponseWriter, r *http.Request) {
}

func (CHanlder *CommentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		CHanlder.getComments(w, r)
		return
	case http.MethodPost:
		switch r.URL.Path[1:] {
		case "api/comment/like":
			CHanlder.LikeComment(w, r)
			return
		case "api/comment/dislike":
			CHanlder.DislikeComment(w, r)
			return
		case "api/comment/":
			CHanlder.addComment(w, r)
			return
		default:
			errJson := models.ErrorJson{Status: 404, Message: "ERROR!! Page Not Found!!"}
			WriteJsonErrors(w, errJson)
			return
		}
	default:
		errJson := models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!!"}
		WriteJsonErrors(w, errJson)
		return

	}
}
