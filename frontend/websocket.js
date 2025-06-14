import { createChatMessageContainer } from "./components/chatMessageContainer.js"
import { ReorderUsers } from "./utils.js";


let socket = null


export function setUpWebsocket() {
    if (!socket) {
        socket = new WebSocket("ws://localhost:8080/ws/chat");
    }
    socket.onopen = function (e) {
        console.log("[open] Connection established");
        console.log("Sending to server");
    };
    socket.onmessage = (event) => receiveMessage(event)

}



export function closeConnection() {
    socket.close()
}


function receiveMessage(event) {
    let data = JSON.parse(event.data)
    switch (data.type) {
        case "message":
            let openChatWindow = document.querySelector(`.chat-window_expanded[data-id="${data.sender_id}"]`) || 
            document.querySelector(`.chat-window_expanded[data-id="${data.receiver_id}"]`)
            if (openChatWindow) createChatMessageContainer(data, openChatWindow, "bottom")
            ReorderUsers(data)
            break;
        case "online":
            console.log(data)
            changeUsersStatus(data.data)
            break;
        case "typing":
            break;
        default:
            break;
    }
}


export function sendMessage(messageContent) {

    let receiver_id = parseInt(document.querySelector(".chat-window_expanded [data-id]").dataset.id)

    const msg = {
        content: messageContent,
        type: "message",
        receiver_id: receiver_id,
        created_at: new Date(Date.now()),
    };
    // Send the msg object as a JSON-formatted string.
    socket.send(JSON.stringify(msg));
}


function changeUsersStatus(data) {
    let onlineUsers = data.map(user=>user.id)
    console.log(onlineUsers)
    // let chatList = document.getElementsByClassName('chat-list')
    let chatList = document.querySelector('.chat-list')
    console.log(chatList)
    let usersCards = document.querySelectorAll('.chat-user-card')
    // let users = document.querySelectorAll('.chat-user-card')
    usersCards.forEach(userCard => {
        let id = userCard.dataset.id
        if (onlineUsers.includes(+id)) {
            userCard.dataset.status = "online"
            userCard.querySelector('.user_status').textContent = "online"
        }else {
            userCard.dataset.status = "offline"
            userCard.querySelector('.user_status').textContent = "offline"
        }
    });
    
    // users.array.forEach(user => {
    //     console.log("card: ", user)
    // });
}