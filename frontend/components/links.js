import { navigateTo } from "/frontend/utils.js"

export function creatLink(content,url,target="_blank") {
    let a = document.createElement('a')
    a.textContent = content
    a.href = url
    a.target = target

    a.addEventListener("click", (e)=>{
        e.preventDefault()
        navigateTo(url)
    })

    return a
}