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
	cookie, _ := r.Cookie("session")
	session, _ := CHanlder.service.GetSessionByTokenEnsureAuth(cookie.Value)
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
	comment.UserId = session.UserId
	comment_created, err_ := CHanlder.service.AddComment(comment)
	if err_ != nil {
		WriteJsonErrors(w, *err_)
		return
	}
	WriteDataBack(w, comment_created)
}

func (CHanlder *CommentHandler) getComments(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session")
	session, _ := CHanlder.service.GetSessionByTokenEnsureAuth(cookie.Value)
	offset, errConvoff := strconv.Atoi(r.URL.Query().Get("offset"))
	postId, err := strconv.Atoi(r.URL.Query().Get("post"))
	if err != nil || errConvoff != nil {
		errJson := models.ErrorJson{Status: 400, Message: "Bad Request!! Post Not Found Or Incorrect offset!"}
		WriteJsonErrors(w, errJson)
		return
	}
	comments, err_ := CHanlder.service.GetComments(session.UserId, postId, offset)
	if err_ != nil {
		WriteJsonErrors(w, *err_)
	}
	if err := json.NewEncoder(w).Encode(&comments); err != nil {
		WriteJsonErrors(w, models.ErrorJson{Status: 500, Message: fmt.Sprintf("%v", err)})
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
		WriteJsonErrors(w, models.ErrorJson{Status: 405, Message: "ERROR!! Method Not Allowed!!"})
		return
	}
}
