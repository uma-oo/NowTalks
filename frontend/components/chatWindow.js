import { MessageForm } from "../const/forms.js";
import { createElement } from "../utils.js";
import { createButton } from "./button.js";
import { createForm } from "./form.js";

let data = {
    user: "bob",
    status: "online",
    id: 1 
}

let messages = [
    {
        content: "Hi!",
        user: "bob",
        timeStamp : "2025-05-06T15:00:16Z"
    },
    {
        content: "Hey man, sup",
        user: "you",
        timeStamp : "2025-05-06T15:00:16Z"
    },
    {
        content: "all good, wbu ?",
        user: "bob",
        timeStamp : "2025-05-06T15:00:16Z"
    },
    {
        content: "wanna hang out some time this weekend?",
        user: "bob",
        timeStamp : "2025-05-06T15:00:16Z"
    },
    {
        content: "For sure man, send me the details later.",
        user: "you",
        timeStamp : "2025-05-06T15:00:16Z"
    }
]


export function openChatWindow(chatUserCard, chatUserCardClone) {
    let chatWindow = document.querySelector('.chat-window')
    chatWindow.classList.add("chat-window_expanded")
    console.log(chatUserCard.dataset.open)
    if (chatUserCard.dataset.open) {
        return
    } 

    const opendChat = document.querySelector('.chat-list > [data-open = "true"]');
    if (opendChat) opendChat.dataset.open = "";

    chatUserCard.dataset.open = "true"
    chatWindow.innerHTML = ""
    let chatWindowHeader  = createElement('div', 'chat-window-header')

    let chatWindowBody = createElement('div', 'chat-window-body')
    let goBackBtn = createButton({icon: "arrowleft"})
    let chatWindowFooter = createElement('div', 'chat-window-footer')
    let messageform = createForm(MessageForm,"message-form")


    goBackBtn.

    chatWindowHeader.append(goBackBtn,chatUserCardClone)
    chatWindowBody.append("Here goes message :)")
    chatWindowFooter.append(messageform)
    chatWindow.append(chatWindowHeader,chatWindowBody,chatWindowFooter)
    return chatWindow
}

export function closeChatWindow() {
    let chatWindow = document.querySelector('.chat-window')
    chatWindow.innerHTML = ""
    chatWindow.classList.remove(".chat-window_expanded")
}