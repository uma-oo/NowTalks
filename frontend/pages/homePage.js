import { isLoggedIn } from "../api/user.js";
import { createChatSection } from "../components/chatSection.js";
import { createHeader } from "../components/header.js";
import { createPostsSection } from "../components/postsSection.js";
import { navigateTo, createElement } from "../utils.js";



export function renderHomePage(app) {
    isLoggedIn().then(data => {
        if (data.is_logged_in) {
            app.dataset.nickname = data.nickname
            app.dataset.id = data.id

            let header = createHeader()
            let main = createElement('main',"home-main")
            let aside = createElement('aside', "chats-container")
            let containersWrapper = createElement('div', "containers-wrapper")
            let chatWindowSection =  createElement('div', "chat-window")
            containersWrapper.append(createPostsSection(),chatWindowSection)
            aside.append(...createChatSection())
            main.append(aside,containersWrapper)
            app.append(header,main)
        } else {navigateTo("/login")}
    })
}



