import { renderErrorPage } from "./pages/errorPage.js";
import { renderHomePage } from "./pages/homePage.js"
import { renderLoginPage } from "./pages/loginPage.js";
import { renderRegisterPage } from "./pages/registerPage.js";

export let app  = document.querySelector('#app')

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
            renderErrorPage(app,'404')
            break;
    }
}

renderApp()

