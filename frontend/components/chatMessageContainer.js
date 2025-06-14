import { createElement, formatTimestamp, timeAgo } from "../utils.js";


export function createChatMessageContainer(
    {
        sender_username,
        content,
        created_at
    },
    openChatWindow,
    position = "top") {
    let chatWindowBody = openChatWindow.querySelector(".chat-window-body")
    let topTargetObserver = chatWindowBody.querySelector(".top-observer-target ")
    let bottomTargetObserver = chatWindowBody.querySelector(".bottom-observer-target ")


    let chatMessageContainer = createElement('div', `chat-message-container ${sessionStorage.getItem("userNickname") === sender_username ? "align-self-end" : ""}`)
    let messageBubble = createElement('div', 'message-bubble')
    let sender = createElement('p', 'message-sender', sender_username)
    let messageContent = createElement('p', 'message-content', content)
    let timeStamp = createElement('span', null, formatTimestamp(created_at))

    messageBubble.append(sender, messageContent, timeStamp)
    chatMessageContainer.append(messageBubble)
    
    if (position == "top") {
        chatWindowBody.insertBefore(chatMessageContainer, topTargetObserver.nextSibling);
    }  else if (position === "bottom") {
        chatWindowBody.insertBefore(chatMessageContainer, bottomTargetObserver);
        if (chatMessageContainer.classList.contains("align-self-end")) chatMessageContainer.scrollIntoView()
    }
}