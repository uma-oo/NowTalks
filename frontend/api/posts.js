
export async function addPost(postData) {
    const response = await fetch('/api/posts', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(postData)
    })
    return response.json()
}

export async function getPosts() {
    const response = await fetch('/api/posts')
    return response.json()
}
