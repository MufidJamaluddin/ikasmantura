import * as React from "react";
import { Admin, Resource } from 'react-admin';

import DashboardView from "./view/DashboardView";
import authProvider from "./authProvider";
import LoginView from "./view/LoginView";

import { createMuiTheme } from '@material-ui/core/styles';
import MyLayout from "./component/MyLayout";

import DataProviderFactory from "../dataprovider/DataProviderFactory";

import * as AboutView from "./view/AboutView";
import * as DepartmentView from "./view/DepartmentView";
import * as UserView from "./view/UserView";
import * as ArticleView from "./view/ArticleView";
import * as EventView from "./view/EventView";
import * as AlbumView from "./view/AlbumView";
import * as GalleryView from "./view/GalleryView";

import AboutIcon from '@material-ui/icons/Info';
import DepartmentIcon from '@material-ui/icons/Work'
import UserIcon from '@material-ui/icons/Group'
import ArticleIcon from '@material-ui/icons/Book'
import EventViewIcon from '@material-ui/icons/EventNote';
import AlbumViewIcon from '@material-ui/icons/PhotoLibrary';
import GalleryIcon from '@material-ui/icons/Panorama';

const history = require("history").createBrowserHistory({ basename: 'panel' })
const dataProvider = DataProviderFactory.getDataProvider()

// @ts-ignore
const theme = createMuiTheme({
    palette: {
        primary: {
            main: '#4400ff'
        },
        secondary: {
            main: '#4400ff'
        },
        error: {
            main: '#f44336'
        },
        contrastThreshold: 3,
        tonalOffset: 0.2,
    },
    typography: {
        // Use the system font instead of the default Roboto font.
        fontFamily: [
            '-apple-system',
            'BlinkMacSystemFont',
            '"Segoe UI"',
            'Arial',
            'sans-serif',
        ].join(','),
    },
    overrides: {
        MuiButton: { // override the styles of all instances of this component
            root: { // Name of the rule
                color: 'white', // Some CSS
            },
        }
    },
});

const AdminApp = (props: any) => (

    <Admin theme={theme}
           layout={(ip) => <MyLayout parenthistory={props.history} {...ip}/>}
           dashboard={DashboardView}
           loginPage={(ip) => <LoginView parenthistory={props.history} {...ip}/>}
           authProvider={authProvider}
           history={history}
           dataProvider={dataProvider}>

        <Resource name="about"
              options={{ "label": "Tentang Kami", "hidden": true }}
              edit={AboutView.AboutEdit}
              show={AboutView.AboutShow}
              icon={AboutIcon}
        />

        <Resource name="departments"
              options={{ "label": "Departemen", "divider": true }}
              list={DepartmentView.DepartmentList}
              edit={DepartmentView.DepartmentEdit}
              show={DepartmentView.DepartmentShow}
              create={DepartmentView.DepartmentCreate}
              icon={DepartmentIcon}
        />

        <Resource name="users"
              options={{ "label": "Anggota", }}
              list={UserView.UserList}
              edit={UserView.UserEdit}
              create={UserView.UserCreate}
              show={UserView.UserView}
              icon={UserIcon}
        />

        <Resource name="articles"
                  options={{ "label": "Artikel", "divider": true }}
              list={ArticleView.PostList}
              edit={ArticleView.PostEdit}
              create={ArticleView.PostCreate}
              show={ArticleView.PostShow}
              icon={ArticleIcon}
        />

        <Resource name="events"
                  options={{ "label": "Agenda" }}
              list={EventView.EventList}
              edit={EventView.EventEdit}
              create={EventView.EventCreate}
              show={EventView.EventView}
              icon={EventViewIcon}
        />

        <Resource name="albums"
                  options={{ "label": "Album", "divider": "true"}}
              list={AlbumView.AlbumList}
              edit={AlbumView.AlbumEdit}
              create={AlbumView.AlbumCreate}
              show={AlbumView.AlbumView}
              icon={AlbumViewIcon}
        />

        <Resource name="photos"
                  options={{ "label": "Gallery", }}
              list={GalleryView.GalleryList}
              edit={GalleryView.GalleryEdit}
              create={GalleryView.GalleryCreate}
              show={GalleryView.GalleryView}
              icon={GalleryIcon}
        />

    </Admin>
)

export default AdminApp;