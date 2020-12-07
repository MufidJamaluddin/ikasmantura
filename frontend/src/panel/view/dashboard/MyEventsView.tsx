import React, {PureComponent, RefObject} from "react";

import {RouteComponentProps} from "react-router";
import moment from "moment";

import FullCalendar from "@fullcalendar/react"
import dayGridPlugin from '@fullcalendar/daygrid'
import listGridPlugin from '@fullcalendar/list'
import {EventClickArg} from '@fullcalendar/core'

import idLocale from '@fullcalendar/core/locales/id'

import DataProviderFactory from "../../../dataprovider/DataProviderFactory";

import {showNotification} from 'react-admin';

import Modal from '@material-ui/core/Modal';
import Button from '@material-ui/core/Button';

import DownloadIcon from '@material-ui/icons/GetApp';
import DeleteIcon from '@material-ui/icons/Delete';
import {download} from "../../../utils/DownloadUtil";

interface MyEventsViewItemState {
    isOpen: boolean
    data?: any
}

class MyEventsViewItem extends PureComponent<{}, MyEventsViewItemState>
{
    constructor(props:any) {
        super(props);
        this.state = {
            isOpen: false,
            data: null
        }
        this.onDownload = this.onDownload.bind(this)
    }

    openModal(data: any) {
        this.setState({ isOpen: true, data: data })
    }

    onDownload(e) {
        try{
            e.preventDefault()
        }
        catch (ex){

        }

        let id = this.state.data?.id
        if(id) {
            download({
                action: 'POST',
                path: '/api/v1/eventsdownload/' + id,
            })
        }

        return false
    }

    render() {
        let state = this.state
        return (
            <Modal
                open={state.isOpen}
                onClose={() => { this.setState({ isOpen: false}); return false; }}
                aria-labelledby="simple-modal-title"
                aria-describedby="simple-modal-description"
                className={'d-modal'}
            >
                <div className={'container c-text-center'}>
                    <h3>{state.data?.title}</h3>
                    <b>
                        {
                            state.data?.start && moment(state.data?.start).format('MMM Do YYYY h:mm')
                        }
                        &nbsp;s.d.&nbsp;
                        {
                            state.data?.end && moment(state.data?.end).format('MMM Do YYYY h:mm')
                        }
                    </b>
                    <img src={state.data?.image} alt={state.data?.title} className={'d-modal-image'}/>
                    <p>{state.data?.description}</p>

                    <Button onClick={this.onDownload} type={'button'} variant="contained" color="primary">
                        <DownloadIcon/> Download Tiket
                    </Button>
                    &nbsp;
                    <Button type={'button'} variant="contained" color="secondary">
                        <DeleteIcon /> Delete / Unregister
                    </Button>
                </div>
            </Modal>
        )
    }
}

export default class MyEventsView extends PureComponent<RouteComponentProps<any>>
{
    onDateRangeChanged(
        fetchInfo: any,
        successCallback,
        failureCallback
    )
    {
        let start = moment(fetchInfo.start.valueOf())
            .format('YYYY-MM-DD')

        let end = moment(fetchInfo.end.valueOf())
            .format('YYYY-MM-DD')

        const dataProvider = DataProviderFactory.getDataProvider()

        let params:any = {
            pagination: {
                page: 1,
                perPage: 1000,
            },
            sort: {
                field: '',
                order: '',
            },
            filter: {
                start_gte: start,
                start_lte: end,
                myEvent: true,
            },
        }

        dataProvider.getList('events', params).then(value => {
            successCallback(value.data);
        }, error => {
            showNotification(error, 'error')
        })

    }

    private eventItemRef: RefObject<MyEventsViewItem> = React.createRef();

    render() {
        const eventItemReference = this.eventItemRef
        return (
            <>
                <MyEventsViewItem ref={this.eventItemRef}  />
                <FullCalendar
                    height={"auto"}
                    plugins={[ dayGridPlugin, listGridPlugin ]}
                    locale={idLocale}
                    weekends={true}
                    themeSystem={"standard"}
                    headerToolbar={{
                        left: 'prev,next today',
                        center: 'title',
                        right: 'dayGridMonth,listWeek'
                    }}
                    events={this.onDateRangeChanged}
                    editable={false}
                    eventClick={function (info: EventClickArg){
                        let data = {
                            id: info.event.id,
                            start: info.event.start,
                            end: info.event.end,
                            title: info.event.title,
                            image: info.event.extendedProps?.image,
                            description: info.event.extendedProps?.description,
                        }

                        eventItemReference.current.openModal(data)
                    }}
                />
            </>
        )
    }

}
