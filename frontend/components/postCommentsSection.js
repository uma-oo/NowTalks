import { getComments } from "/frontend/api/comment.js";
import { CommentForm } from "/frontend/const/forms.js";
import { createElement, navigateTo } from "/frontend/utils.js";
import { createForm } from "/frontend/components/form.js";
import { createComment } from "/frontend/components/comment.js";
import { renderErrorPage } from "/frontend/pages/errorPage.js";
import { createIcon } from "/frontend/components/icon.js"

export function createPostCommentsSection(postId, comments) {
    let postCommentsSection = createElement('div', "post-comments-section")
    let commentsContainer = createElement('div', "comments-container")
    commentsContainer.dataset.offset = 0
    let commentForm = createForm(CommentForm, "comment-form")
    commentForm.dataset.postId = postId
    let seeMore

    if (comments > 0) {
        seeMore = createElement('p', "see-more", "See more comments ...")
        seeMore.addEventListener("click", (e) => {
            fetchComments(postId, commentsContainer)
        })
    }

    fetchComments(postId, commentsContainer)

    postCommentsSection.append(commentsContainer, seeMore || "", commentForm)
    return postCommentsSection
}

function fetchComments(id, commentsContainer) {

    let offset = commentsContainer.lastChild?.dataset.commentId
    getComments(id, offset || 0).then(([status, data]) => {
        if (status === 401) {
            navigateTo("/login")
        }
        if ([400, 429].includes(status)) {
            renderErrorPage(status)
        }
        if (status === 200) {
            if (!data && !commentsContainer.children.length) {
                let noContent = createElement("div", "no-content")
                let icon = createIcon("no-comment")
                let noMessages = createElement("p", null, "No Comments for this post - be the first to comment.")
                noContent.append(icon, noMessages)
                commentsContainer.append(noContent)
                return
            }
            if (!data || data.length < 10) {
                let seeMore = commentsContainer.nextSibling
                if (seeMore) seeMore.remove()
            }
            data?.forEach(commentData => {
                commentsContainer.append(createComment(commentData))
            });
        }
    })

}