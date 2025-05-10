package middleware

import (
	"net/http"

	handler "real-time-forum/backend/handler"
	"real-time-forum/backend/models"
)

// could be returning a boolean but to see again
func (m *Middleware) GetAuthUser(r *http.Request) (*models.Session, *models.ErrorJson) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil, &models.ErrorJson{Status: 401, Message: "ERROR!! Unauthorized Access"}
	}
	// check if the value of the cookie is correct and if not expired!!!
	session, errJson := m.service.GetUserSessionByToken(cookie.Value)
	if errJson != nil || session.IsExpired() {
		return nil, errJson
	}
	return session, nil
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

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := m.GetAuthUser(r)
	if err != nil {
		handler.WriteJsonErrors(w, *err)
		return
	}
	m.MiddlewareHanlder.ServeHTTP(w, r)
}
