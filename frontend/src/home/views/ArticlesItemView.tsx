import React, { PureComponent } from "react";
import {RouteComponentProps} from "react-router";
import DataProviderFactory from "../../dataprovider/DataProviderFactory";

import {NotificationManager} from 'react-notifications';

import moment from "moment";

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

        console.log(this.props.match.params.id)
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
        this.updateData()
    }

    render()
    {
        let item = this.state.data

        if(item === null) return <div className="c-center-box c-loader"/>;

        let createdAt = moment(item.createdAt).format('DD MMMM YYYY');

        return (
            <>
                <img className={"c-img-full"}
                     src={item.image ?? "/static/img/jakarta.jpg"}
                     alt="Avatar"/>
                <div className="c-container-pad c-container-centered">
                    <h4><b>{item.title}</b></h4>
                    <b>{item.createdByName}</b> &nbsp; <small>{item.createdAt}</small>
                    <p>{item.body}</p>
                </div>
            </>
        )
    }

}