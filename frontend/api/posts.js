
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

export async function getPostsApi() {
    try {
        const response = await fetch('/api/post', {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        })
        return await response.json()
    } catch (error) {
        console.error(error);
    }
}





