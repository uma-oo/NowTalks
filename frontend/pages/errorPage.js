import { createButton } from "/frontend/components/button.js"
import { createHeader } from "/frontend/components/header.js"
import { app } from "/frontend/index.js"
import { createElement, navigateTo } from "/frontend/utils.js"

export function renderErrorPage(status){
    app.innerHTML = ""
    let header = createHeader()
    let container = createElement("div", "app-error-container" )
    let statusCode = createElement("h1", "status-code", status)
    let errorMessage = createElement("p", "error-msg", getErrorMessage(status))
    let goBackBtn = createButton({icon:"home", text: "Go back home."},"butotn", "primary-btn")
    goBackBtn.addEventListener("click",()=> {
        navigateTo("/")
    })
    container.append(statusCode, errorMessage)
    if (status != 429) container.append(goBackBtn)
    app.append(header,container)
};

function getErrorMessage(status) {
    switch (status) {
        case 403:
            return "Access Forbidden."
        case 404:
            return "Page Not Found."
        case 429:
            return "Too Many request, please try again later"
        case 500:
            return "Internal Server Error."
    }
}