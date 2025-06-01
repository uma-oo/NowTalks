import { timeAgo } from "../utils.js"
import { Comments } from "../const/data.js"
import { CommentForm } from "../const/forms.js"
import { createElement } from "../utils.js"
import { createButton } from "./button.js"
import { createComment } from "./comment.js"
import { createForm } from "./form.js"
import { createIcon } from "./icon.js"


let reactions = [
    { type: "likes", icon: "like" },
    { type: "dislikes", icon: "dislike" },
    { type: "comments", icon: "comment" },
]

export function createPostCard({
    id,
    user_name,
    title,
    content,
    categories,
    created_at,
    total_comments,
    total_likes,
    total_dislikes
}) {
    let container = createElement('div', 'post-container')
    container.dataset.id = id

    let postHeader = createElement('div', 'post-header')
    let postInfo = createElement('div', 'post-info')
    let postTitle = createElement('p', 'post-title', title)
    let postWriter = createElement('span', null, `${user_name}`)
    let timestamp = createElement('span', null, timeAgo(created_at))
    let categoriesList = createElement('div', 'categories')
    categories.forEach(category => {
        let categoryTag = createElement('span', 'tag', `${category}`)
        // catego
        categoriesList.append(categoryTag)
    });

    let postBody = createElement('div', 'post-body')
    let postContent = createElement('p', 'post-content', content)

    let postFooter = createElement('div', 'post-footer')
    let reactionElements = reactions.map(reaction => {
        let container = createElement('div', 'reaction-container');
        container.dataset.reaction = reaction.type;
        let icon = createIcon(reaction.icon);
        let count = createElement('span', null, '0');
        container.append(icon, count);
        return container;
    });


    let postCommentsContainer = createElement('div', "post-comments-container toggleable hide")
    let comments = Comments.map(comment => createComment(comment));

    let viewPostBtn = createButton({ text: "see post", icon: "arrowright" }, 'button', "linkBtn viewPost toggleable row-reverse")
    let closeBtn = createButton({ text: "close", icon: "xmark" }, "button", "close-btn hide toggleable")

    viewPostBtn.addEventListener('click', (e) => togglePost(e.target, container))
    closeBtn.addEventListener('click', (e) => togglePost(e.target, container))

    let commentForm = createForm(CommentForm, "comment-form")
    commentForm.classList.add("hide", "toggleable")

    postWriter.prepend(createIcon('user'))
    timestamp.prepend(createIcon('calendar'))
    postInfo.append(postWriter, timestamp)
    postHeader.append(postInfo,categoriesList,postTitle)
    postBody.append(postContent)
    postFooter.append(...reactionElements);
    postCommentsContainer.append("Comments", ...comments)
    container.append(postHeader, postBody, postFooter, postCommentsContainer, commentForm, viewPostBtn, closeBtn)
    return container
}

function togglePost(button, container) {
    container.classList.toggle('post-container_expand')
    let toggleableElements = container.querySelectorAll('.toggleable')
    toggleableElements.forEach(elem => elem.classList.toggle('hide'))
    container.scrollIntoView({ block: "center" })
}
