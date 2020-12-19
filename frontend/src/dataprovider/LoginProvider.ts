export default function LoginProvider({ username, password }) {
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
        .then(({ token, data }) => {
            if(
                process.env.NODE_ENV === "development"
                || process.env.NODE_ENV === "test"
            )
            {
                localStorage.setItem('token', token);
            }
            localStorage.setItem('data', JSON.stringify(data));
        });
}
