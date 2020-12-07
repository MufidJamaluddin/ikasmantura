import * as React from 'react';
import {Fragment} from 'react';
import {useDispatch, useSelector} from 'react-redux';
import {Drawer, useMediaQuery} from '@material-ui/core';
import {getResources, MenuItemLink} from 'react-admin';
import DefaultIcon from '@material-ui/icons/ViewList';
import {toSentenceCase} from "../../utils/toSentenceCase";

import Divider from '@material-ui/core/Divider';

import AppsIcon from '@material-ui/icons/Apps'
import HomeIcon from '@material-ui/icons/Home'
import AboutIcon from '@material-ui/icons/Info';

import {makeStyles} from '@material-ui/core/styles';

import IconButton from '@material-ui/core/IconButton';

import ChevronLeftIcon from '@material-ui/icons/ChevronLeft';
import ChevronRightIcon from '@material-ui/icons/ChevronRight';
import {toggleSidebar} from 'ra-core';

import clsx from 'clsx';

const drawerWidth = 240;

const useStyles = makeStyles((theme) => ({
    drawer: {
        width: drawerWidth,
        flexShrink: 0,
        whiteSpace: 'nowrap',
    },
    drawerOpen: {
        width: drawerWidth,
        transition: theme.transitions.create('width', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.enteringScreen,
        }),
    },
    drawerClose: {
        transition: theme.transitions.create('width', {
            easing: theme.transitions.easing.sharp,
            duration: theme.transitions.duration.leavingScreen,
        }),
        overflowX: 'hidden',
        width: theme.spacing(7) + 1,
    },
    toolbar: {
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'flex-end',
        padding: theme.spacing(0, 1),
        // necessary for content to be below app bar
        ...theme.mixins.toolbar,
    },
}));

const MyMenu = ({ parenthistory, onMenuClick, logout }) =>
{
    const isXSmall = useMediaQuery((theme:any) => theme.breakpoints.down('xs'));
    const open = useSelector(state => state.admin.ui.sidebarOpen);
    const resources = useSelector(getResources);
    const classes = useStyles()
    const dispatch = useDispatch()

    return (
        <Drawer
            variant="permanent"
            className={clsx(classes.drawer, {
                [classes.drawerOpen]: open,
                [classes.drawerClose]: !open,
            })}
            classes={{
                paper: clsx({
                    [classes.drawerOpen]: open,
                    [classes.drawerClose]: !open,
                }),
            }}
        >
            <div className={classes.toolbar}>
                <IconButton onClick={() => dispatch(toggleSidebar())}>
                    { open ? <ChevronLeftIcon /> : <ChevronRightIcon /> }
                </IconButton>
            </div>
            <img src={process.env.PUBLIC_URL + "/static/img/logo_ika2.png"}
                 className={'d-centered-img'} alt={"Logo IKA Smantura"}
            />
            <br/>
            <MenuItemLink
                to="/about/1/show"
                primaryText={'About'}
                leftIcon={<AboutIcon />}
                onClick={onMenuClick}
                sidebarIsOpen={open}
            />
            <Divider variant={'middle'}/>
            <MenuItemLink
                to="/"
                primaryText={'Dashboard'}
                leftIcon={<AppsIcon />}
                onClick={onMenuClick}
                sidebarIsOpen={open}
            />
            {resources.map((resource, key) => {
                if(resource.options && !resource.options.hidden) {
                    return (
                        <Fragment key={key}>
                            {resource.options.divider && <Divider variant={'middle'}/>}
                            <MenuItemLink
                                key={resource.name}
                                to={`/${resource.name}`}
                                primaryText={
                                    (resource.options && resource.options.label) ||
                                    toSentenceCase(resource.name)
                                }
                                leftIcon={
                                    resource.icon ? <resource.icon/> : <DefaultIcon/>
                                }
                                onClick={onMenuClick}
                                sidebarIsOpen={open}
                            />
                        </Fragment>
                    )
                }
                else {
                    return []
                }
            })}
            <Divider variant={'middle'}/>
            <MenuItemLink
                to="/"
                primaryText={'Back to Home'}
                leftIcon={<HomeIcon />}
                onClick={(e) => {
                    e.preventDefault()
                    parenthistory.push('/')
                }}
                sidebarIsOpen={open}
            />
            <Divider variant={'middle'}/>
            {isXSmall && logout}
        </Drawer>
    );
};

export default MyMenu;
