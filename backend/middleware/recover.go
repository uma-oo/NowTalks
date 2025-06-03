package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	handler "real-time-forum/backend/handler"
	"real-time-forum/backend/models"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// log the panic and stack trace
				message := "Caught panic: %v, Stack trace: %s"
				log.Printf(message, err, string(debug.Stack()))
				errjson := models.ErrorJson{
					Status:  500,
					Message: "Internal Server Error",
				}

				handler.WriteJsonErrors(w, errjson)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
