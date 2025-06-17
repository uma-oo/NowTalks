
const API_URL = '/api/messages';

export async function getMessages(offset, receiver_id) {
    try {
        const response = await fetch(`${API_URL}?offset=${offset}&receiver_id=${receiver_id}&type=new`, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        })
        return [response.status, await response.json()];
    } catch (error) {
        console.error("error trying to get users", error)
    }
}


// we don't need this anymore 
export async function markMessagesRead(receiver_id) {
    try {
        const response = await fetch(`${API_URL}?receiver_id=${receiver_id}`, {
            method: 'PATCH',
            headers: { 'Content-Type': 'application/json' }
        })
        return [response.status, await response.json()];
    } catch (error) {
        console.error("error trying to mark messages as read", error)
    }
}