
export async function getUsers() {
    const response = await fetch('/api/user');
    return response.json();
}

export async function createUser(data) {
    const response = await fetch('/api/user', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(data),
    });
    return response.json();
}


export async function loginUser(data) {
    console.log(data)
    try {
        const response = await fetch('api/user/login', {
            method : "POST",
            headers : {'Content-Type': 'application/json'},
            body: JSON.stringify(data)
        })
        return response.json()
    } catch (error) {
        console.error(error)
        return response.json()
    }
}