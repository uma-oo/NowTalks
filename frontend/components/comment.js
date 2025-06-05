import { createElement, timeAgo } from "../utils.js"
import { createIcon } from "./icon.js"


export function createComment({
    user_name,
    created_at,
    content
}) {
    let commentContainer = createElement('div', "comment-container")
    let commentContent = createElement('div', "comment-content")
    let commentWriter = createElement('span', null, user_name)
    let commentCreatedAt = createElement('span', null, timeAgo(created_at))
    let commentText = createElement('p', null, content)
    let heartIcon = createIcon("heart")

    commentWriter.prepend(createIcon("user"))
    commentCreatedAt.prepend(createIcon("calendar"))
    commentContent.append(commentWriter, commentCreatedAt, commentText)
    commentContainer.append(commentContent, heartIcon)
    return commentContainer
}
