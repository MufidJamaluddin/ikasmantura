import LoginProvider from "./LoginProvider";

const ApiUrl = process.env.PUBLIC_URL + '/api/v1'

const doRemoteLogin = () => fetch(`${ApiUrl}/auth`, {method: 'get'})
    .then(data => {
        if (data.status === 401) {
            throw new Error("Unauthorized User")
        }
        if (data.status < 200 || data.status > 399) {
            throw new Error("Connection Error")
        }
        return data.json()
    }).catch(err => {
        return null
    })


const authProvider = {
    // called when the user attempts to log in
    login: LoginProvider,

    // called when the user clicks on the logout button
    logout: () => {
        localStorage.removeItem('user')
        localStorage.removeItem('refresh')

        return fetch(`${ApiUrl}/auth`, { method: 'delete' })
            .then(_ => {

            })
    },

    // called when the API returns an error
    checkError: ({ status }) => {
        if (status === 401 || status === 403) {
            localStorage.removeItem('user')
            localStorage.removeItem('refresh')
            return Promise.reject();
        }
        return Promise.resolve();
    },

    // called when the user navigates to a new location, to check for authentication
    checkAuth: async () => {
        let dtUser = localStorage.getItem('user')
        let userData

        if(dtUser) {
            userData = JSON.parse(dtUser)
        }

        if (!userData) {
            userData = await doRemoteLogin()

            if(userData === null) {
                return Promise.reject()
            }

            localStorage.setItem('user', JSON.stringify(userData))
        }

        let expiration = userData.exp * 1000
        let currentDate = Date.now()

        if (currentDate >= expiration) {
            return await Promise.reject()
        }

        return await Promise.resolve()
    },

    // called when the user navigates to a new location, to check for permissions / roles
    getPermissions: async () => {
        let user = localStorage.getItem('user')
        let dtUser
        if(user === null) {
            dtUser = await doRemoteLogin()
            if(dtUser === null) {
                return Promise.reject()
            }
            localStorage.setItem('user', JSON.stringify(dtUser))
        } else {
            dtUser = JSON.parse(user)
        }
        return Promise.resolve(dtUser.role)
    },

    getIdentity: () => {
        try {
            const {
                id, fullName, avatar = '/static/img/ringga'
            } = JSON.parse(localStorage.getItem('user'));
            return Promise.resolve({ id, fullName, avatar });
        } catch (error) {
            return Promise.reject(error);
        }
    }
}

export default authProvider
