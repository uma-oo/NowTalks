package middleware

import (
	"net/http"

	"real-time-forum/backend/service"
)

// dunno if this method will work or not !!!!

// bacause we are having http.Hanlder interface types





type Middleware struct {
	MiddlewareHanlder http.Handler
	service           *service.AppService
}
// Hna i wrapped the Middleware type here 7itash 7tajit wa7d l method!!
type LoginRegisterMiddleWare struct {
	MiddlewareHanlder Middleware
    handler http.Handler
}

func NewMiddleWare(handler http.Handler, service *service.AppService) *Middleware {
	return &Middleware{handler, service}
}


func NewLoginMiddleware(m Middleware, handler http.Handler) *LoginRegisterMiddleWare {
	return &LoginRegisterMiddleWare{m, handler}
}
