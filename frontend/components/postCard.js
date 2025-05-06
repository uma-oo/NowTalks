import { formatCreationDate } from "../utils.js"

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
    postWriter.textContent = postData.user
    postTimePosted.textContent = formatCreationDate(postData.created_at)


    postContainer.append(postHeader)
    postContainer.append(postBody)
    postContainer.append(postFooter)
    return postContainer
}