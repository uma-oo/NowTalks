import { getCategories } from "/frontend/api/posts.js";
import { isLoggedIn } from "/frontend/api/user.js";
import { createButton } from "/frontend/components/button.js";
import { createChatSection } from "/frontend/components/chatSection.js";
import { createHeader } from "/frontend/components/header.js";
import { createPostsSection, toggleCreatePostFormContainer } from "/frontend/components/postsSection.js";
import { navigateTo, createElement } from "/frontend/utils.js";
import { setUpWebsocket } from "/frontend/websocket.js";
import { renderErrorPage } from "/frontend/pages/errorPage.js";

export function renderHomePage(app) {
    isLoggedIn().then(data => {
        if (data.is_logged_in) {
            sessionStorage.setItem("userId", data.id)
            sessionStorage.setItem("userNickname", data.nickname)
            sessionStorage.setItem("onlineUsers", [])
            sessionStorage.setItem("openChat", null)
            setCategories(app).then(async () => {
                let header = createHeader()
                let main = createElement('main', "home-main")
                let aside = createElement('aside', "chats-container")
                let createPostBtn = createButton({ text: "Create Post", icon: "edit" }, 'button', 'create-post-btn')
                let chatWindowSection = createElement('div', "chat-window")

                createPostBtn.addEventListener("click", () => {
                    if (!document.querySelector(".create-post-form-container_expanded")) {
                        toggleCreatePostFormContainer()
                    }
                })
                aside.append(createPostBtn, await createChatSection())
                main.append(createPostsSection(), chatWindowSection)
                app.append(header, aside, main)
                setUpWebsocket()
            })
        } else { navigateTo("/login") }
    })
}

async function setCategories(app) {
    app.dataset.categories = ''
    return getCategories().then(([status, data]) => {

        if (status == 200) {
            data.forEach(element => {
                app.dataset.categories += `${element.category_id}-${element.category_name},`
            });
        } else if (status == 401) {
            navigateTo('login')
        } else if ([429,500].includes(status)) {
            renderErrorPage(status)
        }
    }).catch(error => console.error(error))
}
