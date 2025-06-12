import { fetchMessages } from "../components/messagesSection.js";
import { MessageForm } from "../const/forms.js";
import { createElement } from "../utils.js";
import { createButton } from "./button.js";
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

    // const observer = new IntersectionObserver(
    //     entries => {
    //         entires.forEach(entry => {
    //             entry.target.innerText = `${Math.round(entry.intersectionRatio * 100)}%`
    //         })
    //     },
    //     { threshold: [0, 0.25, 0.5, 0.75, 1] }
    // )

    // observer.observe(document.getElementById("test"))





    chatUserCard.dataset.open = "true"
    chatWindow.innerHTML = ""
    let chatWindowHeader = createElement('div', 'chat-window-header')
    let goBackBtn = createButton({ icon: "arrow-square-left" })
    let chatWindowBody = createElement('div', 'chat-window-body')
    chatWindowBody.dataset.last = 0

    fetchMessages(0, chatUserCard.dataset.id, chatWindowBody)
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
    chatWindow.classList.remove("chat-window_expanded")
    chatUserCard.dataset.open = ""
    chatWindow.innerHTML = ""
}