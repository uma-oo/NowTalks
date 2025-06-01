import { createIcon } from "./icon.js";
import { createElement } from "../utils.js";

export function createChatUserCard(user) {
    let chatUserCard = createElement('div','chat-user-card') 
    chatUserCard.dataset.open = ""

    let chatUserCardHeader = createElement('div', 'chat-user-card-header') 
    let chatUserCardBody = document.createElement('div')
    chatUserCardBody.classList.add("chat-user-card_body")
    let userName = document.createElement('p')
    userName.append(user.name)
    chatUserCardBody.append(userName)

    // use it for time for latest message or is typing message 
    let chatUserCardFooter = document.createElement('div')
    chatUserCardFooter.classList.add("chat-user-card_footer")

    user.online ? chatUserCardHeader.classList.add('online') : chatUserCardHeader.classList.add('offline')
    chatUserCard.append(chatUserCardHeader, chatUserCardBody,chatUserCardFooter)
    return chatUserCard
}