import { createButton } from "./button.js"

const tabs = [
    {
        content: "Posts",
        className :'active-tab-btn tab-btn',
        section: 'posts_section'
    },
    {
        content: "Chat",
        className : 'tab-btn',
        section: 'chat_section'
    }
]

export function createTabsContainer() {
    let tabsContainer = document.createElement('div')
    tabsContainer.classList.add('tabs-container')
    tabs.forEach(tab => {
        let tabBtn = createButton(tab.content,'button', tab.className)
        tabBtn.dataset.section = tab.section
        tabsContainer.append(tabBtn)
        
        tabBtn.addEventListener("click",(e)=>{
            let btn = e.target
            let sectionTriggered = document.querySelector(`.${btn.dataset.section}`) 
            if (!sectionTriggered.classList.contains("visible_tab_section")) {
                let tab_btns = document.querySelectorAll('.tab-btn')
                tab_btns.forEach(btn => btn.classList.toggle('active-tab-btn'))
                let sections = document.querySelectorAll(".tab_section")
                sections.forEach(section=>section.classList.toggle('visible_tab_section'))
            }
        })
    })

    return tabsContainer
}