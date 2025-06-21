import { getPostsApi } from "/frontend/api/posts.js"
import { navigateTo, createElement } from "/frontend/../utils.js"
import { createPostCard } from "/frontend/components/postCard.js"
import { createForm } from "/frontend/components/form.js"
import { PostForm } from "/frontend/const/forms.js"
// import { throttle, throttledScrollFetcher } from "/frontend/utils.js"
import { createIcon } from "/frontend/components/icon.js"
import { throttle } from "/frontend/utils.js"
import { renderErrorPage } from "/frontend/pages/errorPage.js"

export function createPostsSection() {
    let postsSection = createElement('section', "posts_section")
    let createPostFormContainer = createElement('div', 'create-post-form-container')
    let postsContainer = createElement('div', 'posts_container')
    postsContainer.dataset.canFetch = "true"
    postsContainer.dataset.offset = 0
    let fetchObserverTarget = createElement("div", null)

    postsContainer.append(fetchObserverTarget)
    postsSection.append(postsContainer, createPostFormContainer)

    postsSectionObserver(postsContainer, fetchObserverTarget)
    return postsSection
}

function postsSectionObserver(container, target) {
    const throttledFetch = throttle((container,offset) => fetchPosts(container, offset), 8000);
    const observer = new IntersectionObserver(
        (entries, observer) => {
            entries.map(entry => {
                let lastPost = target.previousSibling
                let offset = lastPost?.dataset.postId || 0
                if (container.dataset.canFetch === "false") {
                    console.log("can't fetch no more")
                    observer.unobserve(entry.target)
                }
                if (entry.isIntersecting) {
                    fetchPosts(container, offset)
                    // throttledFetch(container,offset);
                }
            })
        })
    observer.observe(target)
}

function fetchPosts(container, offset) {
    console.log("posts container: ", offset)
    getPostsApi(offset).then(([status, data]) => {
        if (status == 401) {
            navigateTo('login')
        } else if ([400,429,500].includes(status)) {
            renderErrorPage(status)
        }
        else if (status == 200) {
            if (!data || data.length < 10) {
                container.dataset.canFetch = "false"
            }
            if (!data) {
                console.log("no data", data)
                let img = document.createElement("img")
                img.src = "../assets/icons/finishline.png"
                img.style.width = "100px"
                let text = createElement("p", null, "You have reached the end :)")
                text.style.fontWeight = 600
                text.style.fontSize = "20px"
                container.append(img, text)
            } else {
                data.map(postData => container.insertBefore(createPostCard(postData), container.lastChild))
            }
        }
    })
}


export function toggleCreatePostFormContainer() {
    let container = document.querySelector('.create-post-form-container')
    container.classList.toggle("create-post-form-container_expanded")

    if (!container.querySelector("#create-post-form")) {
        let title = createElement('h2', null, "Share your thoughts:")
        let goBack = createIcon("arrow-square-left")
        goBack.addEventListener('click', () => toggleCreatePostFormContainer())
        title.prepend(goBack)
        container.append(title, createForm(PostForm, "create-post-form"))
    } else {
        container.innerHTML = ""
    }
}

