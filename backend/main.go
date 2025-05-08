package main

import (
	"fmt"
	"net/http"

	database "real-time-forum/backend/database"
	"real-time-forum/backend/handler"
	"real-time-forum/backend/middleware"
	"real-time-forum/backend/repositories"
	"real-time-forum/backend/routes"
	"real-time-forum/backend/service"
)

func main() {
	db, err := database.InitDB("./database/forum.db")
	if err != nil {
		fmt.Println("err", err)
	}

	err = db.ReadSQL("./database/db.sql")
	if err != nil {
		fmt.Println("err", err)
	}

	// setup layers
	repo := repositories.NewAppRepository(db.Database)
	service := service.NewPostService(repo)
	postHanlder := handler.NewPostHandler(service)
	commentHandler := handler.NewCommentHandler(service)
	userHandler := handler.NewUserHandler(service)
	middleware := middleware.NewMiddleWare(userHandler.Login)
	routes.SetRoutes(postHanlder, commentHandler, userHandler)
	defer db.Database.Close()

	fmt.Println("Listening on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
