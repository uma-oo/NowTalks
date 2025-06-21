import { renderErrorPage } from "/frontend/pages/errorPage.js";
import { renderHomePage } from "/frontend/pages/homePage.js"
import { renderLoginPage } from "/frontend/pages/loginPage.js";
import { renderRegisterPage } from "/frontend/pages/registerPage.js";

export let app = document.querySelector('#app')

export function renderApp() {
    app.innerHTML = ""
    console.log(window.location.pathname)
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
        case "/api":
        case "/assets":
        case "/components":
        case "/const":
        case "/pages":
        case "/styles":
        case "/api/":
        case "/assets/":
        case "/components/":
        case "/const/":
        case "/pages/":
        case "/styles/":
            renderErrorPage(403)
            break;
        default:
            renderErrorPage(404)
            break;
    }
}

renderApp()