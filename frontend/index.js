
import { renderHomePage } from "./pages/home.js"

export const appContainer = document.getElementById('app')

// window.addEventListener('popstate', function (event) {
//     event.preventDefault()
// 	let page = this.window.location.href
//     console.log(page)
// });

// async function startApp() {
//     let posts = await getPosts()
//     console.log(posts)

//     if (posts.status === 401){
//         history.pushState({},"","/login")
//     }
// }


// startApp();




// renderForm(app, registerFom, 'registerForm')
// renderForm(app, LoginForm, 'loginForm')
// renderForm(app, PostForm, 'postForm')
// renderForm(app, CommentForm, 'commentForm')





renderHomePage()


// posts.forEach(post => app.append(createPostCard(post))) 

