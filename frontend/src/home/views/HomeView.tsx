import React from "react";
import {RouteComponentProps} from "react-router";
import RegeTitle from "../component/RegeTitle";
import {Button, Col, Container, Row} from "react-bootstrap";
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
            NotificationManager.error(error, 'Get Data Error');
        })
    }

    componentDidMount()
    {
        this.updateTopics()
    }

    render() {
        return (
            <section className="features-icons bg-light">
                <RegeTitle>
                    <p className="text-center display-4">
                        Selamat Datang di Portal Ikatan Alumni SMAN Situraja
                    </p>
                </RegeTitle>

                <Container>
                    <Row className="text-center">

                        <Col sm={6} md={4} lg={2}>
                            <div className="features-icons-item mx-auto mb-5 mb-lg-0 mb-lg-2">
                                <div
                                    className="features-icons-icon d-flex animate__animated animate__bounceOut animate__slower animate__infinite infinite">
                                    <i className="fas fa-calendar-alt m-auto text-danger"/>
                                </div>
                                <h3 className="text-success">Kegiatan</h3>
                                <p className="lead mb-0">Kegiatan yang akan dan telah dilaksanakan</p>
                            </div>
                            <Link to='/events'>
                                <Button variant="outline-primary mb-1">
                                    Buka
                                </Button>
                            </Link>
                        </Col>

                        <Col sm={6} md={4} lg={2}>
                            <div className="features-icons-item mx-auto mb-5 mb-lg-0 mb-lg-2">
                                <div
                                    className="features-icons-icon d-flex animate__animated animate__bounceOut animate__slower animate__infinite infinite">
                                    <i className="fas fa-images m-auto text-secondary"/>
                                </div>
                                <h3 className="text-success">Galeri</h3>
                                <p className="lead mb-0">Galeri foto acara dan kegiatan</p>
                            </div>
                            <Link to='/gallery'>
                                <Button variant="outline-primary mb-1">
                                    Buka
                                </Button>
                            </Link>
                        </Col>

                        {
                            this.state.topics.map(item => {
                                return (
                                    <Col sm={6} md={4} lg={2} key={item.id}>
                                        <div className="features-icons-item mx-auto mb-5 mb-lg-0 mb-lg-2">
                                            <div
                                                className="features-icons-icon d-flex animate__animated animate__bounceOut animate__slower animate__infinite infinite">
                                                <i className={`${item.icon} m-auto`}/>
                                            </div>
                                            <h3 className="text-success">{item.name}</h3>
                                            <p className="lead mb-0">{item.description}</p>
                                        </div>
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
