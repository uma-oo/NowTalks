package middleware

import (
	"net/http"
	"time"

	"real-time-forum/backend/models"
)

// This is the middleware that comes after the the authentication
// there are many cases to handle here
func (RateLimitM *RateLimitMiddleWareLoggedIn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := RateLimitM.service.GetUsernameFromSession(r)
	userInfo, ok := RateLimitM.Users[userId]
	if ok {
		if userInfo.LastRequest.Add(1 * time.Minute).Before(time.Now()) {
			userInfo.Count = 1
		} else if userInfo.Count > 60 && userInfo.LastRequest.Before(time.Now().Add(1*time.Minute)) {
			WriteJsonErrors(w, models.ErrorJson{Status: 429, Message: "Hey! Too Many requests!!"})
			return
		} else {
			userInfo.Count++
		}
		userInfo.LastRequest = time.Now()
	} else {
		RateLimitM.Users[userId] = &UserInfo{
			UserID:      userId,
			Count:       1,
			LastRequest: time.Now(),
		}
	}
	RateLimitM.MiddlewareHanlder.ServeHTTP(w, r)
}

// rate limiter for the / on the route of /api/




func (rateLimiter *RateLimitter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if rateLimiter.LastRequest.Add(1 * time.Minute).Before(time.Now()) {
		rateLimiter.Count = 1
		rateLimiter.LastRequest = time.Now()
	} else if rateLimiter.Count > 1000 && rateLimiter.LastRequest.Before(time.Now().Add(1*time.Minute)) {
		WriteJsonErrors(w, models.ErrorJson{Status: 429, Message: "Hey! Too Many requests!!"})
		return
	} else {
		rateLimiter.Count++
		rateLimiter.LastRequest = time.Now()
	}
}
