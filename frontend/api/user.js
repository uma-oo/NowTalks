
export async function getUsers() {
    const response = await fetch('api/user');
    return response.json();
}

export async function createUser(data) {
    try {
        const response = await fetch('api/user/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        });
        return response;
    } catch (error) {
        console.error(error)
        return response.json()
    }
}

export async function loginUser(data) {
    try {
        const response = await fetch('api/user/login', {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        })
        return response
    } catch (error) {
        console.error(error);
        return response.json()
    }
}

export async function logoutUser() {
    try {
        await fetch("api/user/logout", {
            method: "POST",
            headers: { 'Content-Type': 'application/json' }
        })
    } catch (error) {
        throw error
    }
}