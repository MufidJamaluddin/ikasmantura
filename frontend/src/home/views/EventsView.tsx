import React, {RefObject} from "react";

import {RouteComponentProps} from "react-router";
import moment from "moment";

import FullCalendar from "@fullcalendar/react"
import dayGridPlugin from '@fullcalendar/daygrid'
import listGridPlugin from '@fullcalendar/list'

import idLocale from '@fullcalendar/core/locales/id'

import {Col, Row} from "react-bootstrap";
import {ThemeContext} from "../component/PageTemplate";
import {initEvents} from "../models/EventsModel";

export default class EventsView extends React.PureComponent<
    RouteComponentProps<unknown, unknown, {view?: string|number, date: Date|string}>|any>
{
    static contextType = ThemeContext;

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

        initEvents({ start, end, successCallback, failureCallback })
    }

    getHeader() {
        return <h1 className="text-center display-4"> Kegiatan <span ref={this.titleRef}>IKA</span></h1>
    }

    componentDidMount() {
        this.context.setHeader({ header: this.getHeader(), title: 'Daftar Kegiatan', showTitle: false })
    }

    private titleRef: RefObject<HTMLSpanElement> = React.createRef();
    private calendarRef: RefObject<FullCalendar> = React.createRef();

    render()
    {
        const props = this.props;
        const titleRef = this.titleRef;
        const calendarRef = this.calendarRef;

        let initialViewCode = this.props.location?.state?.view ?? 0
        let initialView = initialViewCode === 0 ? 'dayGridMonth' : 'listWeek';
        let initialDate: Date;

        if(this.props.location?.state?.date)
        {
            initialDate = moment(this.props.location?.state?.date).toDate();
        }
        else
        {
            initialDate = new Date();
        }

        return (
            <Row className="padding-cont">
                <Col md={{ span: 8, offset: 2 }}>

                    <div className={"container"}>
                        <FullCalendar
                            ref={this.calendarRef}
                            height={"auto"}
                            plugins={[ dayGridPlugin, listGridPlugin ]}
                            locale={idLocale}
                            initialView={initialView}
                            initialDate={initialDate}
                            weekends={true}
                            themeSystem={"standard"}
                            customButtons={{
                                ForwardButton: {
                                    icon: 'chevron-right',
                                    click: function (params) {
                                        let date = moment(initialDate)
                                            .add(1, 'month')
                                            .format('YYYY-MM-DD')

                                        props.history.push('events', {
                                            view: initialViewCode,
                                            date: date
                                        });

                                        try
                                        {
                                            if(calendarRef.current) {
                                                let api = calendarRef.current.getApi();
                                                api.next()

                                                if(titleRef.current) {
                                                    titleRef.current.innerText = api.view.title
                                                }
                                            }
                                        }
                                        catch (e)
                                        {

                                        }
                                    }
                                },
                                BackwardButton: {
                                    icon: 'chevron-left',
                                    click: function () {
                                        let date = moment(initialDate)
                                            .subtract(1, 'month')
                                            .format('YYYY-MM-DD')

                                        props.history.push('events', {
                                            view: initialViewCode,
                                            date: date
                                        });

                                        try
                                        {
                                            let api = calendarRef.current.getApi()
                                            api.prev()

                                            titleRef.current.innerText = api.view.title
                                        }
                                        catch (e)
                                        {

                                        }
                                    }
                                },
                                ListButton: {
                                    text: 'Daftar',
                                    click: function (){
                                        let date = moment(initialDate)
                                            .subtract(1, 'month')
                                            .format('YYYY-MM-DD')

                                        props.history.push('events_list', {
                                            date: date
                                        });
                                    }
                                }
                            }}
                            headerToolbar={{
                                left: 'BackwardButton,ForwardButton today',
                                right: 'dayGridMonth,listWeek,ListButton'
                            }}
                            events={this.onDateRangeChanged}
                            editable={false}
                            eventClick={function (info:any){
                                let id = info.event.id;
                                props.history.push(`events/${id}`)
                            }}
                            eventDidMount={function (params) {
                                let isMonth = params.view.type === 'dayGridMonth' ? 0 : 1
                                let date = moment(params.view.currentStart).format('YYYY-MM-DD')

                                try
                                {
                                    titleRef.current.innerText = params.view.title
                                }
                                catch (e)
                                {

                                }

                                props.history.replace('events', {
                                    view: isMonth,
                                    date: date
                                });
                            }}
                        />
                    </div>

                </Col>
            </Row>
        )
    }

}
