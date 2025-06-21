import { getMessages } from "../api/messages.js";
import { navigateTo } from "../utils.js";
import { createChatMessageContainer } from "./chatMessageContainer.js";

export function fetchMessages(offset, receiver_id, type, chatWindow) {
    let chatWindowBody = chatWindow.querySelector('.chat-window-body');
    const prevScrollHeight = chatWindowBody.scrollHeight
    getMessages(offset, receiver_id, type)
        .then(([status, data]) => {
            if (status === 401) {
                navigateTo("/login")
            }
            if (status === 200) {
                if (!data || data.length < 10) {
                    chatWindow.dataset.topObsorver = "off"
                }
                data?.forEach(message => {
                    createChatMessageContainer(message, chatWindow, "top")
                });
                const diff = chatWindowBody.scrollHeight - prevScrollHeight;
                chatWindowBody.scrollTop += diff;
            }
            if ([400, 429, 500].includes(status)) {
                renderErrorPage(status)
            }
        })
}