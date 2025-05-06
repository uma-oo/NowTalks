package middleware

import "net/http"




type Logger struct {
	http.Handler
}