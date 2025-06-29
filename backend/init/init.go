package init

import (
	"database/sql"
	"net/http"

	"real-time-forum/backend/handler"
	"real-time-forum/backend/repositories"
	"real-time-forum/backend/routes"
	"real-time-forum/backend/service"
)

// SETUP THE LAYERS
func InitSetup(database *sql.DB) *http.ServeMux {
	repo := repositories.NewAppRepository(database)
	service := service.NewPostService(repo)
	postHanlder := handler.NewPostHandler(service)
	commentHandler := handler.NewCommentHandler(service)
	userHandler := handler.NewUserHandler(service)
	logout := handler.NewLogoutHandler(service)
	loggedin := handler.NewUserDataHanlder(service)
	categories := handler.NewCategoriesHandler(service)
	chat := handler.NewChatServer(service)
	reactionHandler := handler.NewReactionHandler(service)
	users := handler.NewUsersHandler(service)
	messages := handler.NewMessagesHandler(service)
	mux := routes.SetRoutes(postHanlder, commentHandler,
		reactionHandler, userHandler, logout,
		users, loggedin, categories, chat, messages, service)
	return mux
}
