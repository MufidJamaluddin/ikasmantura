import LoginProvider from "./LoginProvider";
import inMemoryUserData from './InMemoryUserData'

const ApiUrl = process.env.PUBLIC_URL + '/api/v1'

const authProvider = {
    // called when the user attempts to log in
    login: LoginProvider,

    // called when the user clicks on the logout button
    logout: () => {
        return fetch(`${ApiUrl}/auth`, { method: 'delete' })
            .then(_ => {
                if(!inMemoryUserData.eraseData()){
                    throw new Error("Error Invalidate User Data")
                }
            })
    },

    // called when the API returns an error
    checkError: ({ status }) => {
        if (status === 401 || status === 403) {
            inMemoryUserData.eraseData()
            return Promise.reject();
        }
        return Promise.resolve();
    },

    // called when the user navigates to a new location, to check for authentication
    checkAuth: async () => {
        let userData = inMemoryUserData.getUser()

        if (!userData) {
            userData = await fetch(`${ApiUrl}/auth`, {method: 'get'})
                .then(data => {
                    if (data.status === 401) {
                        throw new Error("Unauthorized User")
                    }
                    if (data.status < 200 || data.status > 399) {
                        throw new Error("Connection Error")
                    }
                    return data.json()
                })
        }

        let expiration = userData.exp * 1000
        let currentDate = Date.now()

        if (currentDate >= expiration) {
            return await Promise.reject()
        }

        return await Promise.resolve()
    },

    // called when the user navigates to a new location, to check for permissions / roles
    getPermissions: () => {
        const user = inMemoryUserData.getUser()
        if(user === null) {
            return Promise.reject()
        }
        return Promise.resolve(user.role)
    },

    getData: () => {
        return inMemoryUserData.getUser()
    },

    getIdentity: () => {
        try {
            const { id, fullName, avatar = '/static/img/ringga.jpg' }
                = inMemoryUserData.getUser()

            return Promise.resolve({ id, fullName, avatar });
        } catch (error) {
            return Promise.reject(error);
        }
    }
}

export default authProvider
