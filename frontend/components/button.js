import { createIcon } from "./icon.js"

export function createButton(content, type, style) {

    let button = document.createElement('button')

    button.setAttribute('type', type)
    Array.isArray(style) ? button.classList.add(...style) : button.classList.add(style)
    let icon = content.icon ? createIcon(content.icon) : "" 
    let contentSpan = document.createElement('span') 
    contentSpan.textContent = content.text ? content.text : ""
    button.append(icon, contentSpan)
    return button
}