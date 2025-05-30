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
type ClientList map[int][]*Client

type Client struct {
	service    *service.AppService
	connection *websocket.Conn
	chatServer *ChatServer
	Message    chan *models.Message
	ErrorJson  chan *models.ErrorJson
	Done       chan struct{}
	CloseOnce   sync.Once
	userId     int
}

type ChatServer struct {
	service  *service.AppService
	clients  ClientList
	upgrader websocket.Upgrader
	sync.RWMutex
}

// https://stackoverflow.com/questions/65034144/how-to-add-a-trusted-origin-to-gorilla-websockets-checkorigin
func NewChatServer(service *service.AppService) *ChatServer {
	return &ChatServer{
		service: service,
		clients: make(ClientList),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func NewClient(conn *websocket.Conn, server *ChatServer) *Client {
	return &Client{
		service:    server.service,
		connection: conn,
		chatServer: server,
		Message:    make(chan *models.Message),
		ErrorJson:  make(chan *models.ErrorJson),
		Done:       make(chan struct{}),
	}
}
