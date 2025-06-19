package middleware

import (
	"net"
	"net/http"
	"time"

	"real-time-forum/backend/models"
)

// This is the middleware that comes after the the authentication
// there are many cases to handle here
func (RateLimitM *RateLimitMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ipAddress := r.RemoteAddr
	// m7taja l ip u salam
	ip, _, _ := net.SplitHostPort(ipAddress)
	value, ok := RateLimitM.Users.Load(ip)
	if ok {
		value.(*ClientInfo).Lock()
		defer value.(*ClientInfo).Unlock()
		if time.Since(value.(*ClientInfo).LastRequest) > RateLimitM.MaxDuration {

			value.(*ClientInfo).Count = 1
			value.(*ClientInfo).LastRequest = time.Now()
		} else {
			if value.(*ClientInfo).Count > RateLimitM.MaxRequests {
				WriteJsonErrors(w, models.ErrorJson{Status: http.StatusTooManyRequests, Message: "ERROR!! Too many Requests"})
				return
			}
			value.(*ClientInfo).Count++
		}
	} else {
		RateLimitM.Users.Store(ip, &ClientInfo{
			Count:       1,
			LastRequest: time.Now(),
		})
	}

	RateLimitM.MiddlewareHanlder.ServeHTTP(w, r)
}
