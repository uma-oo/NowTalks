import { createForm } from "../components/form.js"
import { creatLink } from "../components/links.js"
import { loginForm } from "../const/forms.js"
import { navigateTo } from "../../utils.js"
import { isLoggedIn } from "../api/user.js"

export function renderLoginPage(app) {
    isLoggedIn().then(data => {
        if (!data.is_logged_in) {
            app.innerHTML = ""
            let header = document.createElement('h1')
            let formError = document.createElement('div')
            formError.className = "form-error"
            header.textContent = "Log In"
            let loginFormElement = createForm(loginForm, "login-form")
            let goToRegister = document.createElement('p')
            goToRegister.textContent = "Don't have an account? "
            let registerLink = creatLink("Register", "", "")
            goToRegister.append(registerLink)
            registerLink.addEventListener("click", (e) => {
                e.preventDefault()
                navigateTo('./register')
            })
            app.append(header, formError, loginFormElement, goToRegister)
        }
        else {
            app.dataset.nickname = data.nickname
            app.dataset.id = data.id
            navigateTo("/")
        }
    })

}
