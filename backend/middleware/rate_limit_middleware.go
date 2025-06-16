package middleware

import (
	"net/http"
	"time"

	"real-time-forum/backend/models"
)

// This is the middleware that comes after the the authentication
// there are many cases to handle here
func (RateLimitM *RateLimitMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := RateLimitM.service.GetUsernameFromSession(r)
	userInfo, ok := RateLimitM.Users[userId]
	if ok {
		if userInfo.Count > 11 && userInfo.LastRequest.Before(time.Now().Add(1*time.Minute)) {
			WriteJsonErrors(w, models.ErrorJson{Status: 429, Message: "Hey! Too Many requests!!"})
			return
		} else if !userInfo.LastRequest.Before(time.Now().Add(1 * time.Minute)) {
			userInfo.Count = 1
		} else {
			userInfo.Count++
		}
		userInfo.LastRequest = time.Now()
	} else {
		RateLimitM.Users[userId] = UserInfo{
			UserID:      userId,
			Count:       1,
			LastRequest: time.Now(),
		}
	}
	RateLimitM.MiddlewareHanlder.ServeHTTP(w, r)
}
