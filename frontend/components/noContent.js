export function createNoContent(text){
    let noContent = document.createElement('p')
    noContent.classList.add("noContent")
    noContent.textContent = text
}