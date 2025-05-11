import { createChatSections } from "../components/chatSection.js";
import { createFooter } from "../components/footer.js";
import { createHeader } from "../components/header.js";
import { createPostsSections } from "../components/postsSection.js";
import { createTabsContainer } from "../components/tabsContainer.js";



export async function renderHomePage() {
    let header = createHeader()
    // let footer = createFooter()
    let main =  document.createElement('main')
    main.classList.add("home-main")
    let mainSectionsContainer = document.createElement('div')
    mainSectionsContainer.classList.add("main-sectionsContainer")
    mainSectionsContainer.append(await createPostsSections(), createChatSections())
    main.append(createTabsContainer(),mainSectionsContainer)
    app.append(header,main)
}



