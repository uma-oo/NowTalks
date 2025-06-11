let socket


export function setUpWebsocket() {
    socket = new WebSocket("ws://localhost:8080/ws/chat");
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
    
    console.log(event.data.type);
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