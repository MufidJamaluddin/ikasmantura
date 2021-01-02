import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';

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
    fetched: boolean
}

interface ActionsParamType {
    init: undefined
    done: undefined
}

const ArticleTopicModel: ModelType<StateType, ActionsParamType> = {
    state: {
        data: null,
        fetched: false,
    },
    actions: {
        done: (_, { state }) => {
            return { ...state, fetched: true }
        },
        init: async (_, {state, actions}) => {
            if(state.fetched) {
                return state
            }
            await actions.done()

            let newState = {...state ?? {}}
            if (newState.data === null) {
                newState.data = await getTopics()
                newState.fetched = newState.data === null
            }
            return newState
        }
    }
}

export default ArticleTopicModel
