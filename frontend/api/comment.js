


// You have to specify the id assi ayoub 
export async function getComments(id) {
    const response = await fetch(`/api/comment?postId=${id}`);
    return response.json();
}



// the same applies for the creation !!! khass id f l api 
export async function createComment(commentData) {
    const response = await fetch('/api/comment', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(commentData),
    });
    return response.json();
}