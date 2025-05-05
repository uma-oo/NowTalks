import {getPosts} from "./api/posts.js"


// const registerForm = document.getElementById("registerForm")

// registerForm.addEventListener('submit', e => {
//     e.preventDefault()

//     let form = new FormData(registerForm)

//     const data = Object.fromEntries(form.entries());
// })



// let posts = await getPosts()
// console.log(posts)


    fetch('https://fakestoreapi.com/products')
  .then(response => response.json())
  .then(data => console.log(data))



