package middleware

import (
	"fmt"
	"net/http"
	"time"
)

// This is the middleware that comes after the the authentication

func (RateLimitM *RateLimitMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := RateLimitM.service.GetUsernameFromSession(r)
	userInfo, ok := RateLimitM.Users[userId]
	if ok {
		if userInfo.Count > 11 && userInfo.LastRequest.Before(time.Now()) {
			fmt.Println("salaam")
		}
	} else {
		RateLimitM.Users[userId] = UserInfo{
			UserID: userId,
			Count:  1,

			LastRequest: time.Now(),
		}
	}
}
