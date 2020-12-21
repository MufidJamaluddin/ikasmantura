import React from "react";

import {useLogin} from 'react-admin'
import {ThemeProvider} from '@material-ui/styles'

import {NotificationManager} from 'react-notifications';
import { NotificationContainer } from 'react-notifications';

const LoginView = ({ parenthistory, theme, location }) => {

    const login = useLogin();

    let {
        username = '',
        password = ''
    } = location?.state ?? {}

    if(username)
    {
        if(password)
        {
            if(username.length > 3 && password.length > 3)
            {
                login({
                    username: username,
                    password: password
                }).catch(_ => {
                    NotificationManager.error(
                        'Username/Password salah atau Koneksi Bermasalah!', 'Login Gagal')
                });
            }
        }
    }

    const onSubmit = (e: React.FormEvent<HTMLFormElement>|any) => {
        e.preventDefault();

        let formData = new FormData(e.target)

        login({
            username: formData.get('username'),
            password: formData.get('password')
        }).catch(_ => {
            NotificationManager.error(
                'Username/Password salah atau Koneksi Bermasalah!', 'Login Gagal')
        });
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
                               placeholder="Enter Username" name="username" required minLength={3} maxLength={100}/>

                        <label htmlFor="password" className="lead-sm"><b>Password</b></label>
                        <input type="password" className="c-input-box"
                               placeholder="Enter Password" name="password" required minLength={3} maxLength={100}/>

                        <button type="button" className="c-button info" onClick={onHomeClick}>Kembali</button>
                        &nbsp;
                        <button type="submit" className="c-button info">Login</button>

                    </div>
                </form>
            </div>
            <NotificationContainer/>
        </ThemeProvider>
    )
}

export default LoginView
