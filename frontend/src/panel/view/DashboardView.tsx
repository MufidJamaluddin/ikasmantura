import * as React from "react";
import {Card, CardContent, CardHeader} from '@material-ui/core';
import MyEventsView from "./dashboard/MyEventsView";
import {useGetIdentity, usePermissions } from "react-admin";

export default function DashboardView(props:any) {
    const { identity: { id, fullName } = {} } = useGetIdentity();
    const { permissions } = usePermissions();
    return (
    <Card>
        <CardHeader title={"Selamat Datang, "+ fullName +"."}/>
        {
            (permissions === 'admin' || permissions === 'member') ? (
                <CardContent>
                    <div className={'c-text-center'}>
                        <h3>Agenda {fullName}</h3>
                    </div>
                    <div>
                        <MyEventsView {...props}/>
                    </div>
                </CardContent>
            ) : (
                <CardContent>
                    <div className={'c-text-center'}>
                        <h3>Selamat Datang, {fullName}</h3>
                        <p>
                            Akun anda belum diverifikasi oleh pengurus IKA, mohon ditunggu.
                            Bila tak kunjung diverifikasi, silakan hubungi pengurus IKA di info@ikasmansituraja.org.
                        </p>
                    </div>
                </CardContent>
            )
        }
    </Card>)
};
