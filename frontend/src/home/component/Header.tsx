import React, {PureComponent} from "react";

import {Image, Nav, Navbar, NavDropdown, Container} from "react-bootstrap";

import './Header.css'

import { Link } from "react-router-dom";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";

import {NotificationManager} from 'react-notifications';
import WEB_INFO from "../config";
import AuthProvider from "../../dataprovider/authProvider";

export default class Header extends PureComponent<{}, {topics: Array<any>, isLogin: boolean}>
{
    constructor(props) {
        super(props);
        this.state = {
            topics: [],
            isLogin: false,
        }
    }

    async updateTopics() {
        let dataProvider = DataProviderFactory.getDataProvider()
        return await dataProvider.getList("article_topics", {
            pagination: {
                page: 1,
                perPage: 100,
            },
            sort: {
                field: 'id',
                order: 'ASC'
            },
            filter: {},
        }).then(resp => {
            return {topics: resp.data}
        }, error => {
            NotificationManager.error(error.message, error.name);
            return {}
        })
    }

    componentDidMount()
    {
        try
        {
            Promise.all([
                this.updateTopics(),
                AuthProvider.checkAuth()
                    .then(() => ({ isLogin: true }))
                    .catch(() => ({ isLogin: false })),
            ]).then((values: Array<any>) => {
                let newState = { ...this.state }
                values.forEach(item => {
                    newState = { ...newState, ...item }
                })
                this.setState(newState)
            })
        }
        catch (e)
        {
            NotificationManager.error('Koneksi Internet Terputus!', 'Error Koneksi');
        }
    }

    render() {
        return (
            <Navbar expand="lg" sticky="top">
                <Container>
                    <Navbar.Brand href="#home" as={Link} to={"/"}>
                        <Image src={process.env.PUBLIC_URL + "/static/img/logo.svg"}
                               rounded={true}
                               className="logo-brand"/>
                        {WEB_INFO.name}
                    </Navbar.Brand>
                    <Navbar.Toggle aria-controls="basic-navbar-nav"/>
                    <Navbar.Collapse id="basic-navbar-nav">
                        <Nav className="ml-auto">
                            <NavDropdown title="Profil" id="basic-nav-dropdown">
                                <NavDropdown.Item as={Link} to={"/about"}>
                                    Tentang {WEB_INFO.name}
                                </NavDropdown.Item>
                                <NavDropdown.Item as={Link} to={"/organization"}>
                                    Struktur Organisasi {WEB_INFO.name}
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
                            <NavDropdown title="Artikel" id="basic-nav-dropdown">
                                {
                                    this.state.topics.map(item => {
                                        return <NavDropdown.Item
                                            key={item.id}
                                            as={Link}
                                            to={{
                                                pathname: '/articles',
                                                state: {
                                                    topicId: item.id
                                                }
                                            }}
                                        >
                                            {item.name}
                                        </NavDropdown.Item>
                                    })
                                }
                            </NavDropdown>
                            <Nav.Link as={Link} to={"/gallery"}>Galeri</Nav.Link>
                            {
                                this.state.isLogin ? (
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
                            }
                        </Nav>
                    </Navbar.Collapse>
                </Container>
            </Navbar>
        )
    }
}
