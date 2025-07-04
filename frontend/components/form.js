import { addPostApi } from "/frontend/api/posts.js";
import { createButton } from "/frontend/components/button.js";
import { createUser, isLoggedIn, loginUser } from "/frontend/api/user.js";
import { createElement, loadFormErrors, navigateTo, setAttributes, setOpions } from "/frontend/utils.js";
import { addComment } from "/frontend/api/comment.js";
import { createComment } from "/frontend/components/comment.js";
import { sendMessage } from "/frontend/websocket.js";
import { createCheckboxInput } from "/frontend/components/checkbox.js";
import {createPostCard} from "/frontend/components/postCard.js"
import { renderErrorPage } from "/frontend/pages/errorPage.js";
import { toggleCreatePostFormContainer } from "/frontend/components/postsSection.js"

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
            formData.age = parseInt(formData.age)
            register(event.target, formData)
            break;
        case "create-post-form":
            formData.categories = form.getAll('categories').map(cat => parseInt(cat))
            createPost(event.target, formData)
            break;
        case "comment-form":
            handleCreateComment(event.target, formData)
            break;
        case "message-form":
            isLoggedIn().then(data => {
                if (data.is_logged_in) {
                    sendMessage(formData.chatMessage)
                    event.target.reset()
                }
                else navigateTo("/login")
            })
            break;
    }
}

export function login(form, data) {
    loginUser(data)
        .then(([status, data]) => {
            let formError = form.parentElement.querySelector(".form-error")
            if (status == 200) {
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
            } else if ([429, 500].includes(status)) {
                renderErrorPage(status)
            }
        })
}

export function register(form, data) {
    createUser(data)
        .then(([status, data]) => {
            if (status === 200) {
                navigateTo("/")
            }
            else if (status === 400) {
                loadFormErrors(form, data.errors)
            }
            else if ([429, 500].includes(status)) {
                renderErrorPage(status)
            }
        })
}

export function createPost(form, data) {
    addPostApi(data)
        .then(([status, data]) => {
            if (status === 200) {
                let post = createPostCard(data)
                let postsContainer = document.querySelector(".posts_container")
                toggleCreatePostFormContainer()
                postsContainer.prepend(post)
                form.reset()
            }
            else if (status === 400) {
                loadFormErrors(form, data.errors)
            } else if (status === 401) {
                navigateTo("/login")
            } else if ([429, 500].includes(status)) {
                renderErrorPage(status)
            }
        })
}

function handleCreateComment(form, data) {
    data.post_id = parseInt(form.dataset.postId)
    addComment(data)
        .then(([status, data]) => {
            if (status === 200) {

                data.user_name = sessionStorage.getItem("userNickname")
                let commentsContainer = form.parentElement.querySelector(".comments-container")
                let commentsCount = commentsContainer.parentElement.parentElement.querySelector('.reaction-container[data-reaction="comment"] > span')
                commentsCount.textContent = +commentsCount.textContent + 1
                commentsContainer.prepend(createComment(data))
                let noComent = commentsContainer.querySelector(".no-content")
                if (noComent) noComent.remove()
                form.querySelector('.input-error').textContent = ""
                form.reset()
            } else if (status == 400) {
                loadFormErrors(form, data.errors)
            } else if ([429, 500].includes(status)) {
                renderErrorPage(status)
            }
        })
}
