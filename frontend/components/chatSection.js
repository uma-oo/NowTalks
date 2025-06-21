
import { getUsers } from "/frontend/api/user.js"
import { createElement, navigateTo } from "/frontend/utils.js"
import { createChatUserCard } from "/frontend/components/chatUserCard.js"
import { openChatWindow } from "/frontend/components/chatWindow.js"
import { createIcon } from "/frontend/components/icon.js"

export async function createChatSection() {
    let chatSection = createElement('div', 'chat-section')
    let chatSectionHeader = createElement('div', "chats-section-header")
    let chatSectionHeaderTitle = createElement('h2', null, "Messages: ")
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
    if (status == 401) {
        navigateTo("/login")
    }

    if ([400, 429, 500].includes(status)) {
        renderErrorPage(status)
    }

    if (status == 200) {
        let chats = data?.map(userData => {
            let userCard = createChatUserCard(userData)
            let userCardClone = userCard.cloneNode(true)
            userCard.addEventListener("click", () => {
                openChatWindow(userCard, userCardClone)
            })
            return userCard
        });
        if (chats) {
            chatList.innerHTML = ""
            chatList.append(...chats)
        } else {
            chatList.append("No users")
        }
    }
}