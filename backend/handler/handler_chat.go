package handler

import (
	"net/http"
	"real-time-forum/backend/models"
	"strings"
)

func (server *ChatServer) ChatServerHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		if isHandshakeError(err) {
			WriteJsonErrors(w, *models.NewErrorJson(400, "ERROR!! There is something wrong with request Upgrade"))
			return
		}
		WriteJsonErrors(w, *models.NewErrorJson(500, "ERROR!! Internal Server Error"))
		return
	}
	client := NewClient(connection, server)
	server.AddClient(client)
	go client.ReadMessages()
	go client.WriteMessages()

}

// HERE fin l handler ghadi yt9ad and we'll be handling everything
func (server *ChatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.URL.Path {
	case "/ws/chat":
		server.ChatServerHandler(w, r)
		return
	case "/ws/users":
		w.Write([]byte("hhhhhhhhhhh"))
		return
	default:
		w.Write([]byte("not found !!"))
		return
	}

}

// f had l7ala we need to return 400
func isHandshakeError(err error) bool {
	return strings.Contains(err.Error(), "not a websocket handshake")
}
