
import { MessageForm } from "/frontend/const/forms.js";
import { createElement, navigateTo, throttle } from "/frontend/utils.js";
import { sendMessage } from "/frontend/websocket.js";
import { createButton } from "/frontend/components/button.js";
import { createForm } from "/frontend/components/form.js";
import { createChatMessageContainer } from "./chatMessageContainer.js";
import { getMessages } from "/frontend/api/messages.js";
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
    let targetTopElement = createElement('div', "observer-target top-observer-target ")

    let chatWindowFooter = createElement('div', 'chat-window-footer')
    let messageform = createForm(MessageForm, "message-form")

    goBackBtn.addEventListener('click', () => {
        let chatUserCard = document.querySelector("[data-open = 'true']")
        let chatWindow = document.querySelector(".chat-window_expanded")
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
    console.log("chatUser: ", chatUserCard)
}

function chatWindowObserver(container, target) {
    const throttledFetch = throttle((offset, id, type, container) => {
        fetchMessages(offset, id, type, container)
    }, 8000);
    console.log(throttledFetch)

    const topObserver = new IntersectionObserver(
        (entries, observer) => {
            entries.forEach(entry => {
                let chatData = container.dataset
                let nextSibling = target.nextSibling
                let offset = nextSibling?.dataset.messageId || 0
                if (chatData.topObsorver === "off") {
                    console.log("unobserve the topTarget")
                    observer.unobserve(entry.target)
                };
                if (entry.isIntersecting) {
                    throttledFetch(offset, chatData.id, "old", container)
                }
            })
        }
    )
    topObserver.observe(target)
}

function fetchMessages(offset, receiver_id, type, chatWindow) {
    let chatWindowBody = chatWindow.querySelector('.chat-window-body');
    const prevScrollHeight = chatWindowBody.scrollHeight
    getMessages(offset, receiver_id, type)
        .then(([status, data]) => {
            if (status === 401) {
                navigateTo("/login")
            }
            if (status === 200) {
                if (!data || data.length < 10) {
                    chatWindow.dataset.topObsorver = "off"
                }
                data?.forEach(message => {
                    createChatMessageContainer(message, chatWindow, "top")
                });
                const diff = chatWindowBody.scrollHeight - prevScrollHeight;
                chatWindowBody.scrollTop += diff;
            }
            if ([400, 429, 500].includes(status)) {
                renderErrorPage(status)
            }
        })
}