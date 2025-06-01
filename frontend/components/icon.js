import { createElement } from "../utils.js"

export function createIcon(name) {
    let icon = createElement('img','icon')
    icon.src = `../assets/icons/${name}.svg`
    return icon
}