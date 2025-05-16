import { timeAgo } from "../../utils.js"
import { Comments } from "../const/data.js"
import { CommentForm } from "../const/forms.js"
import { createButton } from "./button.js"
import { createComment } from "./comment.js"
import { createForm } from "./form.js"

export function createPostCard(postData) {
    let postContainer = document.createElement('div')
    postContainer.className = 'postContainer'

    let postHeader = document.createElement('div')
    postHeader.className = 'post-header'
    let postTitle = document.createElement('p')
    postTitle.className = 'post-title'
    postTitle.textContent = postData.title
    postHeader.append(postTitle)

    let postBody = document.createElement('div')
    postBody.className = 'post-body'
    let postContent = document.createElement('p')
    postContent.className = 'post-content'
    postContent.textContent = postData.content
    postBody.append(postContent)

    let postFooter = document.createElement('div')
    postFooter.className = 'post-Footer'
    let postWriter = document.createElement('p')
    postWriter.textContent = postData.user_name
    let postTimePosted = document.createElement('p')
    postTimePosted.textContent = timeAgo(postData.created_at)
    postFooter.append(postWriter,postTimePosted)


    let postCommentsContainer = document.createElement('div')
    postCommentsContainer.classList.add("post-comments-container","toggleable","hide")
    postCommentsContainer.append("Comments")

    Comments.forEach(comment => {
        postCommentsContainer.append(createComment(comment))
    });

    let viewPostBtn = createButton("viewPost >>",'button',["linkBtn", "viewPost","toggleable"])
    viewPostBtn.addEventListener('click', (e)=>togglePost(e.target, postContainer))
    let closeBtn = createButton("<- Go Back.","button",["close-btn","hide","toggleable"])
    closeBtn.addEventListener('click', (e)=>togglePost(e.target, postContainer))

    let commentForm = createForm(CommentForm,"comment-form")
    commentForm.classList.add("hide","toggleable")
    console.log(commentForm)

    postContainer.append(postHeader,postBody,postFooter,postCommentsContainer,commentForm,viewPostBtn,closeBtn)
    return postContainer
}


function togglePost(btnClicked, postContainer) {
    postContainer.classList.toggle("post-container_expand")
    let elementsToHide = postContainer.querySelectorAll(".toggleable")
    elementsToHide.forEach(elem => {
        elem.classList.toggle("hide")
    });
}





function expandPost(viewPostBtn, postContainer) {
    postContainer.classList.add("post-container_expand")

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
    postContainer.classList.remove("post-container_expand")
    let elementsToHide = postContainer.querySelectorAll(".show")
    elementsToHide.forEach(elem => {
        elem.classList.replace("show","hide")
    });
    viewPostBtn.classList.replace("hide","show")

}













