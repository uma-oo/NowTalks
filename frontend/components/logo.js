import { createElement } from "/frontend/utils.js"

export function createLogo () {
    let appLogo = createElement('a','logo')
    let appImg = createElement('img')
    let appName = createElement('span',null, "ⵙⴰⵡⵍ")
    appImg.src = "/frontend/assets/logo.png"
    appLogo.href = 'http://localhost:8080/'
    appLogo.append(appImg,appName)
    return appLogo
}