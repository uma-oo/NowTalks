
export async function addPostApi(postData) {
    try {
        const response = await fetch('/api/post', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(postData)
        })
        return await response.json()
    } catch (error) {
        console.error(error)
    }
}

export async function getPostsApi(offset=0) {
    try {
        const response = await fetch(`/api/post?offset=${offset}&limit=10`, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        })
        return await response.json()
    } catch (error) {
        console.error(error);
    }
}


export function throttle(func, delay=1000) {
    let delayPassed = true
    return function (...arg) {
        if (delayPassed) {
            console.log(func)
            func(...arg);
            delayPassed = false
            setTimeout(() => {
                delayPassed = true
            }, delay)
        }
    }
}

