export function setUpWebsocket() {
    let socket = new WebSocket("ws://localhost:8080/ws/chat");
    socket.onopen = function (e) {
        console.log("[open] Connection established");
        console.log("Sending to server");
        // socket.send(JSON.stringify({
        //     message : "hi"
        // }));
    };
    // socket.onmessage()
    return socket
}