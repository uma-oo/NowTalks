import { createForm } from "../components/form.js"
import { creatLink } from "../components/links.js"
import { registerFom } from "../const/forms.js"
import { navigateTo } from "../../utils.js"
import { isLoggedIn } from "../api/user.js"

export function renderRegisterPage(app) {
    isLoggedIn().then(data => {
        if (!data.is_logged_in) {
            app.innerHTML = ""
            let header = document.createElement('h1')
            let formError = document.createElement('div')
            formError.className = "form-Error"
            header.textContent = "Register"
            let registerFomlement = createForm(registerFom, "register-form")
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
        else navigateTo("/")
    })
}
