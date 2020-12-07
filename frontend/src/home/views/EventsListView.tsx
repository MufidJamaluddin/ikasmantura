import React, {PureComponent} from "react";
import RegeTitle from "../component/RegeTitle";

import {NotificationManager} from 'react-notifications';
import moment from "moment";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {Pagination, Sort} from "ra-core/src/types";
import {Button, Card, Col, Container, Form, Row} from "react-bootstrap";
import Image from "../component/Image";
import {Link} from "react-router-dom";

import ReactPaginate from 'react-paginate';

interface EventData {id: number, title: string, description: string, image: string, thumbnail: string}

interface EventsListViewState {
    currentDate: string
    data: Array<EventData>
    total: number;
    pagination: Pagination;
    sort: Sort;
    filter: { start_gte?: string, end_lte?: string, title?: string };
    loading: boolean
}

export default class EventsListView extends PureComponent<any, EventsListViewState>
{
    constructor(props:any)
    {
        super(props);

        let initialDate = moment().format('YYYY-MM-DD')

        this.state = {
            currentDate: initialDate,
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
            data: [],
            total: 0,
            loading: true
        }

        this.handlePageChange = this.handlePageChange.bind(this)
        this.handleSearchForm = this.handleSearchForm.bind(this)
    }

    updateData()
    {
        if(this.state.loading)
        {
            const dataProvider = DataProviderFactory.getDataProvider()

            dataProvider.getList('events', this.state).then(resp => {
                this.setState(oldState => {
                    let newState = {
                        data: resp.data as Array<EventData>,
                        loading: false,
                        total: resp.total,
                    }
                    return {...oldState, ...newState}
                })
            }, error => {
                NotificationManager.error(error, 'Get Data Error')
                this.setState(state => {
                    let newState = {loading: false}
                    return {...state, ...newState}
                })
            })
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
        let start = formData.get('start_gte')

        if(start)
        {
            start = moment(start as string).format('YYYY-MM-DD')
        }

        console.log(start)

        this.setState((oldState) => {
            return {
                ...oldState,
                loading: true,
                filter: {
                    title: title as string,
                    start_gte: start as string
                },
                pagination: {
                    page: 1,
                    perPage: oldState.pagination.perPage
                }
            }
        })
    }

    componentDidMount()
    {
        this.updateData()
    }

    componentDidUpdate(prevProps: Readonly<any>, prevState: Readonly<EventsListViewState>, snapshot?: any)
    {
        this.updateData()
    }

    render()
    {
        let initialDate;
        let startGte = this.state.filter?.start_gte;

        if(this.props.location?.state?.date)
        {
            initialDate = moment(this.props.location?.state?.date).format('YYYY-MM-DD')
        }
        else
        {
            initialDate = moment().format('YYYY-MM-DD')
        }

        return (
            <>
                <RegeTitle>
                    <h1 className="text-center display-4">Daftar Kegiatan</h1>
                </RegeTitle>
                <Container>
                    <Row>
                        <Col md={{span:8, offset:2}}>
                            <div className="fa-pull-right">
                                <Link
                                    to={{
                                        pathname: '/events',
                                        state: {
                                            view: 0,
                                            date: startGte ?? initialDate
                                        }
                                    }}
                                >
                                    <Button className="btn-kegiatan">Bulan</Button>
                                </Link>
                                <Link
                                    to={{
                                        pathname: '/events',
                                        state: {
                                            view: 1,
                                            date: startGte ?? initialDate
                                        }
                                    }}
                                >
                                    <Button className="btn-kegiatan">Agenda</Button>
                                </Link>
                                <Button className="btn-kegiatan-active">Daftar</Button>
                            </div>
                        </Col>
                        <Col md={12}>
                            <Form inline onSubmit={this.handleSearchForm}
                                  className={"c-filter justify-content-center"}>

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

                                <Form.Group>
                                    <Form.Label htmlFor="formSearch" srOnly>Mulai Acara</Form.Label>
                                    <Form.Control
                                        id="formSearch"
                                        type="date"
                                        name="start_gte"
                                        required={false}
                                        autoComplete="off"
                                        placeholder="Mulai Acara"
                                        defaultValue={this.props.initialDate}
                                    />
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
                                                    alt={item.title}/>
                                                <Card.Body>
                                                    <Card.Title><h4>{item.title}</h4></Card.Title>
                                                    <p>{item.description} &nbsp;
                                                        <Link to={`events/${item.id}`}>
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
        );
    }
}
