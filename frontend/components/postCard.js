import { createPostCommentsSection } from "/frontend/components/postCommentsSection.js"
import { createIcon } from "/frontend/components/icon.js"
import { ToggleLike } from "/frontend/api/reaction.js"
import { createElement } from "/frontend/utils.js"
import { timeAgo } from "/frontend/utils.js"

export function createPostCard({
    id,
    user_name,
    title,
    content,
    categories,
    created_at,
    total_comments,
    total_likes,
    liked
}) {
    let container = createElement('div', 'post-container')
    container.dataset.postId = id
    let postHeader = createElement('div', 'post-header')
    let postInfo = createElement('div', 'post-info')
    let postWriter = createElement('span', null, `${user_name}`)
    let timestamp = createElement('span', null, timeAgo(created_at))
    let categoriesList = createElement('div', 'categories')
    let postTitle = createElement('p', 'post-title', title)
    let postBody = createElement('div', 'post-body')
    let postContent = createElement('p', 'post-content', content)
    let postFooter = createElement('div', 'post-footer')

    categories.forEach(category => {
        let categoryTag = createElement('span', 'tag', `${category}`)
        categoriesList.append(categoryTag)
    });

    let reactions = [
        { type: "like", icon: "heart", "total": total_likes },
        { type: "comment", icon: "comment", "total": total_comments },
    ]

    let reactionElements = reactions.map(({ type, icon, total }) => {
        let containerElem = createElement('div', 'reaction-container');
        containerElem.dataset.reaction = type;
        let iconElem = createIcon(icon, type);
        if (liked != 0 && type == "like") {
            iconElem.style.fill = "red";
        }
        let countElem = createElement('span', null, total ? total : '0');
        containerElem.prepend(iconElem, countElem);
        return containerElem;
    });

    let postCommentsSection = createPostCommentsSection(id,total_comments)
    postWriter.prepend(createIcon('user'))
    timestamp.prepend(createIcon('calendar'))
    postInfo.append(postWriter, timestamp)
    postHeader.append(postInfo, categoriesList, postTitle)
    postBody.append(postContent)
    postFooter.append(...reactionElements);
    container.append(postHeader, postBody, postCommentsSection, postFooter)

    let commentReaction = postFooter.querySelector('.reaction-container[data-reaction="comment"]')
    commentReaction.addEventListener("click", () => {
        container.scrollIntoView()
        postCommentsSection.classList.toggle("post-comments-section_expanded")
    })

    let likeReaction = postFooter.querySelector('.reaction-container[data-reaction="like"]')
    likeReaction.addEventListener("click", () => {
        ToggleLike({
            entity_id: id,
            entity_type: "post"
        }, likeReaction.querySelector("svg"), likeReaction.querySelector("span"))
    })

    return container
}




