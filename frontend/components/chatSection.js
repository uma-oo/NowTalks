import { users } from "../const/data.js"
import { createChatUserCard } from "./chatUserCard.js"

export function createChatSections(){
    let chatSection = document.createElement('section')
    chatSection.classList.add("chat_section","tab_section")


    let chatSectionHeader = document.createElement('div')
    chatSectionHeader.classList.add("chat_section-header")
    let appUser = document.createElement('p')
    appUser.append("Chat To Other Users")
    chatSectionHeader.append(appUser)

    let chatUsersCardsContainer = document.createElement('div')
    chatUsersCardsContainer.classList.add("chat_users_cards-container")

    users.forEach(user => {
        chatUsersCardsContainer.append(createChatUserCard(user))
    });

    chatSection.append(chatUsersCardsContainer)
    return chatSection;
}