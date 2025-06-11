import { getMessages } from "../api/messages.js";
import { navigateTo } from "../utils.js";
import { createChatMessageContainer } from "./chatMessageContainer.js";







export function fetchMessages(offset, receiver_id, MessageContainer) {
    // let offset = messagesContainer.dataset.offset
    // console.log("offset", offset);
    // let receiver_id = document.querySelector(".chat-window_expanded [data-id]").dataset.id
    getMessages(offset, receiver_id).then(([status, data]) => {

        if (status === 401) {
            navigateTo("/login")
        }
        if (status === 400) {
            console.log(data);
        }
        if (status === 200 && data) {
            data.forEach(message => {
                createChatMessageContainer(message, MessageContainer)
            });
        }
    })

}