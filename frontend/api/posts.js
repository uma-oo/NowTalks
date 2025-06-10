
export async function addPostApi(postData) {
    try {
        const response = await fetch('/api/post', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body : JSON.stringify(postData)
        })
        return [response.status, await response.json()]
    } catch (error) {
        console.error(error)
    }
}

export async function getPostsApi(offset) {

    try {
        const response = await fetch(`/api/post?offset=${offset}`, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        })
        return await response.json()
    } catch (error) {
        console.error("trying to fetch posts",error);
    }
}


