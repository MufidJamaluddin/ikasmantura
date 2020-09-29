import * as React from "react";
import Loadable from 'react-loadable';

const AdminApp = Loadable({
    loader: () => import('./App'),
    loading() {
        return <div className="c-center-box c-loader"/>
    }
});

export default AdminApp;