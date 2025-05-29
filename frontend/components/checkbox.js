export function createCheckboxInput(id, value) {
    let label = document.createElement('label')
    let labelText = document.createElement('span')
    label.setAttribute("for", id)
    label.classList.add("category-option")
    labelText.textContent = value

    let checkBoxInput = document.createElement('input')
    checkBoxInput.id = id
    checkBoxInput.type = "checkbox"
    checkBoxInput.value  = value

    label.append(checkBoxInput, labelText)
    return label
}