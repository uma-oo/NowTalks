package middleware

import (
	"net/http"
	handler "real-time-forum/backend/handler"
	"real-time-forum/backend/models"
)

// Login and Register middlwares
// e7m e7m wash hakka hadshi khassu ykun ??
// allahu a3laaam


func (LogRegM *LoginRegisterMiddleWare) ServeHTTP(w http.ResponseWriter, r *http.Request){
	session , _ := LogRegM.MiddlewareHanlder.GetAuthUser(r)
	if session != nil {
		handler.WriteJsonErrors(w, models.ErrorJson{
			Status: 403,
			Message: "ERROR!! Already have credentials!!",
		})
		return
          
	}
	LogRegM.MiddlewareHanlder.ServeHTTP(w,r)
	

}



