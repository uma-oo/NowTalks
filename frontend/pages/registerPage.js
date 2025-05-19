import { createForm } from "../components/form.js"
import { creatLink } from "../components/links.js"
import { registerFom } from "../const/forms.js"
import { navigateTo } from "../../utils.js"
import { isLoggedIn } from "../api/user.js"

export function renderRegisterPage(app) {
    isLoggedIn().then(data => {
        if (!data.is_logged_in) {
            app.dataset.nickname = data.nickname
            app.dataset.id = data.id
            app.innerHTML = ""
            let header = document.createElement('h1')
            let formError = document.createElement('div')
            formError.className = "form-Error"
            header.innerHTML = "Create account </br> and become a new member in our forum"
            let registerFomlement = createForm(registerFom, "register-form")
            let goToLogin = document.createElement('p')
            goToLogin.textContent = "Already have an account ? "
            let LoginLink = creatLink("Log In", "", "")
            goToLogin.append(LoginLink)
            LoginLink.addEventListener("click", (e) => {
                navigateTo('./login')
            })
            app.append(header, formError, registerFomlement, goToLogin)
        }
        else {
            app.dataset.nickname = data.nickname
            app.dataset.id = data.id
            navigateTo("/")
        }
    })
}
