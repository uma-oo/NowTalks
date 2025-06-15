import { getUsers } from "../api/user.js"
import { createElement, navigateTo } from "../utils.js"
import { createChatUserCard } from "./chatUserCard.js"
import { openChatWindow } from "./chatWindow.js"
import { createIcon } from "./icon.js"

export async function createChatSection() {
    let chatSection = createElement('div','chat-section')
    let chatSectionHeader = createElement('div', "chats-section-header")
    let chatSectionHeaderTitle = createElement('h2', null, "Chats: ")
    chatSectionHeaderTitle.prepend(createIcon("chats"))

    let chatList = createElement('div', 'chat-list')
    chatList.dataset.offset = 0

    await fetchUsers(chatList)
    chatSectionHeader.append(chatSectionHeaderTitle)
    chatSection.append(chatSectionHeader, chatList)
    return chatSection;
}


export async function fetchUsers(chatList) {
    let offset = chatList.dataset.offset
    let [status, data] = await getUsers(offset)
    // getUsers(offset).then(([status, data]) => {
        if (status == 401) {
            navigateTo("/login")
        }
        if (status == 200) {
            let chats = data?.map(userData => {
                let userCard = createChatUserCard(userData)
                userCard.dataset.id = userData.id
                let userCardClone = userCard.cloneNode(true)
                userCard.addEventListener("click", _ => {
                    openChatWindow(userCard, userCardClone)
                })
                return userCard
            });
            if (chats) {
                chatList.append(...chats)
            } else {
                chatList.append("No users")
            }
        }
    // })

}