import React, {PureComponent} from "react";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {Pagination, Sort} from "ra-core/src/types";
import {GetListParams} from 'ra-core'

import {NotificationManager} from 'react-notifications';

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
            if(resp.total == 0) {
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
            <div className={"c-g-banner c-p-full-height c-text-center"}>
                <h1 className={"lead"}>Struktur Organisasi</h1>
                <div className={'c-center-wrapper'}>
                    {
                        data.map((item, key) => {
                            return <div key={item.id} className="c-card c-card-large">
                                <img className={"c-img-full"}
                                     src={item.image ??
                                        process.env.PUBLIC_URL + images[key % 3]}
                                     alt="Avatar"/>
                                <div className="c-container">
                                    <p className={"lead"}>{item.userFullname ?? '-'}</p>
                                    <b>{item.name}</b>
                                </div>
                            </div>
                        })
                    }
                </div>
            </div>
        )
    }
}