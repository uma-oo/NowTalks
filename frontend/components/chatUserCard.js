import { createIcon } from "/frontend/components/icon.js";
import { createElement, formatTimestamp, timeAgo } from "/frontend/utils.js";

export function createChatUserCard({
    id,
    nickname,
    lastInteraction,
    lastMessage,
    notifications
}) {
    let chatUserCard = createElement('div', 'chat-user-card')
    let onlineUsers = sessionStorage.getItem("onlineUsers")
    chatUserCard.dataset.id = id
    chatUserCard.dataset.open = + sessionStorage.getItem("openChat") ===  id  ? "true" : ""
    chatUserCard.dataset.status = onlineUsers.includes(id) ? "online" : "offline"
    chatUserCard.dataset.userName = nickname
    chatUserCard.dataset.notifications = notifications || 0

    let chatUserCardBody = createElement('div', 'chat-user-card-body')
    let chatUserCardFooter = createElement('div', 'chat-user-card-footer')

    let user_nickname = createElement('p', "user_name", nickname)
    let user_status = createElement('span', "user_status", chatUserCard.dataset.status)
    let last_message = createElement('p', "latest_message", lastMessage)
    let last_interaction = createElement('span', "latest_interaction", lastInteraction !== "0" ? formatTimestamp(lastInteraction) : "")
    let notifications_container = createElement('div', `${+notifications > 0 ? "notification_container " : "notification_container hide"}`)
    let notifications_count = createElement('span', `user_notifications`, notifications)
    let notificationIcon = createIcon("notification")
    notifications_container.append(notificationIcon, notifications_count)

    chatUserCardBody.append(user_nickname, user_status, last_message)
    chatUserCardFooter.append(last_interaction, notifications_container)
    chatUserCard.append(chatUserCardBody, chatUserCardFooter)
    return chatUserCard
}