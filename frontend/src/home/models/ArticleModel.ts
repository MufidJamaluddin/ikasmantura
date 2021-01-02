import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';
import {GetListParams} from "ra-core";

export interface ArticleItem {
    id: number|string
    title: string
    image: string
    body?: string
    description?: string
    createdByName: string
    createdAt: string
}

async function getArticles({ page, perPage, filter }) {
    let dataProvider = DataProviderFactory.getDataProvider()

    let cFilter:any = {}

    if(filter.title) {
        let title = filter.title.trim()
        if(title) {
            cFilter.title = title
        }
    }
    if(filter.topicId) {
        cFilter.topicId = filter.topicId
    }

    let request: GetListParams = {
        pagination: {
            page: page,
            perPage: perPage,
        },
        sort: {
            field: 'id',
            order: 'DESC'
        },
        filter: cFilter,
    }

    return await dataProvider.getList("articles", request).then(resp => {
        return {
            data: resp.data as Array<ArticleItem>,
            total: resp.total,
        }
    }, error => {
        NotificationManager.error(error.message, `Error Koneksi: ${error.name}`)
        return null
    })
}

async function getArticleById(id: number | string) {
    let dataProvider = DataProviderFactory.getDataProvider()

    return await dataProvider.getOne('articles', {id: id}).then(resp => {
        return resp.data as ArticleItem
    }, error => {
        NotificationManager.error(error.message, `Error Koneksi: ${error.name}`)
        return null
    })
}

interface StateType {
    data: Array<ArticleItem>|null
    total: number
    selected: ArticleItem|null
}

interface ActionsParamType {
    getArticles: { page:number, perPage:number, filter: any }
    getArticleById: number|string
    reset: undefined
}

const ArticleModel: ModelType<StateType|null, ActionsParamType> = {
    state: null,
    actions: {
        getArticles: async ({page, perPage, filter}, {state, actions}) => {
            let data = await getArticles({page, perPage, filter})
            if(data) {
                return { data: data.data ?? [], total: data.total ?? 0 }
            }
            return state
        },
        getArticleById: async (id: number|string, { state, actions }) => {
            let data = await getArticleById(id)
            if (data) {
                return { selected: data }
            }
            return state
        },
        reset: () => {
            return null
        }
    }
}

export default ArticleModel
