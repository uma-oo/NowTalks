import { getCategories } from "../api/posts.js"
import {  categories } from "../const/data.js"
import { createCheckboxInput } from "./checkbox.js"
import { createElement } from "../utils.js"


const interactions = [
    "likedPosts",
    "createdPosts"
]

// const filter




export function createFilterContainer() {
    let filterContainer = createElement('div', 'filter-container toggleable')

    let filterTitle = createElement('h2', null, "Filter:") 
    let subTitlte1 =createElement('h4', null, "Filter by interaction : ")

    let filterOptions1 = document.createElement('div')
    interactions.forEach(option => {
        let optionElem = createCheckboxInput(`filter-${option}`, option)
        filterOptions1.append(optionElem)
    })

    let subTitlte2 = createElement('h4', null, "Filter by categories :")



    let app = document.querySelector('#app')
    let categories = app.dataset.categories.split(',')
    let filterOptions2 = createElement('div', null)
    categories.forEach(category => {
        if (!category) return
        let [id,name] = category.split('-')
        let optionElem = createCheckboxInput(`filter-category${id}`, name)
        filterOptions2.append(optionElem)
    })
    
    filterContainer.append(filterTitle,subTitlte1,filterOptions1,subTitlte2,filterOptions2)
    return filterContainer
}
