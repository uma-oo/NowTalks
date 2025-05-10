import { createButton } from "./button"

const tabs = [
    {
        content: "posts",
        show: true,
        section: 'posts_section'
    },
    {
        content: "Chat",
        hidden: false,
        section: 'chat_section'
    }
]

export function createTabsContainer() {
    let tabsContainer = document.createElement('div')


    tabs.forEach(tab => {
        let tabBtn = createButton(tab.content,'button','tab-btn')
        tabsContainer.append()

        tabBtn.addEventListener("click",()=>{
            let section = document.querySelector("visible_tab_section")
            section.classList.remove("visible_tab_section")
            
        })

    })


}