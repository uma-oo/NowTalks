package middleware

import (
	"net/http"
	"real-time-forum/backend/service"
)

// dunno if this method will work or not !!!!

// bacause we are having http.Hanlder interface types
type Middleware struct {
	MiddlewareHanlder http.Handler
	service *service.AppService
}

func NewMiddleWare(handler http.Handler, service *service.AppService) *Middleware {
	return &Middleware{handler , service}
}
