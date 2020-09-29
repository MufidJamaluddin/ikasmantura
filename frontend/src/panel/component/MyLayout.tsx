import React from "react";

import { Layout } from 'react-admin';
import MyAppBar from './MyAppBar';
import MySidebar from './MySidebar';
import MyNotification from './MyNotification';
import MyMenu from "./MyMenu";

const MyLayout = props => <Layout
    {...props}
    appBar={MyAppBar}
    sidebar={MySidebar}
    menu={(ip) => <MyMenu parenthistory={props.parenthistory} {...ip}/>}
    notification={MyNotification}
/>;

export default MyLayout;