import { createChatMessageContainer } from "./components/chatMessageContainer.js"
import { navigateTo, ReorderUsers } from "./utils.js";

let socket = null

export function setUpWebsocket() {
    if (!socket) {
        socket = new WebSocket("ws://localhost:8080/ws/chat");
    } 
    socket.onopen = function (e) {
        console.log("websocket connection established: ",e, socket)
    };
    socket.onmessage = (event) => receiveMessage(event)

    socket.onclose = (event) => {
        socket = null
        if (window.location.pathname !== "/login") navigateTo('login')
    }
}

export function closeConnection() {
    socket.close(1000,"user logged out")
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
            changeUsersStatus(data.data)
            break;
        case "typing":
            break;
        default:
            break;
    }
}

export function sendMessage(messageContent) {
    let receiver_id = parseInt(document.querySelector(".chat-window_expanded").dataset.id)
    const msg = {
        content: messageContent,
        type: "message",
        receiver_id: receiver_id,
        created_at: new Date(Date.now()),
    };
    socket.send(JSON.stringify(msg));
}


function changeUsersStatus(data) {
    let onlineUsers = data.map(user=>user.id)
    let usersCards = document.querySelectorAll('.chat-user-card')
    usersCards.forEach(userCard => {
        let id = userCard.dataset.id
        if (onlineUsers.includes(+id)) {
            userCard.dataset.status = "online"
            userCard.querySelector('.user_status').textContent = "online"
        } else {
            userCard.dataset.status = "offline"
            userCard.querySelector('.user_status').textContent = "offline"
        }
    });
}