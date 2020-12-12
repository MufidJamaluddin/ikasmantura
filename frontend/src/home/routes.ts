import HomeView from "./views/HomeView";
import ArticlesView from "./views/ArticlesView";
import EventsView from "./views/EventsView";
import ArticlesItemView from "./views/ArticlesItemView";
import EventsItemView from "./views/EventsItemView";
import OrganizationView from "./views/OrganizationView";
import AboutView from "./views/AboutView";
import GalleryView from "./views/GalleryView";
import EventsListView from "./views/EventsListView";
import RegisterView from "./views/RegisterView";
import LoginView from "./views/LoginView";

const ROUTES = [
    { title: 'IKA', path: '', exact: true, menu: true, component: HomeView },

    { title: 'Tentang Kami', path: 'about', exact: true, menu: true, component: AboutView },
    { title: 'Organisasi', path: 'organization', exact: true, menu: true, component: OrganizationView },

    { title: 'Articles Item', path: 'articles/:id', exact: true, menu: false, component: ArticlesItemView },
    { title: 'Articles', path: 'articles', exact: true, menu: true, component: ArticlesView },

    { title: 'Event Item', path: 'events/:id', exact: true, menu: false, component: EventsItemView },
    { title: 'Events', path: 'events', exact: true, menu: true, component: EventsView },
    { title: 'Events List', path: 'events_list', exact: true, menu: true, component: EventsListView },

    { title: 'Galleries', path: 'gallery', exact: true, menu: true, component: GalleryView },

    { title: 'Login', path: 'login', exact: true, menu: true, component: LoginView },
    { title: 'Register', path: 'register', exact: true, menu: true, component: RegisterView },
]

export default ROUTES
