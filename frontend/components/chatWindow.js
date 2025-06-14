import { fetchMessages } from "../components/messagesSection.js";
import { MessageForm } from "../const/forms.js";
import { createElement } from "../utils.js";
import { createButton } from "./button.js";
import { createForm } from "./form.js";




export function openChatWindow(chatUserCard, chatUserCardClone) {
    let chatWindow = document.querySelector('.chat-window')
    chatWindow.dataset.id = chatUserCard.dataset.id
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
    chatWindowBody.dataset.last = 0
    let targetTopElement = createElement('div',"observer-target top-observer-target ","top Observer target")
    let targetBottomElement = createElement('div',"observer-target bottom-observer-target","bottom Observer target") 

    
    let chatWindowFooter = createElement('div', 'chat-window-footer')
    let messageform = createForm(MessageForm, "message-form")
    
    goBackBtn.addEventListener('click', _ => {
        closeChatWindow(chatUserCard, chatWindow)
    })
    
    chatWindowHeader.append(goBackBtn, chatUserCardClone)
    chatWindowBody.append(targetTopElement,targetBottomElement)
    chatWindowFooter.append(messageform)
    chatWindow.append(chatWindowHeader, chatWindowBody, chatWindowFooter)

    fetchMessages(0, chatUserCard.dataset.id, chatWindow)
    chatWindowObservers(chatWindow,targetTopElement,targetBottomElement)
    return chatWindow
}

export function closeChatWindow(chatUserCard, chatWindow) {
    chatWindow.classList.remove("chat-window_expanded")
    chatUserCard.dataset.open = ""
    chatWindow.innerHTML = ""
}

function chatWindowObservers(container,targetTopElement,targetBottomElement) {
    const topObserver = new IntersectionObserver(
        entries => {
            entries.forEach(entry => {
                console.log("Top Entery",entry)
                entry.target.innerText = "fetch old messages"
            })
        }
    )


    const bottomObserver = new  IntersectionObserver(
        entries => {
            entries.forEach(entry => {
                console.log("Bottom Entery",entry)
                entry.target.innerText = "fetch new messages"
            })
        }
    )

    topObserver.observe(targetTopElement)
    bottomObserver.observe(targetBottomElement)
}