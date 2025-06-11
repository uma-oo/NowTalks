import { fetchMessages } from "../components/messagesSection.js";
import { MessageForm } from "../const/forms.js";
import { createElement } from "../utils.js";
import { createButton } from "./button.js";
import { createChatMessageContainer } from "./chatMessageContainer.js";
import { createForm } from "./form.js";




export function openChatWindow(chatUserCard, chatUserCardClone) {
    let chatWindow = document.querySelector('.chat-window')
    chatWindow.classList.add("chat-window_expanded")
    if (chatUserCard.dataset.open) {
        return
    }

    const opendChat = document.querySelector('.chat-list > [data-open = "true"]');

    if (opendChat) {
        opendChat.dataset.open = "";
    }

    chatUserCard.dataset.open = "true"
    chatWindow.innerHTML = ""
    let chatWindowHeader = createElement('div', 'chat-window-header')
    let goBackBtn = createButton({ icon: "arrow-square-left" })
    let chatWindowBody = createElement('div', 'chat-window-body')
    fetchMessages(30, chatUserCard.dataset.id, chatWindowBody)
    let chatWindowFooter = createElement('div', 'chat-window-footer')
    let messageform = createForm(MessageForm, "message-form")


    goBackBtn.addEventListener('click', _ => {
        closeChatWindow(chatUserCard, chatWindow)
    })

    chatWindowHeader.append(goBackBtn, chatUserCardClone)
    chatWindowFooter.append(messageform)
    chatWindow.append(chatWindowHeader, chatWindowBody, chatWindowFooter)
    return chatWindow
}

export function closeChatWindow(chatUserCard, chatWindow) {
    console.log(chatWindow)
    chatWindow.classList.remove("chat-window_expanded")
    chatUserCard.dataset.open = ""
    chatWindow.innerHTML = ""
}