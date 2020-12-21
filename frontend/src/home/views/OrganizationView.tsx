import React, {PureComponent} from "react";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {PaginationPayload, SortPayload} from "ra-core/src/types";
import {GetListParams} from 'ra-core'

import {NotificationManager} from 'react-notifications';
import RegeTitle from "../component/RegeTitle";
import {Card, Container, Alert} from "react-bootstrap";
import Image from "../component/Image";

interface OrganizationViewState {
    total: number;
    pagination: PaginationPayload;
    sort: SortPayload;
    filter: any;
    data: Array<any>;
}

export default class OrganizationView
    extends PureComponent<{},OrganizationViewState>
{
    constructor(props:any) {
        super(props);
        this.state = {
            pagination: {
                page: 1,
                perPage: 1000,
            },
            sort: {
                field: 'id',
                order: 'ASC'
            },
            filter: {

            },
            data: [],
            total: 0
        }
    }

    updateOrganizationData()
    {
        let dataProvider = DataProviderFactory.getDataProvider()

        dataProvider.getList("departments", this.state as GetListParams).then(resp => {
            if(resp.total === 0) {
                NotificationManager.warning('Tidak ada data');
            }
            this.setState(state => {
                let newState = {
                    data: resp.data,
                    total: resp.total,
                }
                return {...state, ...newState}
            })
        }, error => {
            NotificationManager.error(error.message, error.name);

            this.setState(state => {
                let newState = {loading: false}
                return {...state, ...newState}
            })
        })
    }

    componentDidMount()
    {
        try
        {
            this.updateOrganizationData()
        }
        catch (e)
        {
            NotificationManager.error('Koneksi Internet Terputus!', 'Error Koneksi');
        }
    }

    render() {
        let data: Array<any> = this.state.data

        let images = [
            "/static/img/blank_avatar.svg",
            "/static/img/ringga.jpg",
            "/static/img/mufid.svg"
        ]

        return (
            <section className="features-icons bg-light">
                <RegeTitle>
                    <h1 className="text-center display-4">
                        Struktur Organisasi
                    </h1>
                </RegeTitle>
                <Container>
                    <div className="row justify-content-center mb-3">
                        {
                            data.length === 0 && (
                                <Alert variant="warning">
                                    <strong>Belum ada kepengurusan IKA SMAN Situraja</strong>
                                </Alert>
                            )
                        }
                        {
                            data.map((item, key) => {
                                return <div className="col-auto mb-2" key={item.id}>
                                    <Card style={{'width':'12rem'}} className="h-100">
                                        <Image
                                            className="card-img-top"
                                            src={item.image ?? process.env.PUBLIC_URL + images[key % 3]}
                                            fallbackSrc={process.env.PUBLIC_URL + images[0]}
                                            alt={item.name}/>
                                        <Card.Body>
                                            <p className={"lead"}>{item.userFullname ?? '<Kak IKA>'}</p>
                                            <b>{item.name}</b>
                                        </Card.Body>
                                    </Card>
                                </div>
                            })
                        }
                    </div>
                </Container>
            </section>
        )
    }
}
