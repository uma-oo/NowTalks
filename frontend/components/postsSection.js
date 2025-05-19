import { getPostsApi } from "../api/posts.js"
import { navigateTo } from "../../utils.js"
import { createButton } from "./button.js"
import { createPostCard } from "./postCard.js"
import { createForm } from "./form.js"
import { PostForm } from "../const/forms.js"
import { throttledScrollFetcher } from "../utils.js"
import { createFilterContainer } from "./filter.js"

export function createPostsSections() {
    let postsSection = document.createElement('section')
    postsSection.classList.add("posts_section")

    let postsContainer = document.createElement('div')
    postsContainer.classList.add("posts_Container")
    postsContainer.dataset.offset = 0

    let createPostFormContainer = document.createElement('div')
    createPostFormContainer.classList.add('create-Post-Form-Container')

    let addPostBtn = createButton({icon:"plus"}, 'button', "add-post-btn")
    addPostBtn.addEventListener('click', (e) => {
        toggleCreatePostFormContainer(createPostFormContainer)
    })

    let filterContainer = createFilterContainer()
    console.log(filterContainer)
    

    let filterBtn = createButton({icon:"filter"}, 'button', "filter-btn")
    filterBtn.addEventListener('click', (e)=> {
        toggleFilterContainer()
    })

    fetchPosts(postsContainer)
    const throttledScrollHandler = throttledScrollFetcher(fetchPosts)
    postsContainer.addEventListener('scroll', throttledScrollHandler)
    postsSection.append(postsContainer, createPostFormContainer, addPostBtn)
    return postsSection
}

function fetchPosts(container) {
    let offset = container.dataset.offset
    let filterData = {
        categories : container.dataset.categories,
        likedPosts : container.dataset.likedPosts,
        createdPosts : container.dataset.createdPosts
    }
    getPostsApi(filterData, offset).then(data => {
        if (data?.status == 401) {
            navigateTo('/login')
        } else if (data) {
            container.append(...createPostCards(data))
            container.dataset.offset = +container.dataset.offset + 10
        }
    }).catch(error => console.error(error))
}

function createPostCards(data) {
    if (data == null) return "No Posts Available"
    return data.map(postData => createPostCard(postData))
}

function toggleCreatePostFormContainer(container) {
    container.classList.toggle("create-Post-Form-Container_expanded")
    let bottonIcon = document.querySelector(".add-post-btn>i")
    console.log(bottonIcon)
    if (!container.querySelector("#create-post-form")) {
        container.append(createForm(PostForm, "create-post-form"))
        bottonIcon.classList.replace("fa-plus","fa-xmark")
    }else {
        bottonIcon.classList.replace("fa-xmark","fa-plus")
        container.innerHTML = ""
    }
}

function toggleFilterContainer(container) {
    container.classList.toggle("filter-container_expanded")
    let elementsToHide = container.querySelectorAll(".toggleable")
    elementsToHide.forEach(elem => {
        elem.classList.toggle("hide")
    });
}