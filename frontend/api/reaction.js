import { navigateTo } from "../utils.js"

async function addReaction(reactionData) {
    console.log(reactionData);
    try {
        let response = await fetch("/api/react/like", {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(reactionData)
        })
        return [response.status, await response.json()]
    } catch (err) {
        console.error(err)
    }
}




export function ToggleLike(reactionData, likedElement) {

    addReaction(reactionData).then(
        ([status, response]) => {
            console.log(response);
            if (status == 401) {
                navigateTo("login")
            }
            if (status == 200 && response) {
                let reaction = parseInt(response.reaction)
                switch (reaction) {
                    case 1:
                        likedElement.style.fill = "red"
                        break;
                    default:
                        likedElement.style.fill = "white"
                }

            }
        }
    )


}


