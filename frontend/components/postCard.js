import { timeAgo } from "../utils.js"
import { createButton } from "./button.js"

export function createPostCard(postData) {
    let postContainer = document.createElement('div')
    postContainer.className = 'postContainer'

    let postHeader = document.createElement('div')
    let postTitle = document.createElement('p')
    postHeader.append(postTitle)
    postHeader.className = 'post-header'
    postTitle.className = 'post-title'
    postTitle.textContent = postData.title

    let postBody = document.createElement('div')
    let postContent = document.createElement('p')
    postBody.append(postContent)
    postBody.className = 'post-body'
    postContent.className = 'post-content'
    postContent.textContent = postData.content

    let postFooter = document.createElement('div')
    let postWriter = document.createElement('p')
    let postTimePosted = document.createElement('p')
    postFooter.append(postWriter)
    postFooter.append(postTimePosted)
    postFooter.className = 'post-Footer'
    postWriter.textContent = postData.user_name
    postTimePosted.textContent = timeAgo(postData.created_at)

    let viewPostBtn = createButton("viewPost >>",'button',["linkBtn","show", "viewPost"])
    viewPostBtn.addEventListener('click', ()=>expandPost(viewPostBtn, postContainer))

    postContainer.append(postHeader,postBody,postFooter,viewPostBtn)
    return postContainer
}





function expandPost(viewPostBtn, postContainer) {
    postContainer.classList.add("expand-post")
    let elementsToHide = postContainer.querySelectorAll(".hide")
    elementsToHide.forEach(elem => {
        elem.classList.replace("hide","show")
    });
    let closeBtn = createButton("<- Go Back.","button",["close-btn","show"])
    if (!postContainer.querySelector('.close-btn')) {
        postContainer.append(closeBtn)
        viewPostBtn.classList.replace("show","hide")
        closeBtn.addEventListener("click", ()=>shrinkPost(viewPostBtn,postContainer) )
    }
}



function shrinkPost(viewPostBtn, postContainer){
    postContainer.classList.remove("expand-post")
    let elementsToHide = postContainer.querySelectorAll(".show")
    elementsToHide.forEach(elem => {
        elem.classList.replace("show","hide")
    });
    viewPostBtn.classList.replace("hide","show")

}













