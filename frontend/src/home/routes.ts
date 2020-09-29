import HomeView from "./views/HomeView";
import ArticlesView from "./views/ArticlesView";
import EventsView from "./views/EventsView";
import ArticlesItemView from "./views/ArticlesItemView";
import EventsItemView from "./views/EventsItemView";
import OrganizationView from "./views/OrganizationView";
import AboutView from "./views/AboutView";

const ROUTES = [
    { title: 'IKA', path: '', exact: true, menu: true, component: HomeView },

    { title: 'About', path: 'about', exact: true, menu: true, component: AboutView },
    { title: 'Organization', path: 'organization', exact: true, menu: true, component: OrganizationView },

    { title: 'Articles Item', path: 'news/:id', exact: true, menu: false, component: ArticlesItemView },
    { title: 'Articles', path: 'news', exact: true, menu: true, component: ArticlesView },

    { title: 'Event Item', path: 'events/:id', exact: true, menu: false, component: EventsItemView },
    { title: 'Events', path: 'events', exact: true, menu: true, component: EventsView },
]

export default ROUTES