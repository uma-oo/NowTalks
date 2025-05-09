package middleware

import (
	"net/http"

	handler "real-time-forum/backend/handler"
	"real-time-forum/backend/service"
)

// dunno if this method will work or not !!!!

// bacause we are having http.Hanlder interface types
type Middleware struct {
	MiddlewareHanlder http.Handler
	service *service.AppService
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ , err := m.GetAuthUser(r)
	if err != nil {
		handler.WriteJsonErrors(w, *err)
		return
	}
	m.MiddlewareHanlder.ServeHTTP(w,r)
}

func NewMiddleWare(handler http.Handler, service *service.AppService) *Middleware {
	return &Middleware{handler , service}
}
