import * as React from "react";
import AppBar from './AppBar'
import Typography from '@material-ui/core/Typography';
import {makeStyles} from '@material-ui/core/styles';
import MyUserMenu from "./MyUserMenu";

import {useSelector} from 'react-redux';

import {useMediaQuery} from '@material-ui/core';

import clsx from 'clsx';

const drawerWidth = 240;

const useStyles = makeStyles(theme => ({
    title: {
        flex: 1,
        textOverflow: 'ellipsis',
        whiteSpace: 'nowrap',
        overflow: 'hidden',
        textAlign: 'center'
    },
    appBar: {
        zIndex: theme.zIndex.drawer + 1,
        transition: theme.transitions.create(['width', 'margin'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
        }),
    },
    appBarShift: {
        marginLeft: drawerWidth,
        width: `calc(100% - ${drawerWidth}px)`,
        transition: theme.transitions.create(['width', 'margin'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.enteringScreen,
        }),
    },
    appBarUnShift: {
        marginLeft: theme.spacing(7) + 1,
        width: `calc(100% - ${theme.spacing(7) + 1}px)`,
        transition: theme.transitions.create(['width', 'margin'], {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.enteringScreen,
        }),
    },
}));

const MyAppBar = props => {
    const isSmall = useMediaQuery((theme:any) => theme.breakpoints.down('sm'));
    const open = useSelector(state => state.admin.ui.sidebarOpen);
    const classes = useStyles();

    return (
        <AppBar {...props}
            userMenu={<MyUserMenu/>}
            position="fixed"
            className={isSmall ? null : clsx(
                classes.appBar,
                open ? classes.appBarShift : classes.appBarUnShift,
            )}
        >
            <Typography
                variant="h6"
                color="inherit"
                className={classes.title}
                id="react-admin-title"
            />
        </AppBar>
    );
};

export default MyAppBar;
