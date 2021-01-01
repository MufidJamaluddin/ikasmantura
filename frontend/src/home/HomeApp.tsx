import React from "react";
import Header from "./component/Header";
import Footer from "./component/Footer";

import {Route, Switch} from "react-router";
import ROUTES from "./routes";

import { NotificationContainer } from 'react-notifications';

import 'bootstrap/dist/css/bootstrap.min.css'
import '@fortawesome/fontawesome-free/css/all.min.css'

import './Landing.css'

import 'react-notifications/lib/notifications.css'
import PageTemplate from "./component/PageTemplate";

import { Provider } from "react-model";

export default function HomeApp(){
    return (
        <div className="c-app">
            <Header/>
            <div className="c-main">
                <Provider>
                    <PageTemplate>
                        <Switch>
                            {
                                ROUTES.map((item, key) => <Route
                                    key={key}
                                    path={`/${item.path}`}
                                    exact={item.exact}
                                    component={item.component}/>)
                            }
                        </Switch>
                    </PageTemplate>
                </Provider>
            </div>
            <NotificationContainer/>
            <Footer/>
        </div>
    )
}
