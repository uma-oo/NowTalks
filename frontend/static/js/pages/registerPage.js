import { renderForm } from "../components/form.js"
import { creatLink } from "../components/links.js"
import { registerFom } from "../const/forms.js"
import { navigateTo } from "../../utils.js"

export function renderRegisterPage(app) {
    console.log("asdfasdfasfdasdf")
    app.innerHTML = ""
    let header = document.createElement('h1')
    let formError = document.createElement('div')
    formError.className = "form-Error"
    header.textContent = "Register"
    let registerFomlement = renderForm(registerFom, "register-form")
    let goToLogin = document.createElement('p')
    goToLogin.textContent = "Already have an account ? "
    let LoginLink = creatLink("Log In", "", "")
    goToLogin.append(LoginLink)
    LoginLink.addEventListener("click", (e) => {
        e.preventDefault()
        navigateTo('./login')
    })
    app.append(header, formError, registerFomlement, goToLogin)
}
