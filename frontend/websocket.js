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
    console.log(event.data);
}


export function sendMessage(messageContent) {

    let receiver_id = parseInt(document.querySelector(".chat-window_expanded [data-id]").dataset.id)
    console.log(receiver_id);
    const msg = {
        type: "message",
        content: messageContent,
        created_at: Date.now(),
        receiver_id: receiver_id
    };

    // Send the msg object as a JSON-formatted string.
    socket.send(JSON.stringify(msg));
}