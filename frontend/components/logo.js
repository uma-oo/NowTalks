import { createElement } from "../utils.js"

export function createLogo () {
    let appLogo = createElement('a','logo')
    let appImg = createElement('img')
    let appName = createElement('span',null, "Sawel")
    appImg.src = "../assets/logo.png"
    appLogo.href = './'
    appLogo.append(appImg,appName)
    return appLogo
}