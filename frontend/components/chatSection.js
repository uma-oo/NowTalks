import { users } from "../const/data.js"
import { createElement } from "../utils.js"
import { createChatUserCard } from "./chatUserCard.js"
import { openChatWindow } from "./chatWindow.js"
import { createIcon } from "./icon.js"

export function createChatSection() {
    let chatSectionHeader = createElement('div', "chats-section-header")
    let chatSectionHeaderTitle = createElement('h2', null, "Chats: ")
    chatSectionHeaderTitle.prepend(createIcon("chats"))

    let chatList = createElement('div', 'chat-list')
    let chats = users.map(user => {
        let userCard = createChatUserCard(user)
        let userCardClone = userCard.cloneNode(true)
        userCard.addEventListener("click",e => openChatWindow(userCard, userCardClone))
        return userCard
    });

    chatSectionHeader.append(chatSectionHeaderTitle)
    chatList.append(...chats)
    return [chatSectionHeader, chatList];
}