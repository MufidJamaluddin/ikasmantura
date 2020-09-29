import React, {Fragment, PureComponent} from "react";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";

import {NotificationManager} from 'react-notifications';

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
    constructor(props:any) {
        super(props);
        this.state = {
            data: {}
        }
    }

    componentDidMount() {

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

    render() {
        let data = this.state.data
        return (
            <div className={"c-g-banner c-text-center c-p-full-height"}>
                <div>
                    <h2>Tentang Kami</h2>
                    <p>{data?.description ?? ''}</p>
                    <hr/>
                    <h2>Visi</h2>
                    <p>{data?.vision ?? ''}</p>
                    <hr/>
                    <h2>Misi</h2>
                    <p>{data?.mission ?? ''}</p>
                </div>
            </div>
        )
    }
}