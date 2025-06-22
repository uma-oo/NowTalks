package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func isFrontendPath(path string) bool {
	return path == "/login" || path == "/" || path == "/register" || path == "error"
}

func ServeStaticFiles(w http.ResponseWriter, r *http.Request) {
	staticDir := filepath.Join("..", "frontend")
	fileName := filepath.Clean(r.URL.Path)
	if strings.HasPrefix(filepath.Clean(r.URL.Path), "/frontend") {
		splited := strings.Split(filepath.Clean(r.URL.Path), "/")
		fileName = strings.Join(splited[2:], "/")
	}
	requestedFile := filepath.Join(staticDir, fileName)

	file_info, err := os.Stat(requestedFile)
	if isFrontendPath(r.URL.Path) {
		http.ServeFile(w, r, "../frontend/index.html")
		return
	}
	if err != nil && !isFrontendPath(r.URL.Path) {
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
