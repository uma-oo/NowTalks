import { creatLink } from "/frontend/components/links.js"

export function createFooter() {
    let footer = document.createElement('footer')
    let ul = document.createElement('ul')
    ul.textContent =  "Created By :"  
    let links  = [
        {
            content : "ayoub",
            url : "https://github.com/DarkMethoss"        
        },
        {
            content: "oumayma",
            url : "https://github.com/uma-oo"
        }
    ]

    links.forEach(link=> {
        let li = document.createElement('li')
        li.append(creatLink(link.content,link.url))
        ul.append(li)
    })
    footer.append(ul)
    return footer
}