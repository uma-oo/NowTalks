import { users } from "../const/data.js"
import { createChatUserCard } from "./chatUserCard.js"
import { createIcon } from "./icon.js"

export function createChatSections(){
    let chatSection = document.createElement('section')
    chatSection.classList.add("chat_section","tab_section")


    let chatSectionHeader = document.createElement('div')
    chatSectionHeader.classList.add("chat_section-header")
    let chatsIcon = createIcon("chats")
    let appUser = document.createElement('h2')
    appUser.append("Chats: ")
    chatSectionHeader.append(appUser)

    let chatUsersCardsContainer = document.createElement('div')
    chatUsersCardsContainer.classList.add("chat_users_cards-container")

    users.forEach(user => {
        chatUsersCardsContainer.append(createChatUserCard(user))
    });

    chatSection.append(chatSectionHeader,chatUsersCardsContainer)
    return chatSection;
}