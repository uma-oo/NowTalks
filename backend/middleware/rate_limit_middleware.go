package middleware

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
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
	staticDir := filepath.Join("..", "frontend")
	// fmt.Println("filepath.Clean(r.URL.Path)", filepath.Clean(r.URL.Path))
	fmt.Printf("filepath.Clean(r.URL.Path): %v\n", filepath.Clean(r.URL.Path))
	fileName := filepath.Clean(r.URL.Path)
	if strings.HasPrefix(filepath.Clean(r.URL.Path), "/frontend") {
		splited := strings.Split(filepath.Clean(r.URL.Path), "/")
		fileName = strings.Join(splited[2:], "/")
	}
	requestedFile := filepath.Join(staticDir, fileName)
	// file_path := strings.Split(requestedFile, "frontend")
	// fmt.Println("file_path", file_path[len(file_path)-1])

	file_info, err := os.Stat(requestedFile)
    fmt.Printf("requestedFile: %v\n", requestedFile)
	if isFrontendPath(r.URL.Path) {
		fmt.Println("isfronted", requestedFile)
		http.ServeFile(w, r, "../frontend/index.html")
		return
	}
	if err != nil && !isFrontendPath(r.URL.Path) {
		fmt.Println("is not found", requestedFile)
		serveIndexWithStatus(w, 404, filepath.Join(staticDir, "index.html"))
		return
	}
	if file_info.IsDir() && file_info.Name() != "frontend" {
		serveIndexWithStatus(w, 403, filepath.Join(staticDir, "index.html"))
		return
	}

	http.ServeFile(w, r, filepath.Join("..", "frontend", requestedFile))
}

func serveIndexWithStatus(w http.ResponseWriter, statusCode int, indexPath string) {
	data, err := os.ReadFile(indexPath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)
	w.Write(data)
}
