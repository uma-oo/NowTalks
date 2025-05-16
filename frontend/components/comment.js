import {timeAgo} from "../utils.js"


export function createComment(data) {
    let commentContainer = document.createElement('div')
    commentContainer.className = 'commentContainer'
    
    let commentHeader = document.createElement('p')
    commentHeader.classList.add("comment-header")
    commentHeader.textContent = data.username
    
    
    let commentBody = document.createElement('p')
    commentBody.classList.add("comment-body")
    commentBody.textContent = data.content

    let commentFooter = document.createElement('div')
    commentFooter.classList.add("comment-footer")
    let createdAt = document.createElement('p')  
    createdAt.textContent = timeAgo(data.createdAt)
    commentFooter.append(createdAt)

    commentContainer.append(commentHeader,commentBody,commentFooter)
    return commentContainer
}
