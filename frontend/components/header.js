import { createButton } from "./button.js"

export function createHeader() {
    let header = document.createElement('header')
    let logo = document.createElement('a')
    logo.href = '/'
    logo.textContent = 'Talkaa'
    logo.className = 'logo'
    let logOut = createButton("Log-Out",'button','logout-btn')
    
    header.append(logo, logOut)
    return header
}