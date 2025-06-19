import { getComments } from "../api/comment.js";
import { CommentForm } from "../const/forms.js";
import { createElement, navigateTo } from "../utils.js";
import { createForm } from "./form.js";
import { createComment } from "./comment.js";
import { renderErrorPage } from "../pages/errorPage.js";

export function createPostCommentsSection(postId) {
    let postCommentsSection = createElement('div', "post-comments-section")
    let commentsContainer = createElement('div', "comments-container")
    commentsContainer.dataset.offset = 0
    let seeMore = createElement('p', "see-more", "See more comments ...")
    let commentForm = createForm(CommentForm, "comment-form")
    commentForm.dataset.postId = postId

    fetchComments(postId, commentsContainer)
    seeMore.addEventListener("click", () => {
        commentsContainer.dataset.offset = + commentsContainer.dataset.offset +10
        fetchComments(postId, commentsContainer)
    })
    postCommentsSection.append(commentsContainer, seeMore, commentForm)
    return postCommentsSection
}




function fetchComments(id, commentsContainer) {
    let offset = commentsContainer.dataset.offset
    getComments(id, offset).then(([status, data]) => {
        if (status === 401) {
            navigateTo("/login")
        }
        if ([400,429].includes(status)){
            renderErrorPage(status)
        }
        if (status === 200 && data) {
            data.forEach(commentData => {
                commentsContainer.append(createComment(commentData))
            });
        }
    })

}