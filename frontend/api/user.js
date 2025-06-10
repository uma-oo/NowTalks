import { closeConnection } from "../websocket.js";

export async function getUsers(offset) {
    try {
        const response = await fetch(`api/users?offset=${offset}`);
        return [ response.status , await response.json()];
    } catch (error) {
        console.error("error trying to get users", error)
    }


}

export async function createUser(data) {
    try {
        const response = await fetch('api/user/register', {
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
        const response = await fetch('api/user/login', {
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
        await fetch("api/user/logout", {
            method: "POST",
            headers: { 'Content-Type': 'application/json' }
        })
        closeConnection()
    } catch (error) {
        console.error(`Error trying to logout: ${error}`)
    }
}


export async function isLoggedIn() {
    try {
        let response = await fetch("api/loggedin")
        return await response.json()
    } catch (error) {
        console.error(error)
    }
}