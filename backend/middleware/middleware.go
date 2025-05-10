package middleware

import (
	"net/http"

	"real-time-forum/backend/service"
)


type Middleware struct {
	MiddlewareHanlder http.Handler
	service           *service.AppService
}


type LoginRegisterMiddleWare struct {
	MiddlewareHanlder http.Handler
    service           *service.AppService
}

func NewMiddleWare(handler http.Handler, service *service.AppService) *Middleware {
	return &Middleware{handler, service}
}


func NewLoginMiddleware( handler http.Handler, service *service.AppService) *LoginRegisterMiddleWare {
	return &LoginRegisterMiddleWare{handler, service}
}
