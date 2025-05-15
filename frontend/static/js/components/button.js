export function createButton(content, type, style) {
    let button = document.createElement('button')
    button.setAttribute('type', type)
    button.textContent = content
    button.classList.add(...style)
    return button
}