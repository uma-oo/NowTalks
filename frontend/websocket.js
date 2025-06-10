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
    return socket
}



function receiveMessage(event) {
        console.log(event.data);
}


function sendMessage() {
    // Construct a msg object containing the data the server needs to process the message from the chat client.
    const msg = {
        type: "message",
        text: document.getElementById("text").value,
        id: clientID,
        date: Date.now(),
    };

    // Send the msg object as a JSON-formatted string.
    socket.send(JSON.stringify(msg));

    // Blank the text input element, ready to receive the next line of text from the user.
    document.getElementById("text").value = "";
}