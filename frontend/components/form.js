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
        formInput.addEventListener('blur', () => {
            formInput.classList.add('input-filled');
        });
    });

    formRepresentaion.buttons.forEach(button => {
        formElement.append(createButton(button.content, button.type, [button.style]))
    })

    formElement.addEventListener('submit', (e) => { handleForm(e) })
    return formElement
}

export function handleForm(event) {
    event.preventDefault()
    let form = new FormData(event.target)
    const formData = Object.fromEntries(form.entries())
    console.log(formData)
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
    loginUser(data).then(response => {
        console.log(response
        )
        let formError = form.parentElement.querySelector(".form-error")
        if (response.status == 200) navigateTo("/")
        else if (response.status == 400) {
            formError.innerText = ""
            loadFormErrors(form, response.errors)
        } else if (response.status == 401) {
            let errors = form.querySelectorAll(".input-error")
            errors.forEach(error => error.textContent = "")
            formError.innerText = "ERROR!! Username or Email does not exist! Or Password Incorrect!"
        }   
    }).catch(error => console.error("Error submitting login form", error))
}

export function register(form, data) {
    data.age = parseInt(data.age)
    createUser(data)
        .then(response => {
            if (response.ok ) {
                navigateTo("/")
            }
            else if (response.status === 400 ) {
                console.log("bad request", response)
                loadFormErrors(form, response.errors)
            }
        })
        .catch(error => console.log("error submitting register form: ", error))
}

export function createPost(form, data) {
    data.user_id = app.dataset.id
    addPostApi(data)
        .then(response => {
            console.log(response)
        })
        .catch(error => console.log("error creating new post"))
}



function loadFormErrors(form, data) {
    console.log(data)
    for (let [fieldId, error] of Object.entries(data)) {
        let inputError = form.querySelector(`#${fieldId}`).nextSibling
        inputError.textContent = error;
    }

}