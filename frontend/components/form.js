import { addPostApi } from "../api/posts.js";
import { createButton } from "./button.js";
import { createCheckboxInput } from "./checkbox.js";
import { createUser, loginUser } from "../api/user.js";
import { createElement, loadFormErrors, navigateTo, setAttributes, setOpions } from "../utils.js";

export function createForm(formRepresentaion, id) {
    let formElement = document.createElement('form')
    formElement.noValidate = true
    formElement.id = id

    formRepresentaion.elements.forEach((elem) => {
        let formGrp = createElement('div', 'form-grp')
        formGrp.dataset.for = elem.attributes.name
        let label = createElement('label', null, elem.label)
        label.setAttribute('for', elem.attributes.id)
        let formInput = createElement(elem.tag, null)
        setAttributes(formInput, elem.attributes)
        if (elem.tag == 'select') {
            setOpions(formInput, elem.options)
        }
        formGrp.style.width = elem.style.width
        let inputError = createElement('p', 'input-error')
        formGrp.append(label, formInput, inputError)
        formElement.append(formGrp)
    });

    let formButtons = createElement("div", 'form-buttons')

    formRepresentaion.buttons.forEach(button => {
        formButtons.append(createButton(button.content, button.type, button.style))
    })

    if (id == 'create-post-form') {
        let categoriesFormGrp = createElement('div', 'form-grp')
        categoriesFormGrp.dataset.for = "categories"
        let categoriesLabel = createElement('label', null, 'Post Categories')
        categoriesLabel.setAttribute("for", "categories")
        let app = document.querySelector('#app')
        let categories = app.dataset.categories.split(',')
        let categoriesList = createElement('div', 'categories-list')
        categories.forEach(category => {
            if (!category) return
            let [id, name] = category.split('-')

            let optionElem = createCheckboxInput(`category${id}`, id, name)
            categoriesList.append(optionElem)
        })
        let inputError = createElement('p', 'input-error')
        categoriesFormGrp.append(categoriesLabel, categoriesList, inputError)
        formElement.append(categoriesFormGrp)
    }

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
            formData.categories = form.getAll('categories').map(cat => parseInt(cat))
            createPost(event.target, formData)
            break;
        default:
            break;
    }
}

export function login(form, data) {
    loginUser(data).then(([status, data]) => {
        let formError = form.parentElement.querySelector(".form-error")
        if (status == 200) {
            console.log(data)
            navigateTo("/")
        } else if (status == 400) {
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
        .then(([status, data]) => {
            if (status === 200) {
                navigateTo("/")
            }
            else if (status === 400) {
                loadFormErrors(form, data.errors)
            }
        })
        .catch(error => console.log("error submitting register form: ", error))
}

export function createPost(form, data) {
    addPostApi(data)
        .then(([status, data]) => {
            if (status === 200) {
                form.reset()
                navigateTo("/")
            }
            else if (status === 400) {
                loadFormErrors(form, data.errors)
            }
        })
        .catch(error => console.log("error submitting register form: ", error))
}

