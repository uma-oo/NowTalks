import { createElement } from "/frontend/utils.js"
import { createIcon } from "/frontend/components/icon.js"

export function createButton(content, type, className) {
    let button = document.createElement('button')
    button.setAttribute('type', type)

    if (className) button.className = className
    let btnIcon = content.icon ? createIcon(content.icon) : "" 
    let btnText = createElement('span',null, content.text)
    button.append(btnIcon, btnText)
    return button
}