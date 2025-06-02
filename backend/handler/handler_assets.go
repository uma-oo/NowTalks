package handler

import (
	"net/http"
	"path/filepath"
)

// frontend/api/comment.js



//  Impelement the API 

func HandleAssets(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join("../frontend", r.URL.Path[1:]))
}
