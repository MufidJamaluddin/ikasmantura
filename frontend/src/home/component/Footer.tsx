import React from "react";
import {Col, Container, Row} from "react-bootstrap";

import './Footer.css'
import RegeImage from "./RegeImage";

export default function Footer() {
    return (
        <div className="footerContainer">
            <RegeImage/>
            <footer>
                <Container className="footer">
                    <Row>
                        <Col lg={6} className="h-100 text-center text-lg-left my-auto">
                            <ul className="list-inline mb-2">
                                <li className="list-inline-item">
                                    IKA SMAN SITURAJA
                                </li>
                                <li className="list-inline-item">&sdot;</li>
                                <li className="list-inline-item">
                                    <a href="mailto:info@ikasmansituraja.org">
                                        Kontak
                                    </a>
                                </li>
                            </ul>
                            <p className="text-muted small mb-4 mb-lg-0">
                                &copy; IKA SMAN SITURAJA 2020. Hak Cipta Dilindungi Oleh Undang-Undang.
                            </p>
                        </Col>
                        <Col lg={6} className="h-100 text-center text-lg-right my-auto">
                            <ul className="list-inline mb-0">
                                <li className="list-inline-item mr-3">

                                    <i className="fab fa-facebook fa-2x fa-fw"/>

                                </li>
                                <li className="list-inline-item mr-3">

                                    <i className="fab fa-twitter-square fa-2x fa-fw"/>

                                </li>
                                <li className="list-inline-item">

                                    <i className="fab fa-instagram fa-2x fa-fw"/>

                                </li>
                            </ul>
                        </Col>
                    </Row>
                </Container>
            </footer>
        </div>
    )
}
