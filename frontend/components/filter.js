import { getCategories } from "../api/posts.js"

export function createFilterContainer() {
    let filterContainer = document.createElement('div')
    filterContainer.classList.add("filter-container","toggleable")

    
    
    getCategories.then(response=>console.log(response))

    

    return filterContainer
}
