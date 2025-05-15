export function createButton(content, type, style) {
    let button = document.createElement('button')
    button.setAttribute('type', type)
    button.classList.add(...style)
    let contentSpan = document.createElement('span')
    contentSpan.textContent = content
    button.append(contentSpan)
    return button
}