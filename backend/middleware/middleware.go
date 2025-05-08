package middleware

import (
	"net/http"

	handler "real-time-forum/backend/handler"
)

// dunno if this method will work or not !!!!
type MiddlewareHanlder  func(w http.ResponseWriter, r *http.Request)

type Middleware struct {
	MiddlewareHanlder
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := GetAuthUser(r)
	if err != nil {
		handler.WriteJsonErrors(w, *err)
		return
	}
	m.MiddlewareHanlder(w, r)
}

func NewMiddleWare(handler MiddlewareHanlder) *Middleware {
	return &Middleware{handler}
}
