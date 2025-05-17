
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
    let data = {
                test : "abcd"
            }
    try {
        const response = await fetch(`/api/post?offset=${offset}&limit=10`, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
            body:  JSON.stringify(data)
        })
        return await response.json()
    } catch (error) {
        console.error(error);
    }
}




