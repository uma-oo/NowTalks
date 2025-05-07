// import {getPosts} from "./api/posts.js"
import { renderForm } from "./components/form.js"
import { createPostCard } from "./components/postCard.js"
import { posts } from "./const/data.js"
import { CommentForm, LoginForm, PostForm, registerFom } from "./const/forms.js"

export const app = document.getElementById('app')


// renderForm(app, registerFom, 'registerForm')
// renderForm(app, LoginForm, 'loginForm')
// renderForm(app, PostForm, 'postForm')
// renderForm(app, CommentForm, 'commentForm')



posts.forEach(post => app.append(createPostCard(post))) 

