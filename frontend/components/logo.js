import { createElement,navigateTo } from "/frontend/utils.js"

export function createLogo () {
    let appLogo = createElement('a','logo')
    let appImg = createElement('img')
    let appName = createElement('span',null, "ⵙⴰⵡⵍ")
    appImg.src = "/frontend/assets/logo.png"
    appLogo.addEventListener("click",(e)=> {
        e.preventDefault()
        navigateTo("/")
    })
    appLogo.append(appImg,appName)
    return appLogo
}