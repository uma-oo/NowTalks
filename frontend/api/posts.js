
export async function addPostApi(postData) {
    const response = await fetch('/api/post', {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(postData)
    })
    return response.json()
}

export async function getPostsApi() {
    try {
        const response = await fetch('/api/post', {
            method: 'GET',
            headers: { 'Content-Type': 'application/json'},
        })
        return await response.json()
    } catch (error) {
        console.error(error);
    }
}





