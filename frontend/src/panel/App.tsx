import * as React from "react";
import {Admin, Resource} from 'react-admin';

import DashboardView from "./view/DashboardView";
import authProvider from "../dataprovider/authProvider";
import LoginView from "./view/LoginView";

import {createMuiTheme} from '@material-ui/core/styles';
import MyLayout from "./component/MyLayout";

import DataProviderFactory from "../dataprovider/DataProviderFactory";

import * as AboutView from "./view/AboutView";
import * as ClassroomView from "./view/ClassroomView";
import * as DepartmentView from "./view/DepartmentView";
import * as UserView from "./view/UserView";
import * as TempUserView from "./view/TempUserView";
import * as ArticleView from "./view/ArticleView";
import * as ArticleTopicView from "./view/ArticleTopicView";
import * as EventView from "./view/EventView";
import * as AlbumView from "./view/AlbumView";
import * as GalleryView from "./view/GalleryView";

import AboutIcon from '@material-ui/icons/Info';
import DepartmentIcon from '@material-ui/icons/Work'
import UserIcon from '@material-ui/icons/Group'
import ArticleIcon from '@material-ui/icons/Book'
import ArticleTopicIcon from '@material-ui/icons/LibraryBooks';
import EventViewIcon from '@material-ui/icons/EventNote';
import AlbumViewIcon from '@material-ui/icons/PhotoLibrary';
import GalleryIcon from '@material-ui/icons/Panorama';

const history = require("history").createBrowserHistory({ basename: 'panel' })
const dataProvider = DataProviderFactory.getDataProvider()

// @ts-ignore
const theme = createMuiTheme({
    palette: {
        primary: {
            main: 'rgb(128, 100, 161)'
        },
        secondary: {
            main: '#3498db'
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
           loginPage={(ip) => <LoginView
               parenthistory={props.history} theme={theme} {...ip}/>
           }
           authProvider={authProvider}
           history={history}
           dataProvider={dataProvider}
           disableTelemetry
    >
        { permissions => [
            <Resource name="about"
                      options={{"label": "Tentang Kami", "hidden": true}}
                      edit={permissions === 'admin' ? AboutView.AboutEdit : null}
                      show={AboutView.AboutShow}
                      icon={AboutIcon}
            />,

            <Resource name="classrooms"
                options={{"label": "Kelas"}}
                list={ClassroomView.ClassroomList}
                edit={permissions === 'admin' ? ClassroomView.ClassroomEdit : null}
                create={permissions === 'admin' ? ClassroomView.ClassroomCreate: null}
                icon={AboutIcon}
            />,

            <Resource name="departments"
                options={{"label": "Departemen", "divider": true}}
                list={DepartmentView.DepartmentList}
                show={DepartmentView.DepartmentShow}
                edit={permissions === 'admin' ? DepartmentView.DepartmentEdit : null}
                create={permissions === 'admin' ? DepartmentView.DepartmentCreate : null}
                icon={DepartmentIcon}
            />,


            <Resource name="users"
                options={{"label": "Anggota", "hidden": permissions !== 'admin'}}
                list={permissions === 'admin' ? UserView.UserList : null}
                show={permissions === 'admin' ? UserView.UserView : null}
                edit={permissions === 'admin' ? UserView.UserEdit : null}
                create={permissions === 'admin' ? UserView.UserCreate : null}
                icon={UserIcon}
            />,

            <Resource name="temp_users"
                options={{"label": "Registrasi Anggota", "hidden": permissions !== 'admin' }}
                list={permissions === 'admin' ? TempUserView.TempUserList : null}
                show={permissions === 'admin' ? TempUserView.TempUserView : null}
                edit={permissions === 'admin' ? TempUserView.TempUserEdit : null}
                create={permissions === 'admin' ? TempUserView.TempUserCreate : null}
                icon={UserIcon}
            />,

            <Resource name="article_topics"
                options={{"label": "Topik Artikel", "divider": true}}
                list={ArticleTopicView.TopicList}
                edit={permissions === 'admin' ? ArticleTopicView.TopicEdit : null}
                create={permissions === 'admin' ? ArticleTopicView.TopicCreate : null}
                icon={ArticleTopicIcon}
            />,

            <Resource name="articles"
                options={{"label": "Artikel", "divider": false}}
                list={ArticleView.PostList}
                show={ArticleView.PostShow}
                edit={
                  permissions === 'admin' || permissions === 'member' ?
                      ArticleView.PostEdit : null}
                create={
                  permissions === 'admin' || permissions === 'member' ?
                      ArticleView.PostCreate : null}
                icon={ArticleIcon}
            />,

            <Resource name="events"
                options={{"label": "Agenda", "divider": true}}
                list={EventView.EventList}
                show={EventView.EventView}
                edit={permissions === 'admin' ? EventView.EventEdit : null}
                create={permissions === 'admin' ? EventView.EventCreate : null}
                icon={EventViewIcon}
            />,

            <Resource name="albums"
                options={{"label": "Album", "divider": true}}
                list={AlbumView.AlbumList}
                show={AlbumView.AlbumView}
                edit={permissions === 'admin' ? AlbumView.AlbumEdit : null}
                create={permissions === 'admin' ? AlbumView.AlbumCreate : null}
                icon={AlbumViewIcon}
            />,

            <Resource name="photos"
                options={{"label": "Gallery",}}
                list={GalleryView.GalleryList}
                show={GalleryView.GalleryView}
                edit={permissions === 'admin' ? GalleryView.GalleryEdit : null}
                create={permissions === 'admin' ? GalleryView.GalleryCreate : null}
                icon={GalleryIcon}
            />,

            ]
        }
    </Admin>
)

export default AdminApp;
