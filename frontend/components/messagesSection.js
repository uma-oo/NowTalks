import { getMessages } from "../api/messages.js";
import { navigateTo } from "../utils.js";
import { createChatMessageContainer } from "./chatMessageContainer.js";

export function fetchMessages(offset, receiver_id, type, chatWindow) {
    if (type === "new") {
        console.log("Fetching new messages ")
    } else {
        console.log("Fetching old messages ")
    }


    getMessages(offset, receiver_id, type)
    .then(([status, data]) => {
        if (status === 401) {
            navigateTo("/login")
        }

        if (status === 400) {
            console.log(data);
        }

        if (status === 200) {
            console.log(data)
            
            if (!data || data.length < 10) {
                type === "old" ? chatWindow.dataset.topObsorver = "off" : chatWindow.dataset.bottomObserver = "off"
            }
            if (type=="old") data = data.reverse()
            data?.forEach(message => {
                createChatMessageContainer(message, chatWindow , type == "new" ? "bottom": "top")
            });
            // messagesContainer.dataset.offset = +messagesContainer.dataset.offset + 10
        }
    })

}