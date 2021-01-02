import React, {useContext, useEffect} from "react";

import moment from "moment";
import 'moment/locale/id';

import {Card, Col, Row} from "react-bootstrap";
import Image from "../component/Image";
import DOMPurify from "../../utils/Sanitizer";
import {useStore} from "../models";
import {ThemeContext} from "../component/PageTemplate";
import {ArticleItem} from "../models/ArticleModel";

export default function ArticlesItemView(props)
{
    const id = props.match.params.id

    const [state, actions] = useStore('ArticleModel')
    const theme = useContext(ThemeContext)

    let item: ArticleItem = { ...state?.selected }

    useEffect(() => {
        theme.setHeader({ title: item?.title ?? 'Artikel', showTitle: true })
        actions.getArticleById(id)
        return actions.reset
    }, [id])

    let {
        title,
        image,
        body,
        createdByName,
        createdAt,
    } = item

    let fCreatedAt = moment(createdAt).format('LLLL');

    return (
        <Row>
            <Col md={{span:10, offset:1}}>
                <Card>
                    <Image
                        className="card-img-top"
                        src={image}
                        alt={title}/>
                    <Card.Title>
                        <h1 className="text-center">{title}</h1>
                    </Card.Title>
                    <Card.Text className="lead text-center">
                        <b>Oleh: {createdByName ?? 'Kakak Anonim'}</b> &nbsp;
                        <small>Pada: {fCreatedAt}</small>
                    </Card.Text>
                    <Card.Body>
                        <div dangerouslySetInnerHTML={{
                            __html: body ? DOMPurify.sanitize(body) : ''
                        }} />
                    </Card.Body>
                </Card>
            </Col>
        </Row>
    )
}
