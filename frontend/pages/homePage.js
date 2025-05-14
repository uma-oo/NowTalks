import { createChatSections } from "../components/chatSection.js";
import { createFooter } from "../components/footer.js";
import { createHeader } from "../components/header.js";
import { createPostsSections } from "../components/postsSection.js";
import { createTabsContainer } from "../components/tabsContainer.js";



export function renderHomePage(app) {
    let header = createHeader()

    let main =  document.createElement('main')
    main.classList.add("home-main")
    
    main.append(createPostsSections(), createChatSections())
    app.append(header,main)
}



