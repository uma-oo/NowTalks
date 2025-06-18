
import { fetchMessages } from "../components/messagesSection.js";
import { MessageForm } from "../const/forms.js";
import { createElement } from "../utils.js";
import { sendMessage } from "../websocket.js";
import { createButton } from "./button.js";
import { createForm } from "./form.js";

export function openChatWindow(chatUserCard) {
    let user = chatUserCard.dataset
    console.log(user);
    let notificationsContainer = chatUserCard.querySelector(".notification_container")

    notificationsContainer.classList.add("hide")
    let chatWindow = document.querySelector('.chat-window')
    chatWindow.dataset.id = user.id
    chatWindow.dataset.firstFetch = "true"
    chatWindow.dataset.topObsorver = "on"

    chatWindow.classList.add("chat-window_expanded")
    if (chatUserCard.dataset.open) {
        return
    }
    console.log(+notificationsContainer.querySelector("span").textContent != 0, +notificationsContainer.querySelector("span").textContent);
    if (+notificationsContainer.querySelector("span").textContent != 0) {
        sendMessage("read", "read")
    }
    notificationsContainer.querySelector("span").textContent = 0

    const previousOpendChat = document.querySelector('.chat-list > [data-open = "true"]');
    if (previousOpendChat) {
        // here you can mark the messages for the previous chat read
        previousOpendChat.dataset.open = "";
    }

    chatUserCard.dataset.open = "true"
    chatWindow.innerHTML = ""

    let chatWindowHeader = createElement('div', 'chat-window-header')
    let goBackBtn = createButton({ icon: "arrow-square-left" })
    let receiver = createElement('h3', null, user.userName)

    let chatWindowBody = createElement('div', 'chat-window-body')
    let targetTopElement = createElement('div', "observer-target top-observer-target ", "top Observer target")

    let chatWindowFooter = createElement('div', 'chat-window-footer')
    let messageform = createForm(MessageForm, "message-form")

    goBackBtn.addEventListener('click', () => {
        closeChatWindow(chatUserCard, chatWindow)
    })

    chatWindowHeader.append(goBackBtn, receiver)
    chatWindowBody.append(targetTopElement)
    chatWindowFooter.append(messageform)
    chatWindow.append(chatWindowHeader, chatWindowBody, chatWindowFooter)

    chatWindowObserver(chatWindow, targetTopElement)
    return chatWindow
}

export function closeChatWindow(chatUserCard, chatWindow) {
    chatWindow.innerHTML = ""
    chatWindow.classList.remove("chat-window_expanded")
    chatUserCard.dataset.open = ""
}

function chatWindowObserver(container, targetTopElement) {
    const topObserver = new IntersectionObserver(
        (entries, observer) => {
            entries.forEach(entry => {
                let chatData = container.dataset
                let nextSibling = targetTopElement.nextSibling
                let offset = nextSibling?.dataset.messageId || 0
                if (chatData.topObsorver === "off") {
                    console.log("unobserve the topTarget")
                    observer.unobserve(entry.target)
                };
                if (entry.isIntersecting) {
                    fetchMessages(offset, chatData.id, "old", container)
                }
            })
        }, { rootMargin: "20px" }
    )
    topObserver.observe(targetTopElement)
}