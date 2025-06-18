
import { fetchUsers } from "./components/chatSection.js";
import { renderApp } from "./index.js";

export async function navigateTo(pathname) {
    history.replaceState({}, "", pathname)
    renderApp()
}

export function timeAgo(timestamp, locale = 'en') {
    let value;
    const diff = Math.floor((new Date().getTime() - new Date(timestamp).getTime()) / 1000);
    const minutes = Math.floor(diff / 60);
    const hours = Math.floor(minutes / 60);
    const days = Math.floor(hours / 24);
    const months = Math.floor(days / 30);
    const years = Math.floor(months / 12);

    const rtf = new Intl.RelativeTimeFormat(locale, { numeric: "auto" });
    if (years > 0) {
        value = rtf.format(-  years, "year");
    } else if (months > 0) {
        value = rtf.format(-  months, "month");
    } else if (days > 0) {
        value = rtf.format(-  days, "day");
    } else if (hours > 0) {
        value = rtf.format(-  hours, "hour");
    } else if (minutes > 0) {
        value = rtf.format(-  minutes, "minute");
    } else {

        value = rtf.format(-  diff, "second");
    }
    return value;
}

export function formatTimestamp(date) {
    const now = new Date();
    const d = new Date(date);
    const diffTime = now - d;
    const diffDays = Math.floor(diffTime / (1000 * 60 * 60 * 24));

    const options = { hour: 'numeric', minute: 'numeric', hour12: true };

    if (diffDays === 0 && d.getDate() === now.getDate()) {
        // Today
        return d.toLocaleTimeString([], options);
    } else if (diffDays === 1 || (
        now.getDate() - d.getDate() === 1 &&
        now.getMonth() === d.getMonth() &&
        now.getFullYear() === d.getFullYear()
    )) {
        return "Yesterday";
    } else if (diffDays < 7 && d.getDay() !== now.getDay()) {
        // Within the same week
        return d.toLocaleDateString(undefined, { weekday: 'long' });
    } else {
        return d.toLocaleDateString();
    }
}

export function throttledScrollFetcher(func) {
    return throttle((e) => {
        const container = e.target
        const scrollTop = container.scrollTop
        const scrollHeight = container.scrollHeight
        const clientHeight = container.clientHeight

        if (scrollTop + clientHeight >= scrollHeight * 0.8) {
            func(container)
        }
    }, 300)
}

export function throttle(func, delay) {
    let delayPassed = true
    return function (...arg) {
        if (delayPassed) {
            func(...arg);
            delayPassed = false
            setTimeout(() => {
                delayPassed = true
            }, delay)
        }
    }
}

export function createElement(tag, className, text = '') {
    let element = document.createElement(tag)
    if (className) element.className = className
    if (text) element.textContent = text
    return element
}

export function setAttributes(elem, attributes) {
    for (let [key, val] of Object.entries(attributes)) {
        elem.setAttribute(key, val)
    }
}

export function setOpions(selectElement, options) {
    options.forEach(option => {
        let optionElement = document.createElement('option')
        optionElement.setAttribute('value', option)
        optionElement.textContent = option
        selectElement.append(optionElement)
    })
}

export function loadFormErrors(form, data) {
    for (let [field, error] of Object.entries(data)) {
        let inputError = form.querySelector(`.form-grp[data-for="${field}"]>.input-error`)
        if (inputError) {
            inputError.textContent = error;
        }
    }
}

export function ReorderUsers(data) {
    let chatList = document.querySelector('.chat-list')
    let user1 = document.querySelector(`.chat-user-card[data-id="${data.receiver_id}"]`)
    let user2 = document.querySelector(`.chat-user-card[data-id="${data.sender_id}"`)
    let userCard = user1 || user2
    chatList.prepend(editUserCard(userCard, data))

    fetchUsers(chatList).then(()=> editUserCard(userCard))
}

function editUserCard(userCard) {
    if (userCard.dataset.open !== "true") {
        console.log("updating user card notifications")
        console.log(userCard.dataset.notifications)
        userCard.querySelector('.notification_container').classList.remove('hide')
        userCard.querySelector('.user_notifications').textContent = + userCard.dataset.notifications + 1
        userCard.dataset.notifications = + userCard.dataset.notifications + 1
    }
    return userCard
}