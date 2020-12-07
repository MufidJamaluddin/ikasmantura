import React, {PureComponent} from "react";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";

import {NotificationManager} from 'react-notifications';
import {Col, Container, Row} from "react-bootstrap";
import RegeTitle from "../component/RegeTitle";

interface AboutItem {
    id: number|string
    description: string
    vision: string
    mission: string
}

interface AboutViewState {
    data: AboutItem|any
}

export default class AboutView extends PureComponent<{}, AboutViewState>
{
    constructor(props:any)
    {
        super(props);
        this.state = {
            data: {}
        }
    }

    componentDidMount()
    {
        let dataProvider = DataProviderFactory.getDataProvider()

        dataProvider.getOne("about", { id: 1 }).then(resp => {
            this.setState({ data: resp.data as AboutItem })
        }, error => {
            NotificationManager.error(error, 'Get Data Error');

            this.setState(state => {
                let newState = {loading: false}
                return {...state, ...newState}
            })
        })
    }

    render()
    {
        let data = this.state.data
        return (
            <>
                <RegeTitle>
                    <h1 className="text-center display-4">Tentang Kami</h1>
                </RegeTitle>
                <Container>
                    <Row className="padding-cont">
                        <Col sm={12}>
                            <h3 className="text-center">Deskripsi</h3>
                            <p>{data?.description ?? ''}</p>
                        </Col>
                        <hr/>
                        <Col sm={12}>
                            <h3 className="text-center">Visi</h3>
                            <p>{data?.vision ?? ''}</p>
                        </Col>
                        <hr/>
                        <Col sm={12}>
                            <h3 className="text-center">Misi</h3>
                            <p>{data?.mission ?? ''}</p>
                        </Col>
                    </Row>
                </Container>
            </>
        )
    }
}
