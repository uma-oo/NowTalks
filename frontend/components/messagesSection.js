import { getMessages } from "../api/messages.js";
import { navigateTo } from "../utils.js";
import { createChatMessageContainer } from "./chatMessageContainer.js";

export function fetchMessages(offset, receiver_id, type, chatWindow) {
    getMessages(offset, receiver_id).then(([status, data]) => {
        if (status === 401) {
            navigateTo("/login")
        }
        if (status === 400) {
            console.log(data);
        }
        if (status === 200 && data) {
            data.forEach(message => {
                createChatMessageContainer(message, chatWindow)
            });
            // messagesContainer.dataset.offset = +messagesContainer.dataset.offset + 10
        }
    })

}