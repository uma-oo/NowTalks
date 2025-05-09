package middleware

import (
	"net/http"

	"real-time-forum/backend/models"
)

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
