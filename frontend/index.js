import { renderErrorPage } from "/frontend/pages/errorPage.js";
import { renderHomePage } from "/frontend/pages/homePage.js"
import { renderLoginPage } from "/frontend/pages/loginPage.js";
import { renderRegisterPage } from "/frontend/pages/registerPage.js";

export let app  = document.querySelector('#app')

export function renderApp() {
    console.log("inside the renderApp function:");
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
            renderErrorPage(404)
            break;
    }
}

renderApp()

