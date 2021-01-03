import * as React from "react";
import {Link} from "react-router-dom";

import ReactPaginate from 'react-paginate';

import {Button, Card, Col, Form, Row} from "react-bootstrap";

import Image from "../component/Image";

import {useStore} from "../models";
import {useContext, useEffect, useState} from "react";

import {ThemeContext} from "../component/PageTemplate";
import {ArticleTopicItem} from "../models/ArticleTopicModel";
import {ArticleItem} from "../models/ArticleModel";
import {strip_tags} from "../../utils/Security";

function ArticleItemView(props) {
    let {
        id, thumbnail, title, body
    } = props

    body = strip_tags(body)

    return (
        <div className="col-auto" key={id}>
            <Card className="h-100" style={{'width':'15rem'}}>
                <Image
                    className="card-img-top"
                    src={thumbnail}
                    alt={title}/>
                <Card.Body>
                    <Card.Title><h4>{title}</h4></Card.Title>
                    <p>
                        {body}
                        &nbsp;
                        <Link to={`articles/${id}`}>
                            <small><b>Read More...</b></small>
                        </Link>
                    </p>
                </Card.Body>
            </Card>
        </div>
    )
}

function ArticleSearch(props) {
    let {
        history, topicId, topics, title, setTitle
    } = props

    const [cTitle, cSetTitle] = useState(title)

    function onSubmit(e) {
        if(e.cancelable){
            e.preventDefault()
        }
        setTitle(cTitle)
    }

    return (
        <Form onSubmit={onSubmit}>
            <Form.Row className="justify-content-center">
                <Form.Group as={Col} md="5">
                    <Form.Label htmlFor="topicSearch" srOnly>Topik</Form.Label>
                    <Form.Control
                        as="select"
                        id="topicSearch"
                        required={true}
                        value={topicId}
                        onChange={(e) => {
                            let state = {
                                topicId: e.target.value
                            };

                            history.replace({ ...history.location, state});
                        }}
                        name="topicId">
                        <option value={0}>
                            Semua
                        </option>
                        {
                            topics.map(item => (
                                <option key={item.id} value={item.id}>
                                    {item.name}
                                </option>
                            ))
                        }
                    </Form.Control>
                </Form.Group>

                <Form.Group as={Col} md="5">
                    <Form.Label htmlFor="titleSearch" srOnly>Judul</Form.Label>
                    <Form.Control
                        id="titleSearch"
                        type="text"
                        name="title"
                        maxLength={100}
                        minLength={3}
                        required={false}
                        onChange={(e) => {
                            cSetTitle(e.target.value)
                        }}
                        value={cTitle}
                        autoComplete="off"
                        placeholder="Judul Berita yang Anda Cari" />
                </Form.Group>

                <Col md="2">
                    <Button type="submit" variant="info" size="sm">
                        Cari
                    </Button>
                </Col>
            </Form.Row>
        </Form>
    )
}

function ArticlesView(props)
{
    const cTopicId = props.location?.state?.topicId ?? 0
    const history = props.history

    const [stateTopic, actionsTopic] = useStore('ArticleTopicModel')
    const [stateArticle, actionsArticle] = useStore('ArticleModel')
    const [cTitle, setTitle] = useState('')

    const theme = useContext(ThemeContext)

    let articles: Array<ArticleItem> = stateArticle?.data ?? []
    let articlesTotal = stateArticle?.total ?? 0
    let contentPerPage = 4
    let topics: Array<ArticleTopicItem> = stateTopic?.data ?? []

    function searchArticles({ topicId, title }) {
        actionsArticle.getArticles({
            page: 1, perPage: contentPerPage, filter: { topicId: topicId, title: title } })
    }

    function handlePageChange({selected}) {
        actionsArticle.getArticles(
            { page: selected+1, perPage: contentPerPage, filter: { topicId: cTopicId } })
    }

    useEffect(() => {
        theme.setHeader({ title: 'Artikel', showTitle: true })

        actionsTopic.init()
        searchArticles( { topicId: cTopicId, title: cTitle })

        return actionsArticle.reset
    }, [cTopicId, cTitle])

    if(!stateArticle?.data) {
        return <div className="c-center-box c-loader"/>
    }

    return (
        <Row>
            <Col md={{span:8, offset:2}}>
                <ArticleSearch
                    perPage={contentPerPage}
                    topicId={cTopicId}
                    history={history}
                    topics={topics}
                    title={cTitle}
                    setTitle={setTitle}
                />
            </Col>
            <Col md={12}>
                <div className="row justify-content-center mb-3">
                {
                    articles.map(item => (<ArticleItemView key={item.id} {...item} />))
                }
                </div>
            </Col>
            <Col md={12}>
                <ReactPaginate
                    previousLabel={'previous'}
                    nextLabel={'next'}
                    breakLabel={'...'}
                    breakClassName={'break-me'}
                    pageCount={ articlesTotal / contentPerPage }
                    marginPagesDisplayed={2}
                    pageRangeDisplayed={5}
                    onPageChange={handlePageChange}
                    containerClassName={'pagination justify-content-center'}
                    subContainerClassName={'pages pagination'}
                    activeClassName={'active'}
                />
            </Col>
        </Row>
    )
}

export default ArticlesView
