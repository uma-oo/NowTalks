import { timeAgo } from "../utils.js"
import { Comments } from "../const/data.js"
import { CommentForm } from "../const/forms.js"
import { createElement } from "../utils.js"
import { createButton } from "./button.js"
import { createComment } from "./comment.js"
import { createForm } from "./form.js"
import { createIcon } from "./icon.js"




export function createPostCard({
    id,
    user_name,
    title,
    content,
    categories,
    created_at,
    total_comments,
    total_likes
}) {
    let container = createElement('div', 'post-container')
    container.dataset.id = id
    let postHeader = createElement('div', 'post-header')
    let postInfo = createElement('div', 'post-info')
    let postTitle = createElement('p', 'post-title', title)
    let postWriter = createElement('span', null, `${user_name}`)
    let timestamp = createElement('span', null, timeAgo(created_at))
    let categoriesList = createElement('div', 'categories')

    if (categories) {
        categories.forEach(category => {
            let categoryTag = createElement('span', 'tag', `${category}`)
            categoriesList.append(categoryTag)
        });
    }

    let postBody = createElement('div', 'post-body')
    let postContent = createElement('p', 'post-content', content)
    let postFooter = createElement('div', 'post-footer')

    let reactions = [
        { type: "like", icon: "heart", "total": total_likes },
        { type: "comment", icon: "comment", "total": total_comments },
    ]

    let reactionElements = reactions.map(({ type, icon, total }) => {
        let text = type === "like" ? `Like` : total === 0 ? "Add comment" : "See comments"
        let containerElem = createElement('div', 'reaction-container',text);
        containerElem.dataset.reaction = type;
        let iconElem = createIcon(icon);
        let countElem = createElement('span', null, '0');
        
        containerElem.prepend(iconElem, countElem);
        return containerElem;
    });

    let postCommentsSection = createElement('div', "post-comments-section")
    let commentsContainer = createElement('div', "comments-container")
    let comments = Comments.map(comment => createComment(comment));

    let seeMore = createElement('p', "see-more", "See more posts...")

    let commentForm = createForm(CommentForm, "comment-form")

    postWriter.prepend(createIcon('user'))
    timestamp.prepend(createIcon('calendar'))
    postInfo.append(postWriter, timestamp)
    postHeader.append(postInfo, categoriesList, postTitle)
    postBody.append(postContent)
    postFooter.append(...reactionElements);
    commentsContainer.append(...comments)
    postCommentsSection.append(commentsContainer,seeMore ,commentForm)
    container.append(postHeader, postBody, postCommentsSection, postFooter)
    return container
}