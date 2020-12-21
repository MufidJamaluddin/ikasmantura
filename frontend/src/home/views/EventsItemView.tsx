import React, {PureComponent} from "react";
import {RouteComponentProps} from "react-router";

import moment from "moment";
import 'moment/locale/id';

import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';
import authProvider from "../../dataprovider/authProvider";
import RegeTitle from "../component/RegeTitle";
import {Button, Card, Col, Container, Row} from "react-bootstrap";
import Image from "../component/Image";
import DOMPurify from "../../utils/Sanitizer";

interface EventItemState {
    data: any
}

export default class EventItemView
    extends PureComponent<RouteComponentProps<{id: string}>, EventItemState>
{

    constructor(props:any) {
        super(props);
        this.state = {
            data: null
        }
        moment.locale('id');
    }

    componentDidMount()
    {
        let dataProvider = DataProviderFactory.getDataProvider()
        let id = this.props.match.params.id

        dataProvider.getOne('events', { id: id }).then(resp => {
            let data = resp.data

            this.setState({
                data: data
            })
        }, error => {
            NotificationManager.error(error.message, error.name);
        })
    }

    onDaftarClick(eventId: number|string)
    {
        const id = eventId;

        authProvider.checkAuth().then(item => {

            fetch('/api/v1/eventregister/' + id, {
                method: 'POST',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json',
                   // 'Authentication': 'Bearer ' + localStorage.getItem('token')
                },
            })
            .then(item => {
                NotificationManager.success('Anda telah terdaftar', 'Pendaftaran Sukses');
                this.setState(state => {
                    let newData = { ...state.data, myEvent: true }
                    return { data: newData }
                })
            }).catch(err => {
                NotificationManager.error(err.toString(), 'Pendaftaran Gagal')
            })

        }, err => {
            NotificationManager.error('Anda harus masuk terlebihdahulu', 'Anda Belum Login')
        })

        return false
    }

    render()
    {
        let data = this.state.data

        if(this.state.data === null) return <div className="c-center-box c-loader"/>;

        return (
            <>
                <RegeTitle/>
                <Container>
                    <Row>
                        <Col md={{span:10, offset:1}}>
                            <Card>
                                <Image
                                    className="card-img-top"
                                    src={data.image ?? "/static/img/jakarta.jpg"}
                                    fallbackSrc={"/static/img/jakarta.jpg"}
                                    alt={data.name}/>
                                <Card.Body>
                                    {
                                        data.myEvent ?
                                            (<span
                                                className={"c-button info"}>Anda Telah Terdaftar</span>)
                                            :
                                            (
                                                <Button className={"c-button info"}
                                                        onClick={() => this.onDaftarClick(data.id)}>
                                                    Daftar Jadi Peserta
                                                </Button>
                                            )
                                    }
                                </Card.Body>
                                <Card.Title>
                                    <h1 className="text-center">{data.title}</h1>
                                </Card.Title>
                                <Card.Text className="lead text-center">
                                    Diselenggarakan Oleh: {data.organizer}
                                    - {data.createdByName ?? 'Kakak Anonim'}
                                </Card.Text>
                                <Card.Text className="lead text-center">
                                    Mulai Acara: {moment(data.start).format('LLLL')}
                                    <br/>
                                    Akhir Acara: {moment(data.end).format('LLLL')}
                                </Card.Text>
                                <Card.Body>
                                    <div dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(data.description) }} />
                                </Card.Body>
                            </Card>
                        </Col>
                    </Row>
                </Container>
            </>
        )
    }

}
