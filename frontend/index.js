import { isLoggedIn } from "./api/user.js";
import { renderErrorPage } from "./pages/errorPage.js";
import { renderHomePage } from "./pages/homePage.js"
import { renderLoginPage } from "./pages/loginPage.js";
import { renderRegisterPage } from "./pages/registerPage.js";

let app  = document.querySelector('#app')

console.log(document.cookie)

export function renderApp() {
    app.innerHTML = ""
    switch (window.location.pathname) {
        case "/register":
            renderRegisterPage(app)
            break;
        case "/login":
            renderLoginPage(app)
            break;
        case "/":
            renderHomePage(app)
            break;
        default:
            renderErrorPage(app)
            break;
    }
}

renderApp()

