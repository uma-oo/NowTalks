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
    // socket.send =
    // socket.onmessage()
}



export function closeConnection() {
    socket.close()
}


function receiveMessage(event) {
    let data = JSON.parse(event.data)
    switch (data.type) {
        case "message":
            console.log("hhhhhhhhhh");
            createChatMessageContainer(data, document.querySelector(".chat-window_expanded .chat-window-body"), "bottom")
            ReorderUsers(data)
            break;
        case "read":
            break;
        case "typing":
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