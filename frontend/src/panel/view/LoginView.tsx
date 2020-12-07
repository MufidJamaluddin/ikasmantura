import React from "react";

import {Notification, useLogin, useNotify} from 'react-admin'
import {ThemeProvider} from '@material-ui/styles'

const LoginView = ({ parenthistory, theme }) => {

    const login = useLogin();
    const notify = useNotify();

    const onSubmit = (e: React.FormEvent<HTMLFormElement>|any) => {
        e.preventDefault();

        let formData = new FormData(e.target)

        login({
            username: formData.get('username'),
            password: formData.get('password')
        })
        .catch(() => notify('Invalid email or password'));
    };

    const onHomeClick = (e: any) => {
        e.preventDefault()
        parenthistory.push('/')
    }

    return (
        <ThemeProvider theme={theme}>
            <div className="c-g-banner c-f-full-height">
                <form onSubmit={onSubmit} className="c-container-box c-center-box primary">
                    <div className="c-container c-text-center">

                        <label htmlFor="username" className="lead-sm"><b>Username</b></label>
                        <input type="text" className="c-input-box"
                               placeholder="Enter Username" name="username" required/>

                        <label htmlFor="password" className="lead-sm"><b>Password</b></label>
                        <input type="password" className="c-input-box"
                               placeholder="Enter Password" name="password" required/>

                        <button type="button" className="c-button info" onClick={onHomeClick}>Back to Home</button>
                        &nbsp;
                        <button type="submit" className="c-button info">Login</button>

                        <br/>

                        <label>
                            <input type="checkbox" name="remember"/> Remember me
                        </label>
                    </div>
                </form>
            </div>
            <Notification />
        </ThemeProvider>
    )
}

export default LoginView
