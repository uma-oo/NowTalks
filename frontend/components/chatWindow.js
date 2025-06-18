
import { fetchMessages } from "../components/messagesSection.js";
import { MessageForm } from "../const/forms.js";
import { createElement, navigateTo } from "../utils.js";
import { sendMessage } from "../websocket.js";
import { createButton } from "./button.js";
import { createForm } from "./form.js";

export function openChatWindow(chatUserCard) {

    if (chatUserCard.dataset.open) { // if chat already open
        return
    }

    const previousOpendChat = document.querySelector('.chat-list > [data-open = "true"]');
    if (previousOpendChat) { // if there is an already open chat
        previousOpendChat.dataset.open = "";
    }

    // open new chat
    let user = chatUserCard.dataset
    let chatWindow = document.querySelector('.chat-window')
    chatWindow.classList.add("chat-window_expanded")
    chatWindow.dataset.id = user.id
    chatWindow.dataset.firstFetch = "true"
    chatWindow.dataset.topObsorver = "on"
    let notificationsContainer = chatUserCard.querySelector(".notification_container")
    let notificationsCounter = notificationsContainer.querySelector(".user_notifications")
    notificationsContainer.classList.add("hide")
    
    if (user.notifications != 0) {
        sendMessage("read", "read")
    }

    notificationsCounter.textContent = 0
    chatUserCard.dataset.open = "true"
    sessionStorage.setItem("openChat", +user.id)
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
        }
    )
    topObserver.observe(targetTopElement)
}

