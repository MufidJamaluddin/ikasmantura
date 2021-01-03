import React, {PureComponent} from "react";
import {RouteComponentProps} from "react-router";

import moment from "moment";
import 'moment/locale/id';

import {Button, Card, Col, Row} from "react-bootstrap";
import Image from "../component/Image";
import DOMPurify from "../../utils/Sanitizer";
import { ThemeContext } from "../component/PageTemplate";
import {getEventById, registerEventById} from "../models/EventsModel";

interface EventItemState {
    data: any
}

export default class EventItemView
    extends PureComponent<RouteComponentProps<{id: string}>|any, EventItemState>
{

    constructor(props:any) {
        super(props);
        this.state = {
            data: null
        }
        moment.locale('id');
    }

    static contextType = ThemeContext;

    async componentDidMount() {
        let id = this.props.match.params.id

        let data = await getEventById(id)

        if (data) {
            this.setState({
                data: data
            })
        }

        let title = data?.title ?? 'Kegiatan IKA'
        this.context.setHeader({title: title, showTitle: false})
    }

    async onDaftarClick(eventId: number | string) {
        let success = await registerEventById(eventId)
        if(success) {
            let data = {...this.state?.data}
            data.myEvent = true
            this.setState(state => ({ ...state, data: data }))
        }
        return false
    }

    render()
    {
        let data = this.state.data

        if(this.state.data === null) return <div className="c-center-box c-loader"/>;

        let {
            image, myEvent, id, start, end, description, title, createdByName
        } = data

        return (
            <Row>
                <Col md={{span:10, offset:1}}>
                    <Card>
                        <Image
                            className="card-img-top"
                            src={image}
                            alt={title}/>
                        <Card.Body>
                            {
                                myEvent ?
                                    (<span
                                        className={"c-button info"}>Anda Telah Terdaftar</span>)
                                    :
                                    (
                                        <Button className={"c-button info"}
                                                onClick={() => this.onDaftarClick(id)}>
                                            Daftar Jadi Peserta
                                        </Button>
                                    )
                            }
                        </Card.Body>
                        <Card.Title>
                            <h1 className="text-center">{title}</h1>
                        </Card.Title>
                        <Card.Text className="lead text-center">
                            Diselenggarakan Oleh: {data.organizer}
                            - {createdByName ?? 'Kakak Anonim'}
                        </Card.Text>
                        <Card.Text className="lead text-center">
                            {
                                start && (<>Mulai Acara: {moment(start).format('LLLL')}</>)
                            }
                            <br/>
                            {
                                end && (<>Akhir Acara: {moment(end).format('LLLL')}</>)
                            }
                        </Card.Text>
                        <Card.Body>
                            <div className="text-justify" dangerouslySetInnerHTML={{
                                __html: description ? DOMPurify.sanitize(description) : '' }} />
                        </Card.Body>
                    </Card>
                </Col>
            </Row>
        )
    }

}
