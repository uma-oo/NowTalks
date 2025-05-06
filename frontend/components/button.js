export function createButton(content, type, style) {
    let button = document.createElement('button')
    button.setAttribute('type', type)
    button.classList.add(style)
    button.textContent = content
    return button
}