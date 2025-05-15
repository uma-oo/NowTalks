import { renderForm } from "../components/form.js"
import { creatLink } from "../components/links.js"
import { loginForm } from "../const/forms.js"
import { navigateTo } from "../../utils.js"

export function renderLoginPage(app) {
    app.innerHTML = ""
    let header = document.createElement('h1')
    let formError = document.createElement('div')
    formError.className = "form-Error"
    header.textContent = "Log In"
    let loginFormElement = renderForm(loginForm,"login-form")
    let goToRegister = document.createElement('p')
    goToRegister.textContent = "Don't have an account? "
    let registerLink = creatLink("Register","","")
    goToRegister.append(registerLink)

    registerLink.addEventListener("click", (e) => {
        e.preventDefault()
        navigateTo('./register')
    })
    app.append(header,formError,loginFormElement,goToRegister)
}
