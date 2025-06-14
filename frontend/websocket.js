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
    let users = document.getElementsByClassName('chat-user-card')
    // let users = document.querySelectorAll('.chat-user-card')
    console.log(users)
    console.log(Array.from(users)[0].dataset)
    
    // users.array.forEach(user => {
    //     console.log("card: ", user)
    // });
}