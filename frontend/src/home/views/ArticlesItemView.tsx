import React, {PureComponent} from "react";
import {RouteComponentProps} from "react-router";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";

import {NotificationManager} from 'react-notifications';

import moment from "moment";
import 'moment/locale/id';

import RegeTitle from "../component/RegeTitle";
import {Card, Col, Container, Row} from "react-bootstrap";
import Image from "../component/Image";
import DOMPurify from "../../utils/Sanitizer";

interface ArticlesItemState {
    data: any
}

export default class ArticlesItemView
    extends PureComponent<RouteComponentProps<{id: string}>, ArticlesItemState>
{
    constructor(props:any) {
        super(props);
        this.state = {
            data: null
        }
        moment.locale('id');
    }

    updateData()
    {
        let dataProvider = DataProviderFactory.getDataProvider()
        let id = this.props.match.params.id

        dataProvider.getOne('articles', { id: id }).then(resp => {
            let data = resp.data

            this.setState({
                data: data
            })
        }, error => {
            NotificationManager.error(error, 'Get Data Error');
        })
    }

    componentDidMount()
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
        let item = this.state.data

        if(item === null) return <div className="c-center-box c-loader"/>;

        let createdAt = moment(item.createdAt).format('LLLL');

        return (
            <>
                <RegeTitle/>
                <Container>
                    <Row>
                        <Col md={{span:10, offset:1}}>
                            <Card>
                                <Image
                                    className="card-img-top"
                                    src={item.image ?? "/static/img/jakarta.jpg"}
                                    fallbackSrc={"/static/img/jakarta.jpg"}
                                    alt={item.name}/>
                                <Card.Title>
                                    <h1 className="text-center">{item.title}</h1>
                                </Card.Title>
                                <Card.Text className="lead text-center">
                                    <b>Oleh: {item.createdByName ?? 'Kakak Anonim'}</b> &nbsp;
                                    <small>Pada: {createdAt}</small>
                                </Card.Text>
                                <Card.Body>
                                    <div dangerouslySetInnerHTML={{ __html: DOMPurify.sanitize(item.body) }} />
                                </Card.Body>
                            </Card>
                        </Col>
                    </Row>
                </Container>
            </>
        )
    }

}
