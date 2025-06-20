package middleware

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"real-time-forum/backend/models"
)

func (rl *RateLimitMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Invalid IP address", http.StatusInternalServerError)
		return
	}

	val, ok := rl.Users.Load(ip)
	if ok {
		clientInfo, ok := val.(*ClientInfo)
		if !ok {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		clientInfo.Lock()
		defer clientInfo.Unlock()

		if time.Since(clientInfo.LastRequest) > rl.MaxDuration {
			clientInfo.Count = 1
			clientInfo.LastRequest = time.Now()
		} else {
			if clientInfo.Count >= rl.MaxRequests {
				fmt.Println("", clientInfo.Count)
				WriteJsonErrors(w, models.ErrorJson{
					Status:  http.StatusTooManyRequests,
					Message: "ERROR!! Too many Requests",
				})
				return
			}
			clientInfo.Count++
		}
	} else {
		rl.Users.Store(ip, &ClientInfo{
			Count:       1,
			LastRequest: time.Now(),
		})
	}

	rl.MiddlewareHanlder.ServeHTTP(w, r)
}

func isFrontendPath(path string) bool {
	return path == "/login" || path == "/" || path == "/register" || path == "error"
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	file_info, err := os.Stat(filepath.Join("../frontend/", r.URL.Path[1:]))
	if isFrontendPath(r.URL.Path) {
		http.ServeFile(w, r, "../frontend/index.html")
		return
	}
	if err != nil && !isFrontendPath(r.URL.Path) {
		WriteJsonErrors(w, models.ErrorJson{Status: 404, Message: "Page Not Found :)"})
		return
	}
	if file_info.IsDir() && file_info.Name() != "frontend" {
		WriteJsonErrors(w, models.ErrorJson{Status: 403, Message: "Access Forbiden :)"})
		return
	}
	// http.FileServer(http.Dir("../frontend")).ServeHTTP(w, r)
	http.ServeFile(w, r, filepath.Join("../frontend/", r.URL.Path[1:]))
}
