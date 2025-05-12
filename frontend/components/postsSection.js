// import { getPosts } from "../api/posts.js"
import { Posts } from "../const/data.js"
import { createButton } from "./button.js"
import { createPostCard } from "./postCard.js"

export async function createPostsSections(){
    let postsSection = document.createElement('section')
    postsSection.classList.add("posts_section", "tab_section","visible_tab_section") 
    
    let addPostBtn = createButton("+", 'button',"")
    
    addPostBtn.addEventListener('click',(e)=>{
        
    })

    // let posts = await getPosts()

    Posts.forEach(post => {
        postsSection.append(createPostCard(post))
    });

    return postsSection

}