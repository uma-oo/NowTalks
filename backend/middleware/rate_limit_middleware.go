package middleware

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"real-time-forum/backend/models"
)

var COUNT = 0

// This is the middleware that comes after the the authentication
// there are many cases to handle here
func (RateLimitM *RateLimitMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	COUNT++
	fmt.Printf("COUNT: %v\n", COUNT)
	ipAddress := r.RemoteAddr
	// m7taja l ip u salam
	ip, _, _ := net.SplitHostPort(ipAddress)
	value, ok := RateLimitM.Users.Load(ip)

	if ok {
		fmt.Printf("IF value.(*ClientInfo).Count: %v  value.(*ClientInfo).IP :%v\n", value.(*ClientInfo).Count, value.(*ClientInfo).IP)
		value.(*ClientInfo).Lock()
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
		value.(*ClientInfo).Unlock()
	} else {
		RateLimitM.Users.Store(ip, &ClientInfo{
			Count:       1,
			LastRequest: time.Now(),
			IP:          ip,
		})
		value, _ := RateLimitM.Users.Load(ip)
		fmt.Printf(" ELSE :value.(*ClientInfo).Count: %v  value.(*ClientInfo).IP :%v\n", value.(*ClientInfo).Count, value.(*ClientInfo).IP)
	}

	RateLimitM.MiddlewareHanlder.ServeHTTP(w, r)
}
