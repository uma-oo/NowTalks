package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"real-time-forum/backend/models"
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

type UserInfo struct {
	UserID      int
	Count       int
	LastRequest time.Time
}

type RateLimitMiddleWare struct {
	MiddlewareHanlder http.Handler
	service           *service.AppService
	Users             map[int]*UserInfo
}

func NewRateLimitMiddleWare(handler http.Handler, service *service.AppService) *RateLimitMiddleWare {
	return &RateLimitMiddleWare{handler, service, map[int]*UserInfo{}}
}

func NewMiddleWare(handler http.Handler, service *service.AppService) *Middleware {
	return &Middleware{handler, service}
}

func NewLoginMiddleware(handler http.Handler, service *service.AppService) *LoginRegisterMiddleWare {
	return &LoginRegisterMiddleWare{handler, service}
}

func WriteJsonErrors(w http.ResponseWriter, errJson models.ErrorJson) {
	w.WriteHeader(errJson.Status)
	json.NewEncoder(w).Encode(errJson)
}
