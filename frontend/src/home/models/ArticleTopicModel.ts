import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';
import InMemoryCache from "../../dataprovider/InMemoryCache";

export interface ArticleTopicItem {
    id: number|string
    name: string
    icon: string
    description: string
}

async function getTopics () {
    let dataProvider = DataProviderFactory.getDataProvider()

    return await dataProvider.getList("article_topics", {
        pagination: {
            page: 1,
            perPage: 100,
        },
        sort: {
            field: 'id',
            order: 'ASC'
        },
        filter: {
        },
    }).then(resp => {
        return resp.data as Array<ArticleTopicItem>
    }, error => {
        NotificationManager.error(error.message, `Error Koneksi: ${error.name}`)
        return null
    })
}

interface StateType {
    data: Array<ArticleTopicItem>|null
}

interface ActionsParamType {
    init: undefined
}

const ArticleTopicModel: ModelType<StateType, ActionsParamType> = {
    state: {
        data: null,
    },
    actions: {
        init: async (_, {state, actions}) => {
            let cacheKey = 'fetched_topics'
            if(InMemoryCache.getCache(cacheKey)) {
                return state
            }
            InMemoryCache.setCache(cacheKey, true)

            let newState = {...state ?? {}}
            if (newState.data === null) {
                newState.data = await getTopics()
                if(newState.data === null) {
                    InMemoryCache.setCache(cacheKey, false)
                }
            }
            return newState
        }
    }
}

export default ArticleTopicModel
