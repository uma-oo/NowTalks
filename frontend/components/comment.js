import { ToggleLike } from "../api/reaction.js"
import { createElement, timeAgo } from "../utils.js"
import { createIcon } from "./icon.js"


export function createComment({
    id,
    user_name,
    created_at,
    content,
    total_likes,
    liked
}) {
    let commentContainer = createElement('div', "comment-container")
    let commentContent = createElement('div', "comment-content")
    let commentWriter = createElement('span', null, user_name)
    let commentCreatedAt = createElement('span', null, timeAgo(created_at))
    let commentText = createElement('p', null, content)
    let heartIcon = createIcon("heart", "like")
    if (liked != 0) {
        heartIcon.style.fill = "red"
    }
    let count = createElement('span', null, total_likes ? total_likes : '0')
    heartIcon.addEventListener("click", () => {
        ToggleLike(
            {
                entity_id: id,
                entity_type: "comment"
            },
            heartIcon, count
        )
    })
    commentWriter.prepend(createIcon("user"))
    commentCreatedAt.prepend(createIcon("calendar"))
    commentContent.append(commentWriter, commentCreatedAt, commentText)
    commentContainer.append(commentContent, heartIcon, count)
    return commentContainer
}
