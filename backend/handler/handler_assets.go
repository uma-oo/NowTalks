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
	fmt.Println("filename", file_info)
	if err != nil {
		fmt.Println("")
		log.Fatal(err)
		log.Fatal(file_info)
		w.WriteHeader(http.StatusNotFound)
		return

	} else if file_info == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if file_info != nil && file_info.IsDir() {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, filepath.Join("../frontend", r.URL.Path[1:]))
}
