package handler

import (
	"encoding/json"
	"net/http"
	"sync"

	"real-time-forum/backend/models"
	"real-time-forum/backend/service"

	"github.com/gorilla/websocket"
)

type PostHandler struct {
	service *service.AppService
}
type CommentHandler struct {
	service *service.AppService
}

type UserHanlder struct {
	service *service.AppService
}

type UserData struct {
	service *service.AppService
}

type CategoriesHandler struct {
	service *service.AppService
}

type Logout UserHanlder

func NewLogoutHandler(service *service.AppService) *Logout {
	return &Logout{service: service}
}

func NewCommentHandler(service *service.AppService) *CommentHandler {
	return &CommentHandler{service: service}
}

func NewPostHandler(service *service.AppService) *PostHandler {
	return &PostHandler{service: service}
}

func NewUserHandler(service *service.AppService) *UserHanlder {
	return &UserHanlder{service: service}
}

func NewUserDataHanlder(service *service.AppService) *UserData {
	return &UserData{service: service}
}

func NewCategoriesHandler(service *service.AppService) *CategoriesHandler {
	return &CategoriesHandler{service: service}
}

// NewPostService

func WriteJsonErrors(w http.ResponseWriter, errJson models.ErrorJson) {
	w.WriteHeader(errJson.Status)
	json.NewEncoder(w).Encode(errJson)
}

func WriteDataBack(w http.ResponseWriter, data any) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(&data)
}

// so the server will be a map of connexions and a mutex with it
// section for the chat implemenatation
type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	chatServer *ChatServer
}

type ChatServer struct {
	clients ClientList
	upgrader websocket.Upgrader
	sync.RWMutex
}




// https://stackoverflow.com/questions/65034144/how-to-add-a-trusted-origin-to-gorilla-websockets-checkorigin
func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(ClientList),
		upgrader: websocket.Upgrader {
			ReadBufferSize: 1024,
			WriteBufferSize: 1024,
		},
	}
}

func NewClient(conn *websocket.Conn, server *ChatServer) *Client {
	return &Client{
		connection: conn,
		chatServer: server,
	}
}
