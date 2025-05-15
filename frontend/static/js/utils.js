import { renderApp } from "./index.js";

export function currentRoute() {
    let sep  = new RegExp("/")
    let currentRoute = window.location.href.split('/')
}

export async function navigateTo(pathname) {
    history.replaceState({},"",pathname)
    renderApp();
}


export function timeAgo(timestamp, locale = 'en') {
    let value;
    const diff = (new Date().getTime() - new Date(timestamp).getTime()) / 1000;
    const minutes = Math.floor(diff / 60);
    const hours = Math.floor(minutes / 60);
    const days = Math.floor(hours / 24);
    const months = Math.floor(days / 30);
    const years = Math.floor(months / 12);
    const rtf = new Intl.RelativeTimeFormat(locale, { numeric: "auto" });

    if (years > 0) {
        value = rtf.format(0 - years, "year");
    } else if (months > 0) {
        value = rtf.format(0 - months, "month");
    } else if (days > 0) {
        value = rtf.format(0 - days, "day");
    } else if (hours > 0) {
        value = rtf.format(0 - hours, "hour");
    } else if (minutes > 0) {
        value = rtf.format(0 - minutes, "minute");
    } else {
        value = rtf.format(0 - diff, "second");
    }
    return value;
}






export function setAttributes(elem, attributes) {
    for ( let [key,val] of Object.entries(attributes)) {
        elem.setAttribute(key,val)
    }
}

export function setOpions(selectElement, options) {
    options.forEach(option=> {
        let optionElement = document.createElement('option')
        optionElement.setAttribute('value', option)
        optionElement.textContent = option
        selectElement.append(optionElement)
    })
}