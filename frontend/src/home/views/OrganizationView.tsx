import React, {PureComponent} from "react";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {Pagination, Sort} from "ra-core/src/types";
import {GetListParams} from 'ra-core'

import {NotificationManager} from 'react-notifications';
import RegeTitle from "../component/RegeTitle";
import {Card, Container} from "react-bootstrap";
import Image from "../component/Image";

interface OrganizationViewState {
    total: number;
    pagination: Pagination;
    sort: Sort;
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

    componentDidMount() {

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
            NotificationManager.error(error, 'Get Data Error');

            this.setState(state => {
                let newState = {loading: false}
                return {...state, ...newState}
            })
        })
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
