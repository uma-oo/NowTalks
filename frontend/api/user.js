export async function getUsers(offset) {
    try {
        const response = await fetch(`http://localhost:8080/api/users?offset=${offset}`);
        return [ response.status , await response.json()];
    } catch (error) {
        console.error("error trying to get users", error)
    }
}

export async function createUser(data) {
    try {
        const response = await fetch('http://localhost:8080/api/user/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        });
        return [response.status, await response.json()];
    } catch (error) {
        console.error(`Error trying to register user: ${error}`)
    }
}

export async function loginUser(data) {
    try {
        const response = await fetch('http://localhost:8080/api/user/login', {
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(data)
        })
        return [response.status, await response.json()]
    } catch (error) {
        console.error(`Error trying to login${error}`);
    }
}

export async function logoutUser() {
    try {
        const response = await fetch("http://localhost:8080/api/user/logout", {
            method: "POST",
            headers: { 'Content-Type': 'application/json' }
        })
        console.log()
        return response.status
    } catch (error) {
        console.error(`Error trying to logout: ${error}`)
    }
}

export async function isLoggedIn() {
    try {
        let response = await fetch("http://localhost:8080/api/loggedin")
        return await response.json()
    } catch (error) {
        console.error(error)
    }
}