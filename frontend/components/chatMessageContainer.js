import { createElement, formatTimestamp, timeAgo } from "/frontend/utils.js";


export function createChatMessageContainer (
    {
        message_id,
        sender_username,
        content,
        created_at
    },
    openChatWindow,
    position = "top") {
    let messageBubble = createElement('div', `message-bubble ${sessionStorage.getItem("userNickname") === sender_username ? "align-self-end" : ""}`)
    messageBubble.dataset.messageId = message_id

    let chatWindowBody = openChatWindow.querySelector(".chat-window-body")
    let topTargetObserver = chatWindowBody.querySelector(".top-observer-target ")


    let sender = createElement('p', 'message-sender', sender_username)
    let messageContent = createElement('p', 'message-content', content)
    let timeStamp = createElement('span', null, formatTimestamp(created_at))

    messageBubble.append(sender, messageContent, timeStamp)

    if (position == "top") {
        chatWindowBody.insertBefore(messageBubble, topTargetObserver.nextSibling);
    } else if (position === "bottom") {
        chatWindowBody.append(messageBubble);
        if (messageBubble.classList.contains("align-self-end")) messageBubble.scrollIntoView()
    }
}