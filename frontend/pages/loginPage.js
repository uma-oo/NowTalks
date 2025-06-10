import { createForm } from "../components/form.js"
import { creatLink } from "../components/links.js"
import { loginForm } from "../const/forms.js"
import { navigateTo, createElement } from "../utils.js"
import { isLoggedIn } from "../api/user.js"

export function renderLoginPage(app) {
    isLoggedIn().then(data => {
        if (!data.is_logged_in) {
            app.innerHTML = ""
            let header = createElement("div", null)
            let formTititls = createElement("h1", null, "Sign In")
            let formSubTitle = createElement("h2", null, "Welcome Back — Continue your conversations")
            let formError = createElement("div", "form-error") 
            let loginFormElement = createForm(loginForm, "login-form")
            let goToRegister = createElement('p', null, "Don't have an account?")
            let registerLink = creatLink("Register", "./register", "")
            goToRegister.append(registerLink)

            header.append(formTititls, formSubTitle)

            app.append(header, formError, loginFormElement, goToRegister)
        }
        else {
            app.dataset.nickname = data.nickname
            app.dataset.id = data.id
            navigateTo("/")
        }
    })
}
