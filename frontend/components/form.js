import { createButton } from "./button.js";

export function renderForm(target, formRepresentaion, id) {
    let formElement = document.createElement('form')
    formElement.id = id
    
    formRepresentaion.elements.forEach((elem) => {
        let formGrp = document.createElement('div')
        formGrp.className = 'form-grp'
        
        let label = document.createElement('label')
        label.textContent = elem.label
        label.setAttribute('for',elem.attributes.id)
        
        let formInput = document.createElement(elem.tag)
        setAttributes(formInput, elem.attributes)
        if (elem.tag == 'select') {
            setOpions(formInput,elem.options)
        }
        formGrp.style.width = elem.style.width
        formGrp.append(label)
        formGrp.append(formInput)
        formElement.append(formGrp)
        formInput.addEventListener('blur', () => {
            formInput.classList.add('input-filled');
        });
    });

    formRepresentaion.buttons.forEach(button => {
        formElement.append(createButton(button.content,button.type,button.style))    
    })
    target.append(formElement)
    
    formElement.addEventListener('submit',(event)=>{
        event.preventDefault()
        let form = new FormData(formElement)
        const formData = Object.fromEntries(form.entries())
        console.log(formData)
    })

}

function setAttributes(elem, attributes) {
    for ( let [key,val] of Object.entries(attributes)) {
        elem.setAttribute(key,val)
    }
}

function setOpions(selectElement, options) {
    options.forEach(option=> {
        let optionElement = document.createElement('option')
        optionElement.setAttribute('value', option)
        optionElement.textContent = option
        selectElement.append(optionElement)
    })

}