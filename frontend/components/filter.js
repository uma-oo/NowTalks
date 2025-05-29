import { getCategories } from "../api/posts.js"
import {  categories } from "../const/data.js"
import { createCheckboxInput } from "./checkbox.js"


const interactions = [
    "likedPosts",
    "createdPosts"
]

// const filter




export function createFilterContainer() {
    let filterContainer = document.createElement('div')
    filterContainer.classList.add("filter-container","toggleable")


    let filterTitle = document.createElement('h2')
    filterTitle.textContent = "Filter:"

    let subTitlte1 = document.createElement('h4')
    subTitlte1.innerText = "Filter by interaction : "

    let filterOptions1 = document.createElement('div')
    interactions.forEach(option => {
        let optionElem = createCheckboxInput(`filter-${option}`, option)
        filterOptions1.append(optionElem)
    })

    let subTitlte2 = document.createElement('h4')
    subTitlte2.textContent =  "Filter by categories :"




    let filterOptions2 = document.createElement('div')
    categories.forEach(category => {
        let optionElem = createCheckboxInput(`filter-category${category.category_id}`, category.category_name)
        filterOptions2.append(optionElem)
    })
    
    filterContainer.append(filterTitle,subTitlte1,filterOptions1,subTitlte2,filterOptions2)
    return filterContainer
}
