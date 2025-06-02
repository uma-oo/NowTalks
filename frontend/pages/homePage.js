import { getCategories } from "../api/posts.js";
import { isLoggedIn } from "../api/user.js";
import { createButton } from "../components/button.js";
import { createChatSection } from "../components/chatSection.js";
import { createHeader } from "../components/header.js";
import { createPostsSection, toggleCreatePostFormContainer } from "../components/postsSection.js";
import { navigateTo, createElement } from "../utils.js";



export function renderHomePage(app) {
    isLoggedIn().then(data => {
        if (data.is_logged_in) {
            app.dataset.nickname = data.nickname
            app.dataset.id = data.id
            setUpWebsocket()
            setCategories(app).then(() => {
                let header = createHeader()
                let main = createElement('main', "home-main")
                let aside = createElement('aside', "chats-container")
                let createPostBtn = createButton({ text: "Create Post", icon: "edit" }, 'button', 'create-post-btn')
                let chatWindowSection = createElement('div', "chat-window")

                createPostBtn.addEventListener("click", ()=>{
                    if (!document.querySelector(".create-post-form-container_expanded")) {
                        toggleCreatePostFormContainer()
                    }
                })

                aside.append(createPostBtn, ...createChatSection())
                main.append(createPostsSection(), chatWindowSection)
                app.append(header, aside, main)
            })
        } else { navigateTo("/login") }
    })
}



async function setCategories(app) {
    app.dataset.categories = ''
    return getCategories().then(([status, data]) => {
        console.log(status, data)
        if (status == 200) {
            data.forEach(element => {
                app.dataset.categories += `${element.category_id}-${element.category_name},`
            });
        } else if (status == 401) {
            navigateTo('/login')
        }
    }).catch(error => console.error(error))
}

function setUpWebsocket() {
    let socket = new WebSocket("ws://localhost:8080/ws/chat");

    socket.onopen = function (e) {
        console.log("[open] Connection established");
        console.log("Sending to server");
        socket.send("My name is John");
    };
}