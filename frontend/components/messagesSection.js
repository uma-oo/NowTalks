import { getMessages } from "../api/messages.js";
import { navigateTo } from "../utils.js";
import { createChatMessageContainer } from "./chatMessageContainer.js";

export function fetchMessages(offset, receiver_id, type, chatWindow) {

    const prevHeight = chatWindow.scrollHeight

    getMessages(offset, receiver_id, type)
        .then(([status, data]) => {
            if (status === 401) {
                navigateTo("/login")
            }

            if (status === 400) {
                console.log(data);
            }

            if (status === 200) {
                console.log(data)

                if (!data || data.length < 10) {
                    chatWindow.dataset.topObsorver = "off"
                }
                // if (type=="old") data = data.reverse()
                data?.forEach(message => {
                    createChatMessageContainer(message, chatWindow, "top")
                });

                requestAnimationFrame(() => {
                    const diff = chatWindow.scrollHeight - prevHeight;
                    chatWindow.scrollTop += diff;
                });

                if (chatWindow.dataset.firstFetch === "true") {
                    console.log("first fetch")
                    setTimeout(() => {
                        const lastMsg = chatWindow.querySelector(".message-bubble:last-child");
                        lastMsg?.scrollIntoView({ behavior: "auto", block: "end" });
                    }, 0);
                    chatWindow.dataset.firstFetch = "false"
                }
            }
        })

}