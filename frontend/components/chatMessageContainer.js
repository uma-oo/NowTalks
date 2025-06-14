import { createElement, timeAgo } from "../utils.js";


export function createChatMessageContainer(
    {
        sender_username,
        content,
        created_at
    },
    openChatWindow,
    position = "top") {
    
    let chatMessageContainer = createElement('div', `chat-message-container ${sessionStorage.getItem("userNickname") === sender_username ? "align-self-end" : ""}`)
    let messageBubble = createElement('div', 'message-bubble')
    let sender = createElement('p', 'message-sender', sender_username)
    let messageContent = createElement('p', 'message-content', content)
    let timeStamp = createElement('span', null, timeAgo(created_at))

    messageBubble.append(sender, messageContent, timeStamp)
    chatMessageContainer.append(messageBubble)
    
    let messagesContainer = openChatWindow.querySelector(".chat-window-body")
    if (position == "top") {
        messagesContainer?.prepend(chatMessageContainer)
    }  else if (position === "bottom") {
        messagesContainer?.append(chatMessageContainer)
    }
}