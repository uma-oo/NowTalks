import { fetchMessages } from "../components/messagesSection.js";
import { MessageForm } from "../const/forms.js";
import { createElement } from "../utils.js";
import { createButton } from "./button.js";
import { createForm } from "./form.js";

export function openChatWindow(chatUserCard) {
// notification_container
    let user = chatUserCard.dataset
    chatUserCard.querySelector(".user_notifications").textContent = 0
    chatUserCard.querySelector(".notification_container").classList.add("hide")
    let chatWindow = document.querySelector('.chat-window')
    chatWindow.dataset.id = user.id
    chatWindow.dataset.notification = user.notifications

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
    let targetTopElement = createElement('div',"observer-target top-observer-target ","top Observer target")
    let targetBottomElement = createElement('div',"observer-target bottom-observer-target","bottom Observer target")
    
    let chatWindowFooter = createElement('div', 'chat-window-footer')
    let messageform = createForm(MessageForm, "message-form")
    
    goBackBtn.addEventListener('click', () => {
        closeChatWindow(chatUserCard, chatWindow)
    })

    chatWindowHeader.append(goBackBtn, receiver)
    chatWindowBody.append(targetTopElement,targetBottomElement)
    chatWindowFooter.append(messageform)
    chatWindow.append(chatWindowHeader, chatWindowBody, chatWindowFooter)

    chatWindowObservers(chatWindow,targetTopElement,targetBottomElement)
    return chatWindow
}

export function closeChatWindow(chatUserCard, chatWindow) {
    chatWindow.innerHTML = ""
    chatWindow.classList.remove("chat-window_expanded")
    chatUserCard.dataset.open = ""
}

function chatWindowObservers(container,targetTopElement,targetBottomElement) {

    console.log(container.dataset)

    const topObserver = new IntersectionObserver(
        entries => {
            //  fetch old message if there is a message after the top observer || if there is no notification
            entries.forEach(entry => {
                if (!entry.isIntersecting) retrun
                let nextSibling = targetBottomElement.nextSibling
                let offset = nextSibling.dataset.messageId 
                if (+container.dataset.id == 0 || offset) {
                    console.log("fitshing old messages")
                }
            })
        }
    )

    const bottomObserver = new  IntersectionObserver(
        entries => {
            //  fetch new message if there is still notifications
            entries.forEach(entry => {
                if (!entry.isIntersecting) retrun
                let previousSibling = targetBottomElement.previousSibling
                let offset = previousSibling.dataset.messageId
                if (+container.dataset.notification > 0) {
                    console.log()
                }
            })
        }
    )

    topObserver.observe(targetTopElement)
    bottomObserver.observe(targetBottomElement)
}