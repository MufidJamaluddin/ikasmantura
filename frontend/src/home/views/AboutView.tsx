import React, {useContext, useEffect} from "react";

import {Col, Row} from "react-bootstrap";
import DOMPurify from "../../utils/Sanitizer";
import { useStore } from "../models";
import {ThemeContext} from "../component/PageTemplate";

export default function AboutView(props) {

    const [state, actions] = useStore('AboutModel')
    const theme = useContext(ThemeContext);

    useEffect(() => {
        // @ts-ignore
        actions.init()
        theme.setHeader({ title: 'Tentang Kami', showTitle: true })
    }, [])

    let {
        description, vision, mission
    } = state.data ?? {}

    return (
        <Row className="padding-cont">
            <Col sm={12}>
                <hr/>
                <h3 className="text-center">Deskripsi</h3>
                <div className="text-justify" dangerouslySetInnerHTML={
                    {__html: description ? DOMPurify.sanitize(description) : '-'}}/>
            </Col>
            <Col sm={12}>
                <hr/>
                <h3 className="text-center">Visi</h3>
                <div className="text-justify" dangerouslySetInnerHTML={
                    {__html: vision ? DOMPurify.sanitize(vision) : '-'}}/>
            </Col>
            <Col sm={12}>
                <hr/>
                <h3 className="text-center">Misi</h3>
                <div className="text-justify" dangerouslySetInnerHTML={
                    {__html: mission ? DOMPurify.sanitize(mission) : '-'}}/>
                <hr/>
            </Col>
        </Row>
    )
}
