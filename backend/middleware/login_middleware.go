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


func (LogRegM *LoginRegisterMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request){
	session , err := LogRegM.GetAuthUser(r)
	if session != nil {
		handler.WriteJsonErrors(w, models.ErrorJson{
			Status: 403,
			Message: "ERROR!! Already have credentials!!",
		})
		return
          
	}
	if err != nil {
		handler.WriteJsonErrors(w, models.ErrorJson{
			Status: 400,
			Message: fmt.Sprintf("%v", err),
		})
		return
	}
	LogRegM.MiddlewareHanlder.ServeHTTP(w,r)
	

}
func (LogRegM *LoginRegisterMiddleWare) GetAuthUser(r *http.Request) (*models.Session, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			return nil, nil
		}
		return nil, models.NewErrorJson(400, "ERROR!! There was an error in the Request!!")
	}
	session, _ := LogRegM.service.GetUserSessionByToken(cookie.Value)
	if session.IsExpired() {
		return nil, nil
	}
	return session, nil
}



