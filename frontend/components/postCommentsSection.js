import { getComments } from "/frontend/api/comment.js";
import { CommentForm } from "/frontend/const/forms.js";
import { createElement, navigateTo } from "/frontend/utils.js";
import { createForm } from "/frontend/components/form.js";
import { createComment } from "/frontend/components/comment.js";
import { renderErrorPage } from "/frontend/pages/errorPage.js";

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
        if (status === 200) {
            if (!data && !commentsContainer.children.length) {
                let noContent = createElement("div", "no-content")
                let icon = createIcon("no-comment")
                icon.style.width = "74px"
                icon.style.opacity = "0.8px"
                let noMessages = createElement("p", null, "No Comments for this post - be the first to comment.")
                noContent.append(icon, noMessages)
                commentsContainer.append(noContent)
                // commentsContainer.
            }
            data?.forEach(commentData => {
                commentsContainer.append(createComment(commentData))
            });
        }
    })

}