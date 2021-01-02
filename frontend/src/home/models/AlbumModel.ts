import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';

export interface Album {
    id: string|number
    title: string
}

async function getAlbums() {
    let dataProvider = DataProviderFactory.getDataProvider()
    return await dataProvider.getList("albums", {
        pagination: {
            page: 1,
            perPage: 100,
        },
        sort: {
            field: 'id',
            order: 'DESC'
        },
        filter: {},
    }).then(resp => {
        return resp.data as Array<Album>
    }, error => {
        NotificationManager.error(error.message, `Error Koneksi: ${error.name}`);
        return null
    })
}

export {
    getAlbums
}
