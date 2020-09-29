import React, { PureComponent } from "react";
import {RouteComponentProps} from "react-router";
import moment from "moment";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";


import {NotificationManager} from 'react-notifications';
import authProvider from "../../panel/authProvider";

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
            NotificationManager.error(error, 'Get Data Error');
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
                    'Authentication': 'Bearer ' + localStorage.getItem('token')
                },
            })
            .then(item => {

                console.log(item.json())

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
                <img className={"c-img-full"}
                     src={data.image ?? "/static/img/jakarta.jpg"}
                     alt="Avatar"/>
                <div className="c-container-pad c-container-centered">
                    <h4><b>{data.title}</b></h4>
                    <p><b>{data.organizer} | {data.createdByName} </b></p>
                    <p>
                        <small>
                            {moment(data.start).format('DD MMMM YYYY HH:MM')}
                            &nbsp;s.d.&nbsp;
                            {moment(data.end).format('DD MMMM YYYY HH:MM')}
                        </small>
                    </p>
                    &nbsp;
                    {
                        data.myEvent ?
                        (<span
                            className={"c-button info"}>Anda Telah Terdaftar</span>)
                            :
                        (<button type={"button"}
                                 className={"c-button info"}
                                 onClick={() => this.onDaftarClick(data.id)}>Daftar</button>)
                    }
                    <p>
                        {data.description}
                    </p>
                </div>
            </>
        )
    }

}