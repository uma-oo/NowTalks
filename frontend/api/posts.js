
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

export async function getPostsApi(filterData,offset) {
    console.log(filterData)
    let categoriesQuery = filterData.categories ? `&categories=${filterData.categories}` : ""
    let likedpostsQuery = filterData.likedPosts ? `&likedposts=${filterData.likedPosts}` : ""
    let createdPosts = filterData.createdPosts ? `&createdposts=${filterData.createdPosts}`: ""

    try {
        const response = await fetch(`/api/post?offset=${offset}&limit=10&categories=${filterData.categories}&likedposts=${filterData.likedPosts}&createdPosts=${filterData.createdPosts}`, {
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
        return await response.json()
    } catch (error) {
        console.error("ERROR while trying to get posts categories: ", error)
    }
}

