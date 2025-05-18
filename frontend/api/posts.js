
export async function addPostApi(postData) {
    try {
        const response = await fetch('/api/post', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
        })
        return await response.json()
    } catch (error) {
        console.error(error)
    }
}

export async function getPostsApi(offset=0,filterData) {
    let categories = filterData.categories.join("&categories=")
    console.log(categories)
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


export async function getCategories() {
    try {
        let response = await fetch(`/api/categories`)
        console.log(response)
    } catch (error) {
        console.error("ERROR while trying to get posts categories: ", error)
    }
}

