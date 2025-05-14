import { createButton } from "./button.js";
import { setAttributes, setOpions } from "../utils.js";
import { loginUser } from "../api/user.js";


export function renderForm(formRepresentaion, id) {
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
        formGrp.append(label,formInput)
        formElement.append(formGrp)
        formInput.addEventListener('blur', () => {
            formInput.classList.add('input-filled');
        });
    });

    formRepresentaion.buttons.forEach(button => {
        formElement.append(createButton(button.content,button.type,[button.style]))    
    })

    formElement.addEventListener('submit',(event)=>{
        event.preventDefault()
        let form = new FormData(formElement)
        const formData = Object.fromEntries(form.entries())

        switch (event.target.id) {
            case "login-form":
                login(event.target, formData)
                break;
            default:
                break;
        }
    })
    return formElement
}


export function login(form, data) {
    console.log(form)
    loginUser(data).then(data => {
        if (data.status == 200) console.log("user Loged in successfully")
        else if (data.status == 401) console.log("wrong creadentials", data)
    }).catch(error=> console.log("ERROR !! ", error))


} 