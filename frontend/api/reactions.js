

export function LikeEntity(entity_id , entity_type) {
    try {
        const response = fetch('/api/react/like', {
            method: "POST",
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({entity_id,entity_type})
        })
        return response
    } catch (error) {
        console.error(error)
    }
}