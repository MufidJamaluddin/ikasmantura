import inMemoryUserData from './InMemoryUserData'
import {ParseJwt} from "../utils/Jwt";

export default function LoginProvider({ username, password }) {
    const request = new Request('/api/v1/auth', {
        method: 'POST',
        body: JSON.stringify({ username: username.trim(), password: password.trim() }),
        headers: new Headers({ 'Content-Type': 'application/json' }),
    })

    inMemoryUserData.setRefreshEndpoint('/api/v1/auth')

    return fetch(request)
        .then(response => {
            if (response.status < 200 || response.status >= 300) {
                throw new Error(response.statusText);
            }
            return response.json();
        })
        .then(({ token, refreshToken }) => {
            let userData = ParseJwt(token)

            inMemoryUserData.setUser(userData)
            inMemoryUserData.setRefreshToken(refreshToken)
        });
}
