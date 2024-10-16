import React, {useEffect, useState} from "react";

import {Image, Nav, Navbar, NavDropdown, Container} from "react-bootstrap";

import './Header.css'

import { Link } from "react-router-dom";

import AuthProvider from "../../dataprovider/authProvider";
import {useStore} from "../models";

function UserMenu(props) {

    const [isLogin, setIsLogin] = useState(false)
    const [userData, setUserData] = useState(null)

    useEffect(() => {
        let life = true

        const checkLogin = async () => {
            let { isLogin } = await AuthProvider.checkAuth()
                .then(() => ({ isLogin: true }))
                .catch(() => ({ isLogin: false }));

            if(life) {
                setIsLogin(isLogin)
            }
        }

        checkLogin()

        return () => {
            life = false
        }
    }, [])

    return (
        (isLogin || userData) ? (
            <Nav.Link as={Link} to={"/panel"}>
                <span><i className="fas fa-user"/></span>
            </Nav.Link>
        ) : (
            <NavDropdown title="Alumni" id="basic-nav-dropdown">
                <NavDropdown.Item as={Link} to={"/register"}>
                    Daftar
                </NavDropdown.Item>
                <NavDropdown.Item as={Link} to={"/login"}>
                    Masuk
                </NavDropdown.Item>
            </NavDropdown>
        )
    )
}

export default function Header(props)
{
    const [state, actions] = useStore('AboutModel')

    useEffect(() => {
        // @ts-ignore
        actions.init()
    }, [])

    let data: any = state?.data || {}

    let {
        title
    } = data

    return (
        <Navbar expand="lg" sticky="top">
            <Container>
                <Navbar.Brand href="#home" as={Link} to={"/"}>
                    <Image src={process.env.PUBLIC_URL + "/static/img/logo.svg"}
                           rounded={true}
                           className="logo-brand"/>
                    {title}
                </Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav"/>
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="ml-auto">
                        <NavDropdown title="Profil" id="basic-nav-dropdown">
                            <NavDropdown.Item as={Link} to={"/about"}>
                                Tentang {title}
                            </NavDropdown.Item>
                            <NavDropdown.Item as={Link} to={"/organization"}>
                                Struktur Organisasi {title}
                            </NavDropdown.Item>
                        </NavDropdown>
                        <NavDropdown title="Agenda" id="basic-nav-dropdown">
                            <NavDropdown.Item as={Link} to={"/events"}>
                                Kalender Kegiatan
                            </NavDropdown.Item>
                            <NavDropdown.Item as={Link} to={"/events_list"}>
                                Daftar Kegiatan
                            </NavDropdown.Item>
                        </NavDropdown>
                        <Nav.Link as={Link} to={"/articles"}>
                            Artikel
                        </Nav.Link>
                        <Nav.Link as={Link} to={"/gallery"}>Galeri</Nav.Link>
                        <UserMenu />
                    </Nav>
                </Navbar.Collapse>
            </Container>
        </Navbar>
    )
}
