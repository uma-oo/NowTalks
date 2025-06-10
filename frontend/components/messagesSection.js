import { getMessages } from "../api/messages";
import { createChatMessageContainer } from "./chatMessageContainer";







function fetchMessagess(messagesContainer) {
    let offset = messagesContainer.dataset.offset
    let receiver_id = document.querySelector(".chat-window_expanded [data-id]").dataset.id
    getMessages(offset, receiver_id).then(([status, data]) => {
        if (status === 401) {
            navigateTo("/login")
        }
        if (status === 200 && data) {
            data.forEach(message => {
                commentsContainer.append(createChatMessageContainer(message))
            });
        }
    })

}