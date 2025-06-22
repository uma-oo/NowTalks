import { createElement } from "/frontend/utils.js"
import { isLoggedIn, logoutUser } from "/frontend/api/user.js"
import { navigateTo } from "/frontend/utils.js"
import { createButton } from "/frontend/components/button.js"
import { createLogo } from "/frontend/components/logo.js"
import { closeConnection } from "/frontend/websocket.js"

export function createHeader() {
    let header = createElement('header')
    let logOut = createButton({ text: "log out", icon: "logout" }, 'button', 'logout-btn')
    logOut.addEventListener("click", async (e) => {
     
        let status = await logoutUser()

        if (status == 204 || status == 401) {
         
                navigateTo("/login")
                closeConnection()
            }
    })
    header.append(createLogo())
    isLoggedIn().then(data => {
        if (data.is_logged_in) header.append(logOut)
    })
    return header
}