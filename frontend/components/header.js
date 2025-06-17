import { createElement } from "../utils.js"
import { logoutUser } from "../api/user.js"
import { navigateTo } from "../utils.js"
import { createButton } from "./button.js"
import { createLogo } from "./logo.js"
import { closeConnection } from "../websocket.js"

export function createHeader() {
    let header = createElement('header')
    let logOut = createButton({text: "log out", icon: "logout"},'button','logout-btn')
    logOut.addEventListener("click", (e) => {
        logoutUser().then(data=>{
            navigateTo("/login")
            closeConnection()
        });
    })
    header.append(createLogo(), logOut)
    return header
}