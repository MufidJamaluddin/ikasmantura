import React from "react";
import {Route, RouteComponentProps, Switch} from "react-router";
import {Link} from "react-router-dom";
import ROUTES from "./routes";
import {getSubroute} from "../utils/RouteUtil";

import {NotificationContainer} from 'react-notifications';

import 'react-notifications/lib/notifications.css';

export default class Template extends React.PureComponent<RouteComponentProps<any>>
{
    render() {
        return (
            <div>
                <ul className="c-top_navigation">
                    {
                        ROUTES.map((item, key) => {
                            if(!item.menu) return null;
                            return <li key={key}>
                                <Link to={getSubroute(this.props.match,item.path)}>
                                    {item.title}
                                </Link>
                            </li>
                        })
                    }
                    <li><Link to={'/panel'}>Ruang Alumni</Link></li>
                </ul>
                <Switch>
                    {
                        ROUTES.map((item, key) => {
                            return <Route
                                key={key}
                                exact={item.exact}
                                path={getSubroute(this.props.match, item.path)}
                                component={item.component}/>
                        })
                    }
                </Switch>
                <NotificationContainer/>
            </div>
        )
    }
}