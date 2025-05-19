import { getCategories } from "../api/posts.js"

export function createFilterContainer() {
    let filterContainer = document.createElement('div')
    filterContainer.classList.add("filter-container","toggleable")

    let pickElement = document.createElement('div')
    
    
    getCategories().then(response=>console.log(response))

    

    return filterContainer
}
