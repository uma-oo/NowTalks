// import { getPosts } from "./api/posts"

import { createFooter } from "./components/footer.js"
import { createHeader } from "./components/header.js"

export function renderLayout() {
    let header = createHeader()
    let footer = createFooter()
    let main =  document.createElement('main')


    




    return [header, main,footer]
}


