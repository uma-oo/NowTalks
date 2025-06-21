import { ToggleLike } from "/frontend/api/reaction.js"
import { createElement, timeAgo } from "/frontend/utils.js"
import { createIcon } from "/frontend/components/icon.js"


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

    let reaction = createElement('div', 'reaction-container')
    let heartIcon = createIcon("heart", "like")
    let count = createElement('span', null, total_likes ? total_likes : '0')
    
    if (liked != 0) {
        heartIcon.style.fill = "red"
    }
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
    reaction.append(heartIcon, count)
    commentContent.append(commentWriter, commentCreatedAt, commentText)
    commentContainer.append(commentContent, reaction)
    return commentContainer
}
