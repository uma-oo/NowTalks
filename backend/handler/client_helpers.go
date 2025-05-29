package handler

import (
	"encoding/json"
	"fmt"
	"real-time-forum/backend/models"
)

func (server *ChatServer) AddClient(client *Client, userId int) {
	server.Lock()
	defer server.Unlock()
	server.clients[userId] = append(server.clients[userId], client)
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
		_, payload, err := client.connection.ReadMessage()
		if err != nil {
			break
		}
		message := models.Message{}
		json.Unmarshal(payload, &message)
		fmt.Println("payload", string(payload))

	}

}

func (client *Client) WriteMessages() {

}
