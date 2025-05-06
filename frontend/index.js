// import {getPosts} from "./api/posts.js"
import { registerFom } from "./const/forms.js"
import { renderForm } from "./layout.js"


// const registerForm = document.getElementById("registerForm")

// registerForm.addEventListener('submit', e => {
//     e.preventDefault()

//     let form = new FormData(registerForm)

//     const data = Object.fromEntries(form.entries());
// })






  //   fetch('https://fakestoreapi.com/products')
  // .then(response => response.json())
  // .then(data => console.log(data))



renderForm(document.getElementsByTagName('body'), registerFom, 'registerForm')
