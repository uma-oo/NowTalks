package handler

import (
	"fmt"
	"real-time-forum/backend/models"
)

func (server *ChatServer) AddClient(client *Client) {
	server.Lock()
	defer server.Unlock()
	server.clients[client.userId] = append(server.clients[client.userId], client)
}

func (server *ChatServer) RemoveClient(client *Client) {
	server.Lock()
	defer server.Unlock()
	client.CloseOnce.Do(func() {
		close(client.Message)
		close(client.ErrorJson)
	})
	if _, ok := server.clients[client.userId]; ok {
		client.connection.Close()
		deleteConnection(server.clients, client.userId, client)
	}
}

// first time working with channels and they seem great :!
func (client *Client) ReadMessages() {
	defer client.chatServer.RemoveClient(client)

	for {
		message := &models.Message{}
		err := client.connection.ReadJSON(&message)
		if err != nil {
			close(client.Done)
			break
		}
		message, errJson := client.chatServer.service.ValidateMessage(message)
		if errJson != nil {
			client.ErrorJson <- errJson
			continue
		}
		fmt.Println(message)
		client.Message <- message
	}

}

// i used the channels buy not sure if this is the correct way to handle this 

func (client *Client) WriteMessages() {
	defer client.chatServer.RemoveClient(client)

	for {
		select {
		case message := <-client.Message:
			err := client.connection.WriteJSON(message)
			if err != nil {
				return
			}
		case errJson := <-client.ErrorJson:
			err := client.connection.WriteJSON(errJson)
			if err != nil {
				return
			}
		case <-client.Done:
			return
		}
	}

}


// other than only sending the message back l khuna li sift l mssg 
// broadcastih l connections dyal nfss l user 
// broadcasti l message l receiver (message mashi errJson channel )


// to be done 
// create another endpoint for the online users (hakka bghit if it can be possible)

// finish the like/ dislike things by tomorrow

// in this case 
// 1-  alrady writed  in the same connection 
// 2 need to write to other connections of the same user if it's a message and l reciver 7ta huwa 
func BroadCastTheMessage(sender *Client , receiver *Client, message *models.Message) {
     
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
