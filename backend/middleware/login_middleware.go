package middleware

import (
	"fmt"
	"net/http"

	handler "real-time-forum/backend/handler"
	"real-time-forum/backend/models"
)

// Login and Register middlwares
// e7m e7m wash hakka hadshi khassu ykun ??
// allahu a3laaam

func (LogRegM *LoginRegisterMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := LogRegM.GetAuthUser(r)
	if err != nil {
		handler.WriteJsonErrors(w, models.ErrorJson{
			Status:  err.Status,
			Message: fmt.Sprintf("%v", err.Message),
		})
		return
	}

	LogRegM.MiddlewareHanlder.ServeHTTP(w, r)
}

// the logic of this function is not correct so neeed a quick fix

func (LogRegM *LoginRegisterMiddleWare) GetAuthUser(r *http.Request) *models.ErrorJson {
	cookie, err := r.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil
		}
		return models.NewErrorJson(400, "ERROR!! There was an error in the Request!!")
	}
	has_session, session := LogRegM.service.CheckUserSession(cookie.Value)
	fmt.Println("has_session", has_session,"seesion", session)
	if has_session {
		if !session.IsExpired() {
			return models.NewErrorJson(403, "ERROR!! The User has a session!! Access Forbiden")
		}
	}
	return nil
}
