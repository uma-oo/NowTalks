import { isLoggedIn } from "../api/user.js";
import { createChatSections } from "../components/chatSection.js";
import { createHeader } from "../components/header.js";
import { createPostsSections } from "../components/postsSection.js";
import { navigateTo } from "../utils.js";



export function renderHomePage(app) {
    isLoggedIn().then(data => {
        if (data.is_logged_in) {
            app.dataset.nickname = data.nickname
            app.dataset.id = data.id
            let header = createHeader()
            let main =  document.createElement('main')
            main.classList.add("home-main")
            main.append(createChatSections(),createPostsSections() )
            app.append(header,main)
        } else {navigateTo("/login")}
    })
}



