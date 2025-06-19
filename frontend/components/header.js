import { createElement } from "../utils.js"
import { isLoggedIn, logoutUser } from "../api/user.js"
import { navigateTo } from "../utils.js"
import { createButton } from "./button.js"
import { createLogo } from "./logo.js"
import { closeConnection } from "../websocket.js"

export function createHeader() {
    let header = createElement('header')
    let logOut = createButton({ text: "log out", icon: "logout" }, 'button', 'logout-btn')
    logOut.addEventListener("click", (e) => {
        logoutUser().then(status => {
            console.log(status)
            if (status == 200 || status == 401) {
                navigateTo("/login")
                closeConnection()
            }
        });
    })
    header.append(createLogo())
    isLoggedIn().then(data => {
        if (data.is_logged_in) header.append(logOut)
    })
    return header
}