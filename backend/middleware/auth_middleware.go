package middleware

import (
	"net/http"
	"real-time-forum/backend/models"
	handler "real-time-forum/backend/handler"

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







func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := m.GetAuthUser(r)
	if err != nil {
		handler.WriteJsonErrors(w, *err)
		return
	}
	m.MiddlewareHanlder.ServeHTTP(w, r)
}
