import { createElement } from "../utils.js"

export function createCheckboxInput(label, categoryId, categoryName) {
    let labelElem = createElement('label', "category-option", categoryName)
    labelElem.setAttribute("for", label)

    let checkBoxInput = document.createElement('input')
    checkBoxInput.id = label
    checkBoxInput.type = "checkbox"
    checkBoxInput.value  = categoryId
    checkBoxInput.name = "categories"

    labelElem.append(checkBoxInput)
    return labelElem
}