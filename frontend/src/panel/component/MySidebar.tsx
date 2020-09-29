import React from "react";

import { Sidebar } from 'react-admin';
import { makeStyles } from '@material-ui/core/styles';

const useSidebarStyles = makeStyles({
    drawerPaper: {
        backgroundColor: '#f0f0f0',
    },
});

const MySidebar = props => {
    return (
        <Sidebar classes={useSidebarStyles()} {...props} />
    );
};

export default MySidebar