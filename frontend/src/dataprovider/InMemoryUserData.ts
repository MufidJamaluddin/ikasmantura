import {ParseJwt} from "../utils/Jwt";

const inMemoryUserData = () => {
    let logoutEventName = 'ra-logout'
    let inMemoryUser = null
    let inMemoryRefreshToken = null
    let refreshTimeOutId = null
    let refreshEndpoint

    window.addEventListener('storage', (event) => {
        if (event.key === logoutEventName) {
            inMemoryUser= null;
        }
    })

    const getUser = () => inMemoryUser

    const getRefreshToken = () => inMemoryRefreshToken

    const setRefreshEndpoint = (endpoint) => {
        refreshEndpoint = endpoint
    }

    // This countdown feature is used to renew the JWT in a way that is transparent to the user.
    // before it's no longer valid
    const refreshToken = (delay) => {
        refreshTimeOutId = window.setTimeout(
            getRefreshedToken,
            delay * 1000 - 50000
        ); // Validity period of the token in seconds, minus 50 seconds
    };

    const abortRefreshToken = () => {
        if (refreshTimeOutId) {
            window.clearTimeout(refreshTimeOutId);
        }
    };

    // The method makes a call to the refresh-token endpoint
    // If there is a valid cookie, the endpoint will return a fresh jwt.
    const getRefreshedToken = () => {
        const request = new Request(refreshEndpoint, {
            method: 'PUT',
            headers: new Headers({ 'Content-Type': 'application/json' }),
            credentials: 'include',
        });
        return fetch(request)
            .then((response) => {
                if (response.status !== 200) {
                    eraseData()
                    global.console.log(
                        'Failed to renew the jwt from the refresh token.'
                    );
                    return { token: null };
                }
                return response.json();
            })
            .then(({ token, refreshToken }) => {
                if (token) {
                    let userData = ParseJwt(token)

                    setUser(userData);
                    setRefreshToken(refreshToken)
                    return true;
                }

                return false;
            });
    };

    const setUser = (user) => {
        if(user){
            inMemoryUser = user
            refreshToken(user.exp)
            return true
        }
        return false
    }

    const setRefreshToken = (refreshToken) => {
        inMemoryRefreshToken = refreshToken
        return true
    }

    const eraseData = () => {
        inMemoryUser = null
        inMemoryRefreshToken = null
        abortRefreshToken()
        window.localStorage.setItem(logoutEventName, Date.now().toString());
        return true
    }

    return {
        getUser, setUser,
        getRefreshToken, setRefreshToken, setRefreshEndpoint,
        eraseData
    }
}

export default inMemoryUserData()
