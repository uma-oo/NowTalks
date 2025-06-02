import { createIcon } from "./icon.js";
import { createElement } from "../utils.js";

export function createChatUserCard(nickname) {
    let chatUserCard = createElement('div','chat-user-card') 
    chatUserCard.dataset.open = ""

    let chatUserCardHeader = createElement('div', 'chat-user-card-header online') 
    let chatUserCardBody = createElement('div', 'chat-user-card-body') 
    let chatUserCardFooter =  createElement('div', 'chat-user-card-footer') 
    let userName = createElement('p', null, nickname)
    
    // use it for time for latest message or is typing message 
    
    chatUserCardBody.append(userName)
    chatUserCard.append(chatUserCardHeader, chatUserCardBody,chatUserCardFooter)
    return chatUserCard
}