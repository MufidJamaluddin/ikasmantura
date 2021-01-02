import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';

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
        fetched: false,
    },
    actions: {
        done: (_, { state }) => {
            return { ...state, fetched: true }
        },
        init: async (_, { state, actions  }) => {
            if(state.fetched) {
                return state
            }
            actions.done()

            let newState = { ...state }
            if (newState.data === null) {
                newState.data = await getAbout()
                newState.fetched = newState.data === null
            }
            return newState
        }
    }
}

export default AboutModel
