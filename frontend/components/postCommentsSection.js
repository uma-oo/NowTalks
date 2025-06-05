import { getComments } from "../api/comment.js";
import { CommentForm } from "../const/forms.js";
import { createElement, navigateTo } from "../utils.js";
import { createForm } from "./form.js";
import { createComment } from "./comment.js";

export function createPostCommentsSection(postId) {

    let postCommentsSection = createElement('div', "post-comments-section")
    let commentsContainer = createElement('div', "comments-container")
    commentsContainer.dataset.offset = 0
    let seeMore = createElement('p', "see-more", "See more posts...")
    let commentForm = createForm(CommentForm, "comment-form")



    fetchComments(postId, commentsContainer)

    // commentsContainer.append(...comments)
    postCommentsSection.append(commentsContainer, seeMore, commentForm)

    return postCommentsSection
}

function fetchComments(id, commentsContainer) {
    console.info("Fetching comments for postId", id)
    let offset = commentsContainer.dataset.offset
    getComments(id, offset).then(([status, data])=>{
        if (status === 401 ) {
            navigateTo("/")
        } 
        if (status === 200 && data) {
            console.log(data)
            data.forEach(commentData => {
                commentsContainer.append(createComment(commentData))
            });
        } 
    })

}