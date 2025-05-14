import { renderForm } from "../components/form.js"
import { loginForm } from "../const/forms.js"

export function renderLoginPage(app) {
    app.innerHTML = ""
    let logingForm = renderForm(loginForm,"login-form")
    app.append(logingForm)
}
