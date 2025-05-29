package handler

import (
	"net/http"
)

func (server *ChatServer) ChatServerHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := server.upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.Write([]byte("salaam"))
	}
	client := NewClient(connection, server)
    server.AddClient(client)
	go client.ReadMessages()
	go client.WriteMessages()
	
}

// HERE fin l handler ghadi yt9ad and we'll be handling everything
func (server *ChatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ws/chat":
		server.ChatServerHandler(w, r)
	case "/ws/users":
		w.Write([]byte("hhhhhhhhhhh"))
	}


}
