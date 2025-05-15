import { getPostsApi } from "../api/posts.js"
import { navigateTo } from "../../utils.js"
import { createButton } from "./button.js"
import { createPostCard } from "./postCard.js"

export function createPostsSections() {
    let postsSection = document.createElement('section')
    postsSection.classList.add("posts_section")

    let postsContainer = document.createElement('div')
    postsContainer.classList.add("posts_Container")

    let addPostBtn = createButton("+", 'button', "")
    addPostBtn.addEventListener('click', (e) => {
    })
    let posts = getPostsApi()
    posts.then(data => {
        if ( data?.status == 401){
            navigateTo('/login')
        } else {
            postsContainer.append(...createPostCards(data))
        }
    })
    postsSection.append(postsContainer, addPostBtn)
    return postsSection
}

function createPostCards(data){
    if (data == null) return "No Posts Available"
    
    return data.map(postData => createPostCard(postData))
}
