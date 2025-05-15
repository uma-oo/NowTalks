import { getPostsApi } from "../api/posts.js"
// import { Posts } from "../const/data.js"
import { renderApp } from "../index.js"
import { navigateTo } from "../../utils.js"
import { createButton } from "./button.js"
import { createPostCard } from "./postCard.js"

export function createPostsSections() {
    let postsSection = document.createElement('section')
    postsSection.classList.add("posts_section", "tab_section", "visible_tab_section")

    let addPostBtn = createButton("+", 'button', "")

    addPostBtn.addEventListener('click', (e) => {

    })

    getPostsApi().then(posts => {
        if (posts.status == 401) navigateTo('/login')
    }).catch(error => {
        console.log('error catched')
        console.error('Error fetching posts:', error);
    });

    // if (posts.status && posts.status == 401) {
    //     await navigateTo('/login')
    //     return;
    // }

    // posts.forEach(post => {
    //     postsSection.append(createPostCard(post))
    // });



    return postsSection

}