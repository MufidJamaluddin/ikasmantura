import React from "react";
import {RouteComponentProps} from "react-router";
import RegeTitle from "../component/RegeTitle";
import {Button, Col, Container, OverlayTrigger, Row, Tooltip} from "react-bootstrap";
import {Link} from "react-router-dom";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';

export default class HomeView extends React.PureComponent<RouteComponentProps, {topics: Array<any>}>
{
    constructor(props) {
        super(props);
        this.state = {
            topics: []
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
            NotificationManager.error(error.message, error.name);
        })
    }

    componentDidMount()
    {
        try
        {
            this.updateTopics()
        }
        catch (e)
        {
            NotificationManager.error('Koneksi Internet Terputus!', 'Error Koneksi');
        }
    }

    render() {
        return (
            <section className="features-icons bg-light">
                <RegeTitle>
                    <p className="h1 text-center text-white-50 font-weight-bold">Selamat Datang di Portal</p>
                    <div className="divider-custom">
                        <div className="divider-custom-line"/>
                        <div className="divider-custom-icon">
                            <i className="fa fa-medal"/>
                        </div>
                        <div className="divider-custom-line"/>
                    </div>
                    <p className="display-4 text-center text-white-50 font-weight-bold">Ikatan Alumni SMAN Situraja</p>
                </RegeTitle>

                <Container>
                    <Row className="text-center">

                        <Col sm={6} md={4} lg={2}>
                            <OverlayTrigger
                                placement={'bottom'}
                                overlay={
                                    <Tooltip id={`tooltip-events`}>
                                        Kegiatan yang akan dan telah dilaksanakan
                                    </Tooltip>
                                }>
                                {({ ref, ...triggerHandler }) => (
                                    <div {...triggerHandler}
                                        className="features-icons-item mx-auto mb-5 mb-lg-0 mb-lg-2">
                                        <div ref={ref}
                                            className="features-icons-icon d-flex animate__animated animate__bounceOut animate__slower animate__infinite infinite">
                                            <i className="fas fa-calendar-alt m-auto text-danger"/>
                                        </div>
                                        <h3 className="text-success">Kegiatan</h3>
                                        <Link to='/events'>
                                            <Button variant="outline-primary mb-1">
                                                Buka
                                            </Button>
                                        </Link>
                                    </div>
                                )}
                            </OverlayTrigger>
                        </Col>

                        <Col sm={6} md={4} lg={2}>
                            <OverlayTrigger
                                placement={'bottom'}
                                overlay={
                                    <Tooltip id={`tooltip-gallery`}>
                                        Galeri foto acara dan kegiatan
                                    </Tooltip>
                                }>
                                {({ ref, ...triggerHandler }) => (
                                    <div {...triggerHandler}
                                     className="features-icons-item mx-auto mb-5 mb-lg-0 mb-lg-2">
                                        <div ref={ref}
                                            className="features-icons-icon d-flex animate__animated animate__bounceOut animate__slower animate__infinite infinite">
                                            <i className="fas fa-images m-auto text-secondary"/>
                                        </div>
                                        <h3 className="text-success">Galeri</h3>
                                        <Link to='/gallery'>
                                            <Button variant="outline-primary mb-1">
                                                Buka
                                            </Button>
                                        </Link>
                                    </div>
                                )}
                            </OverlayTrigger>
                        </Col>

                        {
                            this.state.topics.map(item => {
                                return (
                                    <Col sm={6} md={4} lg={2} key={item.id}>
                                        <OverlayTrigger
                                            key={item.id}
                                            placement={'bottom'}
                                            overlay={
                                                <Tooltip id={`tooltip-${item.id}`}>
                                                    {item.description}
                                                </Tooltip>
                                            }>
                                            {({ ref, ...triggerHandler }) => (
                                                <div {...triggerHandler}
                                                    className="features-icons-item mx-auto mb-5 mb-lg-0 mb-lg-2">
                                                    <div
                                                        ref={ref}
                                                        className="features-icons-icon d-flex animate__animated animate__bounceOut animate__slower animate__infinite infinite">
                                                        <i className={`${item.icon} m-auto`}/>
                                                    </div>
                                                    <h3 className="text-success">{item.name}</h3>
                                                    <Link to={{
                                                        pathname: '/articles',
                                                        state: {
                                                            topicId: item.id
                                                        }
                                                    }}>
                                                        <Button variant="outline-primary mb-1">
                                                            Buka
                                                        </Button>
                                                    </Link>
                                                </div>
                                            )}
                                        </OverlayTrigger>
                                    </Col>
                                )
                            })
                        }

                    </Row>
                </Container>

            </section>
        )
    }
}
