import { createElement, timeAgo } from "../utils.js";

export function createChatMessageContainer({
    content,
    user,
    timestamp
}) {

    let chatMessageContainer = createElement('div', `chat-message-container ${user === "you" ? "align-self-end" : ""}`)
    let messageBubble = createElement('div', 'message-bubble')
    let sender = createElement('p', 'message-sender', user)
    let messageContent = createElement('p', 'message-content', content)
    let timeStamp = createElement('span', null, timeAgo(timestamp))

    messageBubble.append(sender, messageContent)
    chatMessageContainer.append(messageBubble, timeStamp)

    return chatMessageContainer
}