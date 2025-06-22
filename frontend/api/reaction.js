import { renderErrorPage } from "/frontend/pages/errorPage.js"
import { navigateTo } from "/frontend/utils.js"

async function addReaction(reactionData) {
    try {
        let response = await fetch("http://localhost:8080/api/react/like", {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(reactionData)
        })
        return [response.status, await response.json()]
    } catch (err) {
        console.error(err)
    }
}

export function ToggleLike(reactionData, svg, count) {
    addReaction(reactionData).then(
        ([status, response]) => {
            if (status == 401) {
                navigateTo("/login")
            }
            if ([400, 429, 500].includes(status)) {
                renderErrorPage(status)
            }
            if (status == 200 && response) {
                let reaction = parseInt(response.reaction)
                switch (reaction) {
                    case 1:
                        count.textContent = + count.textContent + 1
                        svg.style.fill = "red"
                        break;
                    case 0:
                        count.textContent = + count.textContent - 1
                        svg.style.fill = "white"
                        break;
                }
            }
        }
    )
}
