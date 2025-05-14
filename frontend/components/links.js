export function creatLink(content,url,target="_blank") {
    let a = document.createElement('a')
    a.textContent = content
    a.href = url
    a.target = target
    return a
}