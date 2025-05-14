package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// frontend/api/comment.js
func HandleAssets(w http.ResponseWriter, r *http.Request) {
	file_info, err := os.Stat(filepath.Join("../frontend", r.URL.Path[1:]))
	if err != nil {
		fmt.Println("hna", err)
		log.Fatal("hhhhhhhhhhhh")
		w.WriteHeader(http.StatusNotFound)
	}
	if file_info.IsDir() {
		w.WriteHeader(http.StatusForbidden)
	}

	http.ServeFile(w, r, filepath.Join("../frontend", r.URL.Path[1:]))
}
