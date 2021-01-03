import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';
import InMemoryCache from "../../dataprovider/InMemoryCache";

export interface AboutItem {
    id: number|string
    title: string
    email: string
    description: string
    vision: string
    mission: string
    facebook: string
    twitter: string
    instagram: string
}

async function getAbout () {
    let dataProvider = DataProviderFactory.getDataProvider()

    return await dataProvider.getOne("about", {id: 1}).then(resp => {
        return resp.data as AboutItem
    }, error => {
        NotificationManager.error(error.message, `Error Koneksi: ${error.name}`);
        return null
    })
}

const AboutModel = {
    state: {
        data: null,
    },
    actions: {
        init: async (_, { state, actions  }) => {
            let cacheKey = 'fetched_about'
            if(InMemoryCache.getCache(cacheKey)) {
                return state
            }
            InMemoryCache.setCache(cacheKey, true)

            let newState = { ...state }
            if (newState.data === null) {
                newState.data = await getAbout()
                if(newState.data === null) {
                    InMemoryCache.setCache(cacheKey, false)
                }
            }
            return newState
        }
    }
}

export default AboutModel
