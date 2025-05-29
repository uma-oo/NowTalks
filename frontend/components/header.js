import { logoutUser } from "../api/user.js"
import { navigateTo } from "../utils.js"
import { createButton } from "./button.js"

export function createHeader() {
    let header = document.createElement('header')
    let logo = document.createElement('a')
    logo.href = '/'
    logo.textContent = 'Talkaa'
    logo.className = 'logo'
    
    let logOut = createButton({text: "log out", icon: "logout"},'button','logout-btn')
    logOut.addEventListener("click", (e) => {
        logoutUser().then(data=>navigateTo("/login"));
        
    })

    header.append(logo, logOut)
    return header
}