import { createButton } from "./button.js";
import { navigateTo, setAttributes, setOpions } from "../../utils.js";
import { createUser, loginUser } from "../api/user.js";

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

    formElement.addEventListener('submit',(e)=>{handleForm(e)})
    return formElement
}


export function handleForm(event) {
     event.preventDefault()
        let form = new FormData(event.target)
        const formData = Object.fromEntries(form.entries())
        switch(event.target.id) {
            case "login-form":
                login(event.target, formData)
                break;
            case "register-form":
                register(event.target, formData)
            default:
                break;
        }
}


export function login(form, data) {
    loginUser(data).then(response => {
        if (response.status == 200) navigateTo("/")
        else if (response.status == 401) console.log("wrong creadentials", response)
    }).catch(error => console.log("Error submitting login form", error))
} 

export function register(form,data){
    data.age = parseInt(data.age)
    createUser(data)
    .then(response => {
        if (response.ok || response.status === 403){
            navigateTo("/")
        }
        else if (response.status === 400) {
            console.log("bad request", response)
        }
    })
    .catch(error => console.log("error submitting register form: ",error))
}

