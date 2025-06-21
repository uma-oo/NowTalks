import { createForm } from "/frontend/components/form.js"
import { creatLink } from "/frontend/components/links.js"
import { loginForm } from "/frontend/const/forms.js"
import { navigateTo, createElement } from "/frontend/utils.js"
import { isLoggedIn } from "/frontend/api/user.js"
import { createLogo } from "/frontend/components/logo.js"


export function renderLoginPage(app) {
    isLoggedIn().then(data => {
        if (!data.is_logged_in) {
            app.innerHTML = ""
            let header = createElement("div", "form-header")
            let logo = createLogo()
            let formTititls = createElement("h3", null, "Sign In")
            let formSubTitle = createElement("p", null, "Welcome Back — Continue your conversations")
            let formError = createElement("div", "form-error")
            let loginFormElement = createForm(loginForm, "login-form")
            let goToRegister = createElement('p', null, "Don't have an account? ")
            let registerLink = creatLink("Register", "./register", "")
            goToRegister.append(registerLink)

            header.append(formTititls, formSubTitle)

            app.append(logo,header, formError, loginFormElement, goToRegister)
        }
        else {
            app.dataset.nickname = data.nickname
            app.dataset.id = data.id
            navigateTo("/")
        }
    })
}
