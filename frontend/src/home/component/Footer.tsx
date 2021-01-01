import React, {useEffect} from "react";
import {Col, Container, Row} from "react-bootstrap";

import './Footer.css'
import RegeImage from "./RegeImage";
import {useStore} from "../models";

export default function Footer() {
    const [state, actions] = useStore('AboutModel')

    useEffect(() => {
        // @ts-ignore
        actions.init()
    }, [])

    let {
        title, facebook, twitter, instagram, email
    } = state.data ?? {}

    return (
        <div className="footerContainer">
            <RegeImage/>
            <footer>
                <Container className="footer">
                    <Row>
                        <Col lg={6} className="h-100 text-center text-lg-left my-auto">
                            <ul className="list-inline mb-2">
                                <li className="list-inline-item">
                                    {title}
                                </li>
                                <li className="list-inline-item">&sdot;</li>
                                {
                                    email &&
                                    <li className="list-inline-item">
                                        <a href={`mailto:${email}`}>
                                            Kontak
                                        </a>
                                    </li>
                                }
                            </ul>
                            <p className="text-muted small mb-4 mb-lg-0">
                                &copy; {title ?? 'Mufid'} 2020. Hak Cipta Dilindungi Oleh Undang-Undang.
                            </p>
                        </Col>
                        <Col lg={6} className="h-100 text-center text-lg-right my-auto">
                            <ul className="list-inline mb-0">
                                {
                                    facebook &&
                                    <li className="list-inline-item mr-3">
                                        <a href={`https://facebook.com/${facebook}`}
                                           rel="noreferrer"
                                           target="_blank">
                                            <i className="fab fa-facebook fa-2x fa-fw"/>
                                        </a>
                                    </li>
                                }
                                {
                                    twitter &&
                                    <li className="list-inline-item mr-3">
                                        <a href={`https://twitter.com/${twitter}`}
                                           rel="noreferrer"
                                           target="_blank">
                                            <i className="fab fa-twitter-square fa-2x fa-fw"/>
                                        </a>
                                    </li>
                                }
                                {
                                    instagram &&
                                    <li className="list-inline-item">
                                        <a href={`https://instagram.com/${instagram}`}
                                           rel="noreferrer"
                                           target="_blank">
                                            <i className="fab fa-instagram fa-2x fa-fw"/>
                                        </a>
                                    </li>
                                }
                            </ul>
                        </Col>
                    </Row>
                </Container>
            </footer>
        </div>
    )
}
