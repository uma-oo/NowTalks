import { getPostsApi } from "../api/posts.js"
import { navigateTo, createElement } from "../../utils.js"
import { createButton } from "./button.js"
import { createPostCard } from "./postCard.js"
import { createForm } from "./form.js"
import { PostForm } from "../const/forms.js"
import { throttledScrollFetcher } from "../utils.js"
import { createFilterContainer } from "./filter.js"
import { createIcon } from "./icon.js"

export function createPostsSection() {
    let postsSection = createElement('section', "posts_section")

    // post creation elements
    let createPostFormContainer = createElement('div', 'create-post-form-container')
    let addPostBtn = createButton({ icon: "plus" }, 'button', "add-post-btn")
    addPostBtn.addEventListener('click', (e) => {
        toggleCreatePostFormContainer(createPostFormContainer)
    })

    // post filter elements
    let filterContainer = createFilterContainer()
    let filterBtn = createButton({ icon: "filter" }, 'button', "filter-btn")
    filterBtn.addEventListener('click', (e) => {
        toggleFilterContainer(filterContainer)
    })

    let postsContainer = createElement('div', 'posts_container')
    postsContainer.dataset.offset = 0
    fetchPosts(postsContainer)
    const throttledScrollHandler = throttledScrollFetcher(fetchPosts)
    postsContainer.addEventListener('scroll', throttledScrollHandler)
    postsSection.append(postsContainer, createPostFormContainer, filterContainer, filterBtn, addPostBtn)
    return postsSection
}

function fetchPosts(container) {
    let offset = container.dataset.offset
    let filterData = {
        categories: container.dataset.categories,
        likedPosts: container.dataset.likedPosts,
        createdPosts: container.dataset.createdPosts
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

export function toggleCreatePostFormContainer() {
    let container = document.querySelector('.create-post-form-container')
    container.classList.toggle("create-post-form-container_expanded")
    if (!container.querySelector("#create-post-form")) {
        let title = createElement('h2', null , "Share your thoughts:")
        let goBack = createIcon("arrow-square-left")

        goBack.addEventListener('click', ()=> toggleCreatePostFormContainer())
        title.prepend(goBack)
        container.append(title, createForm(PostForm, "create-post-form"))
    } else {
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