/**
 * @author Mufid Jamaluddin
 */
import users from './users.json'
import departments from './departments.json'
import topics from './topics.json'
import articles from './articles.json'
import events from './events.json'
import albums from './albums.json'
import photos from './photos.json'
import about from './about.json'

function getFakeData()
{
    return {
        "about": about,
        "users": users,
        "departments": departments,
        "articles": articles,
        "events_download": [],
        "article_topics": topics,
        "events": events,
        "albums": albums,
        "photos": photos,
    }
}

export default getFakeData
