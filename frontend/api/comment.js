


export async function getComments(postId, offset) {
    try {
        const response = await fetch(`http://localhost:8080/api/comment?post=${postId}&offset=${offset}`);
        return [response.status, await response.json()]
    } catch (error) {
        console.error("Error While Trying to get comments")
    }
}

export async function addComment(commentData) {
    try {
        const response = await fetch('http://localhost:8080/api/comment', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(commentData)
        });
        return [response.status, await response.json()];
    } catch (error) {
        console.error(error)
    }
}