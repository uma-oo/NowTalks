
const API_URL = 'http://localhost:8080/api/messages';

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
