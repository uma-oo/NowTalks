
export async function addPost(postData) {
    const response = await fetch('/api/posts', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(postData)
    })
    return response.json()
}

export async function getPosts() {
    try {
        const response = await fetch('/api/posts', {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        })
        return  response.json()
    } catch (error) {
        throw new Error (error)
    }
}
