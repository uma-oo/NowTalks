import { fetchMessages } from "../components/messagesSection.js";
import { MessageForm } from "../const/forms.js";
import { createElement } from "../utils.js";
import { createButton } from "./button.js";
import { createForm } from "./form.js";

export function openChatWindow(chatUserCard) {
    let user = chatUserCard.dataset
    chatUserCard.querySelector(".user_notifications").textContent = 0
    chatUserCard.querySelector(".notification_container").classList.add("hide")
    let chatWindow = document.querySelector('.chat-window')
    chatWindow.dataset.id = user.id
    chatWindow.dataset.notifications = user.notifications
    chatWindow.dataset.topObsorver = "on"
    chatWindow.dataset.bottomObsorver = "on"

    chatWindow.classList.add("chat-window_expanded")
    if (chatUserCard.dataset.open) {
        return
    }

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
    let targetBottomElement = createElement('div', "observer-target bottom-observer-target", "bottom Observer target")

    let chatWindowFooter = createElement('div', 'chat-window-footer')
    let messageform = createForm(MessageForm, "message-form")

    goBackBtn.addEventListener('click', () => {
        closeChatWindow(chatUserCard, chatWindow)
    })

    chatWindowHeader.append(goBackBtn, receiver)
    chatWindowBody.append(targetTopElement, targetBottomElement)
    chatWindowFooter.append(messageform)
    chatWindow.append(chatWindowHeader, chatWindowBody, chatWindowFooter)

    chatWindowObservers(chatWindow, targetTopElement, targetBottomElement)
    return chatWindow
}

export function closeChatWindow(chatUserCard, chatWindow) {
    chatWindow.innerHTML = ""
    chatWindow.classList.remove("chat-window_expanded")
    chatUserCard.dataset.open = ""
}

function chatWindowObservers(container, targetTopElement, targetBottomElement) {
    const topObserver = new IntersectionObserver(
        (entries, observer) => {
            //  fetch old message if there is a message after the top observer || if there is no notification
            entries.forEach(entry => {
                if (!entry.isIntersecting) return;
                let chatData = container.dataset
                let notifications = +container.dataset.notifications
                let nextSibling = targetBottomElement.nextSibling
                let offset = nextSibling?.dataset.messageId || 0

                if (chatData.topObsorver === "off") {
                    console.log("unobserve the topTarget")
                    observer.unobserve(entry.target)
                };

                if (notifications == 0) {
                    fetchMessages(offset, chatData.id, "old", container)
                }
            })
        }
    )

    const bottomObserver = new IntersectionObserver(
        (entries, observer) => {
            //  fetch new message if there is still notifications
            entries.forEach(entry => {
                if (!entry.isIntersecting) return;
                let chatData = container.dataset
                let previousSibling = targetBottomElement.previousSibling
                let offset = +previousSibling?.dataset.messageId || 0
                // console.log(`notifications: ${chatData.notifications}, offset: ${offset}`)

                if (chatData.bottomObserver === "off") {
                    console.log("unobserve the bottomTarget")
                    observer.unobserve(entry.target)
                };
                if (chatData.notifications > 0) {
                    console.log("offset: ", offset)
                    fetchMessages( offset, chatData.id, "new", container)
                    // console.log("fitching new messages")
                }
            })
        }
    )

    topObserver.observe(targetTopElement)
    bottomObserver.observe(targetBottomElement)
}