import { createForm } from "/frontend/components/form.js"
import { creatLink } from "/frontend/components/links.js"
import { registerFom } from "/frontend/const/forms.js"
import { navigateTo, createElement } from "/frontend/utils.js"
import { isLoggedIn } from "/frontend/api/user.js"
import { createLogo } from "/frontend/components/logo.js"

export function renderRegisterPage(app) {
    isLoggedIn().then(data => {
        if (!data.is_logged_in) {
            let logo = createLogo()
            let header = createElement("div", "form-header")
            let formTititls = createElement("h2", null, "Sign Up")
            let formSubTitle = createElement("p", null, "New to the forum? Create an account to join the conversation.")
            let formError = createElement("div","form-Error")
            formError.className = "form-Error"
            let registerFomlement = createForm(registerFom, "register-form")
            let goToLogin = document.createElement('p')
            goToLogin.textContent = "Already have an account ?  "
            let LoginLink = creatLink("Log In", "/login", "")
            goToLogin.append(LoginLink)
            header.append(formTititls,formSubTitle)
            app.append(logo,header, formError, registerFomlement, goToLogin)
        }
        else {
            navigateTo("/")
        }
    })
}
