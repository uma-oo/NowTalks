import { getPostsApi } from "../api/posts.js"
import { navigateTo } from "../../utils.js"
import { createButton } from "./button.js"
import { createPostCard } from "./postCard.js"
import { createForm } from "./form.js"
import { PostForm } from "../const/forms.js"

export function createPostsSections() {
    let postsSection = document.createElement('section')
    postsSection.classList.add("posts_section")

    let postsContainer = document.createElement('div')
    postsContainer.classList.add("posts_Container")

    let createPostFormContainer = document.createElement('div')
    createPostFormContainer.classList.add('create-Post-Form-Container')

    let addPostBtn = createButton("+", 'button', "")
    addPostBtn.addEventListener('click', (e) => {
        toggleCreatePostFormContainer(createPostFormContainer)
    })

    let posts = getPostsApi()
    posts.then(data => {
        if ( data?.status == 401){
            navigateTo('/login')
        } else {
            postsContainer.append(...createPostCards(data))
        }
    })
    postsSection.append(postsContainer, createPostFormContainer ,addPostBtn)
    return postsSection
}

function toggleCreatePostFormContainer(container) {
    container.classList.toggle("create-Post-Form-Container_expanded")
    console.log(container.querySelector("#create-post-form"))
    if (!container.querySelector("#create-post-form") ) {
        container.append(createForm(PostForm,"create-post-form"))
    } else {
        container.innerHTML = ""
    }
}

function createPostCards(data){
    if (data == null) return "No Posts Available"
    return data.map(postData => createPostCard(postData))
}