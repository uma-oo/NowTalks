import { createForm } from "../components/form.js"
import { creatLink } from "../components/links.js"
import { registerFom } from "../const/forms.js"
import { navigateTo, createElement } from "../utils.js"
import { isLoggedIn } from "../api/user.js"

export function renderRegisterPage(app) {
    isLoggedIn().then(data => {
        if (!data.is_logged_in) {
            app.dataset.nickname = data.nickname
            app.dataset.id = data.id
            app.innerHTML = ""
            let header = createElement("div", null)
            let formTititls = createElement("h1", null, "Sign Up")
            let formSubTitle = createElement("h2", null, "New to the forum? Create an account to join the conversation.")
            let formError = createElement("div","form-Error")
            formError.className = "form-Error"
            let registerFomlement = createForm(registerFom, "register-form")
            let goToLogin = document.createElement('p')
            goToLogin.textContent = "Already have an account ? "
            let LoginLink = creatLink("Log In", "/login", "")
            goToLogin.append(LoginLink)
            header.append(formTititls,formSubTitle)
            app.append(header, formError, registerFomlement, goToLogin)
        }
        else {
            navigateTo("/")
        }
    })
}
