import LoginProvider from "./LoginProvider";

const authProvider = {
    // called when the user attempts to log in
    login: LoginProvider,
    // called when the user clicks on the logout button
    logout: () => {
        localStorage.removeItem('data');
        localStorage.removeItem('token');
        return Promise.resolve();
    },
    // called when the API returns an error
    checkError: ({ status }) => {
        if (status === 401 || status === 403) {
            localStorage.removeItem('data');
            localStorage.removeItem('token');
            return Promise.reject();
        }
        return Promise.resolve();
    },
    // called when the user navigates to a new location, to check for authentication
    checkAuth: () => {
        let dataStr = localStorage.getItem('data')
        if(dataStr)
        {
            let data = JSON.parse(dataStr)
            if(Date.now() >= data.exp * 1000)
            {
                return Promise.reject()
            }
            return Promise.resolve()
        }
        return Promise.reject()
    },
    // called when the user navigates to a new location, to check for permissions / roles
    getPermissions: () => Promise.resolve(),

    getData: () => {
        return JSON.parse(localStorage.getItem("data") ?? "{}")
    }
}

export default authProvider
