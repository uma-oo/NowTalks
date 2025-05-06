// import {getPosts} from "./api/posts.js"
import { renderForm } from "./components/form.js"
import { CommentForm, LoginForm, PostForm, registerFom } from "./const/forms.js"

export const app = document.getElementById('app')




// const registerForm = document.getElementById("registerForm")
// registerForm.addEventListener('submit', e => {
//     e.preventDefault()
//     let form = new FormData(registerForm)
//     const data = Object.fromEntries(form.entries());
// })

//   fetch('https://fakestoreapi.com/products')
// .then(response => response.json())
// .then(data => console.log(data))

renderForm(app, registerFom, 'registerForm')
renderForm(app, LoginForm, 'loginForm')
renderForm(app, PostForm, 'postForm')
renderForm(app, CommentForm, 'commentForm')

