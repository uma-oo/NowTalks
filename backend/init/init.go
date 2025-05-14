package init

import (
	"database/sql"

	"real-time-forum/backend/handler"
	"real-time-forum/backend/repositories"
	"real-time-forum/backend/routes"
	"real-time-forum/backend/service"
)




// SETUP THE LAYERS 
func InitSetup(database *sql.DB) {
	repo := repositories.NewAppRepository(database)
	service := service.NewPostService(repo)
	postHanlder := handler.NewPostHandler(service)
	commentHandler := handler.NewCommentHandler(service)
	userHandler := handler.NewUserHandler(service)
	logout := handler.NewLogoutHandler(service)
	// middleware := middleware.NewMiddleWare(postHanlder)
	routes.SetRoutes(postHanlder, commentHandler, userHandler, service, logout)
}
