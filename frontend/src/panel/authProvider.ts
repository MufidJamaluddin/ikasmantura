const authProvider = {
    // called when the user attempts to log in
    login: ({ username, password }) => {
        const request = new Request('/api/v1/auth', {
            method: 'POST',
            body: JSON.stringify({ username: username.trim(), password: password.trim() }),
            headers: new Headers({ 'Content-Type': 'application/json' }),
        })
        return fetch(request)
            .then(response => {
                if (response.status < 200 || response.status >= 300) {
                    throw new Error(response.statusText);
                }
                return response.json();
            })
            .then(({ data, token }) => {
                console.log(data)
                localStorage.setItem('data', JSON.stringify(data));
                localStorage.setItem('token', token);
            });
    },
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
        return localStorage.getItem('token')
            ? Promise.resolve()
            : Promise.reject();
    },
    // called when the user navigates to a new location, to check for permissions / roles
    getPermissions: () => Promise.resolve(),

    getData: () => {
        return JSON.parse(localStorage.getItem("data") ?? "{}")
    }
}

export default authProvider