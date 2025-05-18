import { createIcon } from "./icon.js"

export function createButton(content, type, style) {

    let button = document.createElement('button')

    button.setAttribute('type', type)
    button.classList.add(...style)
    let icon = createIcon(content.icon)
    let contentSpan = document.createElement('span')
    contentSpan.textContent = content.text
    button.append(contentSpan)
    return button
}