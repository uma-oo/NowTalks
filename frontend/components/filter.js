import { getCategories } from "../api/posts.js"

export function createFilterContainer() {
    let filterContainer = document.createElement('div')
    filterContainer.classList.add("filter-container","toggleable")

    getCategories().then(response=>console.log(response))

    let filterTitle = document.createElement('h3')
    filterTitle.textContent = "Filter:"

    let filterOption1Title = document.createElement(;)
    let filterOption1 = document.createElement('div')
    option1.innerText = "Filter by :"

    let filterOption2 = doc
    let filter

    // FileReader.
    return filterContainer
}
