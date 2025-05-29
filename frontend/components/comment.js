import {createElement, timeAgo} from "../utils.js"


export function createComment(data) {
    let commentContainer = createElement('div', "comment-container")
    let commentHeader = createElement('p', "comment-header", data.username)
    let commentBody = createElement('p', "comment-body", data.content)
    let commentFooter = createElement('div', 'comment-footer')

    let createdAt = createElement('p',null, timeAgo(data.createdAt) )
    document.createElement('p')

    createdAt.textContent = timeAgo(data.createdAt)
    commentFooter.append(createdAt)
    commentContainer.append(commentHeader,commentBody,commentFooter)
    return commentContainer
}
