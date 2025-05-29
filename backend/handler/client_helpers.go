package handler

import (
	"fmt"

	"github.com/gorilla/websocket"
)

func (server *ChatServer) AddClient(client *Client) {
	server.Lock()
	defer server.Unlock()
	server.clients[client] = true
}

func (server *ChatServer) RemoveClient(client *Client) {
	server.Lock()
	defer server.Unlock()
	if _, ok := server.clients[client]; ok {
		client.connection.Close()
		delete(server.clients, client)
	}
}

func (client *Client) ReadMessages() {
    defer client.chatServer.RemoveClient(client)

	for {
		messageType, payload, err := client.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseAbnormalClosure, websocket.CloseGoingAway) {
				fmt.Println(err)
			}
			break
		}
		fmt.Println(messageType)
		fmt.Println("payload", string(payload))

	}
}

func (client *Client) WriteMessages() {

}
