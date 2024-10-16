import React, {PureComponent} from "react";

import {NotificationManager} from 'react-notifications';
import moment from "moment";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {PaginationPayload, SortPayload} from "ra-core/src/types";
import {Button, Card, Col, Form, Row} from "react-bootstrap";
import Image from "../component/Image";
import {Link} from "react-router-dom";

import ReactPaginate from 'react-paginate';
import {ThemeContext} from "../component/PageTemplate";
import {strip_tags} from "../../utils/Security";

interface EventData {id: number, title: string, description: string, image: string, thumbnail: string}

interface EventsListViewState {
    currentDate: string
    data: Array<EventData>
    total: number;
    pagination: PaginationPayload;
    sort: SortPayload;
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

    static contextType = ThemeContext;

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
                NotificationManager.error(error.message, error.name)
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
        try
        {
            this.context.setHeader({ title: 'Daftar Acara', showTitle: true })
            this.updateData()
        }
        catch (e)
        {
            NotificationManager.error('Koneksi Internet Tidak Ada!', 'Error Koneksi');
        }
    }

    componentDidUpdate(prevProps: Readonly<any>, prevState: Readonly<EventsListViewState>, snapshot?: any)
    {
        try
        {
            this.updateData()
        }
        catch (e)
        {
            NotificationManager.error('Koneksi Internet Tidak Ada!', 'Error Koneksi');
        }
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

        let events = this.state.data

        if(!Array.isArray(events)) {
            return <div className="c-center-box c-loader"/>
        }

        return (
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
                <Col md={{span:8, offset:2}}>
                    <br/>
                    <Form onSubmit={this.handleSearchForm}>

                        <Form.Row className="justify-content-center">

                            <Form.Group as={Col} md="5">
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

                            <Form.Group as={Col} md="5">
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

                            <Col md="2">
                                <Button type="submit" variant="info" size="sm">
                                    Cari
                                </Button>
                            </Col>

                        </Form.Row>

                    </Form>

                </Col>
                <Col md={12}>

                    <div className="row justify-content-center mb-3">
                        {
                            events.map(item => {
                                item.description = strip_tags(item.description ?? '').slice(0, 50);
                                return <div className="col-auto" key={item.id}>
                                    <Card className="h-100" style={{'width':'15rem'}}>
                                        <Image
                                            className="card-img-top"
                                            src={item.thumbnail}
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
        );
    }
}
