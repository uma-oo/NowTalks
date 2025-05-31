import { getCategories } from "../api/posts.js";
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
            setCategories(app).then(()=>{
                let header = createHeader()
                let main = createElement('main', "home-main")
                let aside = createElement('aside', "chats-container")
                let containersWrapper = createElement('div', "containers-wrapper")
                let chatWindowSection = createElement('div', "chat-window")
                containersWrapper.append(createPostsSection(), chatWindowSection)
                aside.append(...createChatSection())
                main.append(aside, containersWrapper)
                app.append(header, main)
            })
        } else { navigateTo("/login") }
    })
}



function setCategories(app) {
    app.dataset.categories = ''
    return getCategories().then(([status, data]) => {
        console.log(status, data)
        if (status == 200) {
            data.forEach(element => {
                app.dataset.categories +=`${element.category_id}-${element.category_name},`
            });
        } else if (status == 401) {
            navigateTo('/login')
        }
    }).catch(error => console.error(error))
}