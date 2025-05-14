import { renderErrorPage } from "./pages/errorPage.js";
import { renderHomePage } from "./pages/homePage.js"
import { renderLoginPage } from "./pages/loginPage.js";
import { renderRegisterPage } from "./pages/registerPage.js";

let app  = document.querySelector('#app')

export function renderApp() {
    console.log(window.location.pathname)
    switch (window.location.pathname) {
        case "/register":
            console.log('going to register');
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

window.addEventListener('popstate', () => {
    console.log("url has been changed ")
    renderApp();
});

renderApp()

