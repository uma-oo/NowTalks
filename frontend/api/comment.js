
export async function getComments() {
    const response = await fetch('/api/comments');
    return response.json();
}

export async function createComment(commentData) {
    const response = await fetch('/api/comments', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(commentData),
    });
    return response.json();
}