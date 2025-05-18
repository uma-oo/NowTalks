import { iconsKit } from "../const/icons.js"

export function createIcon(icon) {
    let style = "fa-solid"
    let i = document.createElement('i')
    i.classList.add(style, iconsKit[icon])
}