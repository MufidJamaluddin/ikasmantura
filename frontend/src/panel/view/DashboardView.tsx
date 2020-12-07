import * as React from "react";
import {Card, CardContent, CardHeader} from '@material-ui/core';
import MyEventsView from "./dashboard/MyEventsView";
import authProvider from "../authProvider";

export default (props:any) => {
    const userData = authProvider.getData()
    return (
    <Card>
        <CardHeader title={"Selamat Datang, "+ userData.name +"."}/>
        <CardContent>
            <div className={'c-text-center'}>
                <h3>Agenda Anda</h3>
            </div>
            <div>
                <MyEventsView {...props}/>
            </div>
        </CardContent>
    </Card>)
};
