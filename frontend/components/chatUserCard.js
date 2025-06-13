import { createIcon } from "./icon.js";
import { createElement, formatTimestamp, timeAgo } from "../utils.js";

export function createChatUserCard({
    id,
    nickname,
    lastInteraction,
    lastMessage,
    notifications
}) {
    let chatUserCard = createElement('div','chat-user-card') 
    chatUserCard.dataset.open = ""

    let chatUserCardBody = createElement('div', 'chat-user-card-body') 
    let chatUserCardFooter =  createElement('div', 'chat-user-card-footer') 



    let user_nickname = createElement('p',"user_name",nickname)
    let user_status = createElement('span',"user_status","offline")
    let last_message = createElement('p',"latest_message", lastMessage.length>8?lastMessage.slice(0,8)+"...":lastMessage)
    let last_interaction = createElement('span',"latest_interaction", formatTimestamp(lastInteraction))
    let notifications_container =  createElement('div', "notification_container")
    let notifications_count = createElement('span', null, notifications)
    let notificationIcon = createIcon("notification")
    notifications_container.append(notificationIcon,notifications_count)

    // use it for time for latest message or is typing message 
    
    chatUserCardBody.append(user_nickname, user_status, last_message)
    chatUserCardFooter.append(last_interaction, notifications_container)
    chatUserCard.append( chatUserCardBody,chatUserCardFooter)

    return chatUserCard
}