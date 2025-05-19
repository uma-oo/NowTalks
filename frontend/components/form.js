import { createButton } from "./button.js";
import { navigateTo, setAttributes, setOpions } from "../../utils.js";
import { createUser, loginUser } from "../api/user.js";
import { addPostApi } from "../api/posts.js";
import { app } from "../index.js"

export function createForm(formRepresentaion, id) {
    let formElement = document.createElement('form')
    formElement.noValidate = true
    formElement.id = id


    formRepresentaion.elements.forEach((elem) => {
        let formGrp = document.createElement('div')
        formGrp.className = 'form-grp'
        let label = document.createElement('label')
        label.textContent = elem.label
        label.setAttribute('for', elem.attributes.id)
        let formInput = document.createElement(elem.tag)
        setAttributes(formInput, elem.attributes)
        if (elem.tag == 'select') {
            setOpions(formInput, elem.options)
        }
        formGrp.style.width = elem.style.width

        let inputError = document.createElement('p')
        inputError.classList.add("input-error")

        formGrp.append(label, formInput,inputError)

        formElement.append(formGrp)
    });

    let formButtons = document.createElement('div')
    formButtons.classList.add("form-buttons")

    formRepresentaion.buttons.forEach(button => {
        formButtons.append(createButton(button.content, button.type, [button.style]))
    })

    formElement.append(formButtons)

    formElement.addEventListener('submit', (e) => { handleFormSubmit(e) })
    return formElement
}

export function handleFormSubmit(event) {
    event.preventDefault()
    let form = new FormData(event.target)
    const formData = Object.fromEntries(form.entries())
    switch (event.target.id) {
        case "login-form":
            login(event.target, formData)
            break;
        case "register-form":
            register(event.target, formData)
            break;
        case "create-post-form":
            createPost(event.target, formData)
            break;
        default:
            break;
    }
}

export function login(form, data) {
    loginUser(data).then(([status, data]) => {
        let formError = form.parentElement.querySelector(".form-error")
        if (status == 200 ) navigateTo("/")
        else if (status == 400) {
            formError.innerText = ""
            formError.classList.remove("form-have-error")
            loadFormErrors(form, data.errors)
        } else if (status == 401) {
            let errors = form.querySelectorAll(".input-error")
            errors.forEach(error => error.textContent = "")
            formError.innerText = "ERROR!! Username or Email does not exist! Or Password Incorrect!"
            formError.classList.add("form-have-error")
        }   
    }).catch(error => console.error("Error submitting login form", error))
}

export function register(form, data) {
    data.age = parseInt(data.age)
    createUser(data)
        .then(([status,data]) => {
            if (status === 200 ) {
                navigateTo("/")
            }
            else if (status === 400 ) {
                loadFormErrors(form, data.errors)
            }
        })
        .catch(error => console.log("error submitting register form: ", error))
}

export function createPost(form, data) {
    
    data.user_id = app.dataset.id
    addPostApi(data)
        .then(([status,data]) => {
            if (status === 200 ) {
                navigateTo("/")
            }
            else if (status === 400 ) {
                loadFormErrors(form, data.errors)
            }
        })
        .catch(error => console.log("error submitting register form: ", error))
}



function loadFormErrors(form, data) {
    for (let [fieldId, error] of Object.entries(data)) {
        let inputError = form.querySelector(`#${fieldId}`).nextSibling
        inputError.textContent = error;
    }
}