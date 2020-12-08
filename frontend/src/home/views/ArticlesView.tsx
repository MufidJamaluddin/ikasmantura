import * as React from "react";
import {Link} from "react-router-dom";

import {PaginationPayload, SortPayload} from "ra-core/src/types";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";

import {NotificationManager} from 'react-notifications';

import ReactPaginate from 'react-paginate';

import RegeTitle from "../component/RegeTitle";
import {Button, Card, Col, Container, Form, Row} from "react-bootstrap";
import {RouteComponentProps} from "react-router";
import Image from "../component/Image";

interface ArticleItem {
    title?: string
    createdAt_gte?: string|Date,
    createdAt_lte?: string|Date
}

interface ArticlesViewState
{
    total: number;
    pagination: PaginationPayload;
    sort: SortPayload;
    filter: ArticleItem|any;
    data: Array<any>;
    topics: Array<any>;
    loading: boolean;
    selectedTopic?: string|number;
}

export default class ArticlesView
    extends React.PureComponent<RouteComponentProps<unknown, unknown, {topicId?: string|number}>, ArticlesViewState>
{
    constructor(props:any)
    {
        super(props);

        this.state = {
            pagination: {
                page: 1,
                perPage: 4,
            },
            sort: {
                field: 'id',
                order: 'ASC'
            },
            filter: {

            },
            loading: true,
            data: [],
            total: 0,
            topics: [],
        }

        this.handlePageChange = this.handlePageChange.bind(this)
        this.handleSearchForm = this.handleSearchForm.bind(this)
    }

    updateData()
    {
        if(this.state.loading)
        {
            const dataProvider = DataProviderFactory.getDataProvider()

            dataProvider.getList("articles", this.state).then(resp => {
                this.setState(state => {
                    let newState = {
                        data: resp.data,
                        loading: false,
                        total: resp.total,
                    }
                    return {...state, ...newState}
                })
            }, error => {
                NotificationManager.error(error, 'Get Data Error');

                this.setState(state => {
                    let newState = {loading: false}
                    return {...state, ...newState}
                })
            })
        }
    }

    updateTopics()
    {
        let dataProvider = DataProviderFactory.getDataProvider()
        dataProvider.getList("article_topics", {
            pagination: {
                page: 1,
                perPage: 100,
            },
            sort: {
                field: 'id',
                order: 'ASC'
            },
            filter: {
            },
        }).then(resp => {
            this.setState(state => {
                return {...state, topics: resp.data }
            })
        }, error => {
            NotificationManager.error(error, 'Get Data Error');
        })
    }

    componentDidMount()
    {
        this.updateTopics()
        this.updateData()
    }

    componentDidUpdate(prevProps: Readonly<{}>, prevState: Readonly<ArticlesViewState>, snapshot?: any)
    {
        if(
            prevState.filter !== this.state.filter
            || prevState.pagination !== this.state.pagination
        )
        {
            this.updateData()
        }
    }

    handlePageChange(data)
    {
        this.setState(state => {
            let newState = {
                loading: true, pagination: {
                    page: data.selected + 1, perPage: state.pagination.perPage
                }
            }
            return {...state, ...newState}
        })
    }

    handleSearchForm(e)
    {
        try
        {
            e.preventDefault()
        }
        catch (err)
        {

        }

        let formData = new FormData(e.target);
        let title = formData.get('title')

        let filter = {}

        if(title)
        {
            filter = {
                title: title
            }
        }

        this.setState(state => {
            let newState = {
                loading: true, filter: filter, pagination: {
                    page: 1,
                    perPage: state.pagination.perPage
                }
            }

            return {...state, ...newState}
        })

        return false
    }

    render()
    {
        let topicId = this.props.location?.state?.topicId ?? 1
        const history = this.props.history

        return (
            <>
                <RegeTitle>
                    <h1 className="text-center display-4">Artikel</h1>
                </RegeTitle>
                <Container>
                    <Row>
                        <Col md={12}>
                            <Form inline onSubmit={this.handleSearchForm} className={"c-filter justify-content-center"}>

                                <Form.Group>
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
                                        {
                                            this.state.topics.map(item => <option key={item.id} value={item.id}>
                                                {item.name}
                                            </option>)
                                        }
                                    </Form.Control>
                                </Form.Group>

                                <Form.Group>
                                    <Form.Label htmlFor="titleSearch" srOnly>Judul</Form.Label>
                                    <Form.Control
                                        id="titleSearch"
                                        type="text"
                                        name="title"
                                        maxLength={100}
                                        minLength={3}
                                        required={false}
                                        autoComplete="off"
                                        placeholder="Judul Berita yang Anda Cari" />
                                </Form.Group>

                                <Button type="submit" variant="info" size="sm">
                                    Cari
                                </Button>
                            </Form>

                        </Col>
                        <Col md={12}>

                            <div className="row justify-content-center mb-3">
                            {
                                this.state.data.map(item => {
                                    return <div className="col-auto" key={item.id}>
                                        <Card className="h-100" style={{'width':'15rem'}}>
                                            <Image
                                                className="card-img-top"
                                                src={item.thumbnail ?? "/static/img/jakarta.jpg"}
                                                fallbackSrc={"/static/img/jakarta.jpg"}
                                                alt={item.name}/>
                                            <Card.Body>
                                                <Card.Title><h4>{item.title}</h4></Card.Title>
                                                <p>{item.description} &nbsp;
                                                    <Link to={`articles/${item.id}`}>
                                                        <small><b>Read More...</b></small>
                                                    </Link>
                                                </p>
                                            </Card.Body>
                                        </Card>
                                    </div>
                                })
                            }
                            </div>

                        </Col>
                        <Col md={12}>

                            <ReactPaginate
                                previousLabel={'previous'}
                                nextLabel={'next'}
                                breakLabel={'...'}
                                breakClassName={'break-me'}
                                pageCount={ this.state.total / this.state.pagination.perPage }
                                marginPagesDisplayed={2}
                                pageRangeDisplayed={5}
                                onPageChange={this.handlePageChange}
                                containerClassName={'pagination justify-content-center'}
                                subContainerClassName={'pages pagination'}
                                activeClassName={'active'}
                            />

                        </Col>
                    </Row>
                </Container>
            </>
        )
    }
}
