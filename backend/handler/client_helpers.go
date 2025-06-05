package handler

import (
	"io"

	"real-time-forum/backend/models"

	"github.com/gorilla/websocket"
)

var NUMBER = 0

func (server *ChatServer) AddClient(client *Client) {
	server.Lock()
	server.clients[client.userId] = append(server.clients[client.userId], client)
	defer server.Unlock()
}

func (server *ChatServer) RemoveClient(client *Client) {
	server.Lock()
	defer server.Unlock()
	// client.CloseOnce.Do(func() {
	// 	close(client.Message)
	// 	close(client.ErrorJson)
	// 	close(client.Done)
	// })
	if _, ok := server.clients[client.userId]; ok {
		client.connection.Close()
		deleteConnection(server.clients, client.userId, client)
		go server.BroadCastOnlineStatus()
	}
}

// first time working with channels and they seem great :!
func (client *Client) ReadMessages() {
	defer client.chatServer.RemoveClient(client)

	for {
		message := &models.Message{}
		err := client.connection.ReadJSON(&message)
		if err != nil {
			if err == io.EOF {
				client.ErrorJson <- &models.ErrorJson{
					Status: 400,
					Message: models.MessageErr{
						Message:    "ERROR!! Empty Message field",
						ReceiverID: "ERROR!! Empty Receiver Id field",
					},
				}
				continue
			}
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				break
			}
			// continue
		}

		message.SenderID = client.userId

		message_validated, errJson := client.chatServer.service.ValidateMessage(message)
		if errJson != nil {
			client.ErrorJson <- errJson
			continue
		}

		client.Message <- message_validated
		client.BroadCastTheMessage(message_validated)
	}
}

// i used the channels buy not sure if this is the correct way to handle this

func (client *Client) WriteMessages() {
	defer client.chatServer.RemoveClient(client)

	for {
		select {
		case errJson := <-client.ErrorJson:
			err := client.connection.WriteJSON(errJson)
			if err != nil {
				return
			}
		case message := <-client.Message:
			err := client.connection.WriteJSON(message)
			if err != nil {
				return
			}
		case online_users := <-client.OnlineUsers:
			err := client.connection.WriteJSON(online_users)
			if err != nil {
				return
			}
			// case <-client.Done:
			// 	return
		}
	}
}

func (sender *Client) BroadCastTheMessage(message *models.Message) {
	// braodcast to the connections dyal sender
	sender.chatServer.Lock()
	defer sender.chatServer.Unlock()
	for _, conn := range sender.chatServer.clients[sender.userId] {
		if conn.connection != sender.connection {
			conn.Message <- message
		}
	}
	// dyal receiver
	for _, value := range sender.chatServer.clients[message.ReceiverID] {
		value.Message <- message
	}
}

// dummy way to delete a connection but i'm done
func deleteConnection(clientList map[int][]*Client, userId int, client_to_be_deleted *Client) {
	index := -1
	for i, value := range clientList[userId] {
		if value == client_to_be_deleted {
			index = i
			break
		}
	}
	if index != -1 {
		clientList[userId] = append(clientList[userId][:index], clientList[userId][index+1:]...)
	}
}

// let's do it inside another function and make it specific to the client
func (server *ChatServer) BroadCastOnlineStatus() {
	server.Lock()
	defer server.Unlock()
	online_users := []models.User{}
	for _, connections := range server.clients {
		if len(connections) != 0 {
			online_users = append(online_users, models.User{Id: connections[0].userId, Nickname: connections[0].Username})
		}
	}

	for _, connections := range server.clients {
		for _, conn := range connections {
			conn.OnlineUsers <- online_users
		}
	}
}
