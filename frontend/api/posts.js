
export async function addPostApi(postData) {
    try {
        const response = await fetch('http://localhost:8080/api/post', {
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
        const response = await fetch(`http://localhost:8080/api/post?offset=${offset}`, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        })
        return [response.status, await response.json()]
    } catch (error) {
        console.error("trying to fetch posts",error);
    }
}



export async function getCategories() {
    try {
        let response = await fetch(`http://localhost:8080/api/categories`, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        })
        return [response.status, await response.json()]
    } catch (error) {
        console.error("ERROR while trying to get posts categories: ", error)
    }
}




