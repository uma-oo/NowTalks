


// You have to specify the id assi ayoub 
export async function getComments(postId, offset) {
    try {
        const response = await fetch(`/api/comment?post=${postId}&offset=${offset}`);
        return [response.status, await response.json()]
    } catch (error) {
        console.error("Error While Trying to get comments")
    }
}

// {
// "post_id": 1,
// "content": "Totally agree! I started eating more whole foods and feel amazing."
// }
export async function addComment(commentData) {
    console.log(commentData)
    try {
        const response = await fetch('/api/comment', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(commentData)
        });
        return [response.status, await response.json()];
    } catch (error) {
        console.error(error)
    }
}