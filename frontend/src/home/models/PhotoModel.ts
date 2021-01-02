import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';

export interface Photo {
    id: string|number
    title: string
    image: string
    thumbnail: string
}

async function getPhotoByAlbumIds(albumIds: Array<string | number>) {
    let searchData = {
        pagination: {
            page: 1,
            perPage: 1000,
        },
        sort: {
            field: 'id',
            order: 'ASC'
        },
        filter: {
            albumId: albumIds
        },
    }

    const dataProvider = DataProviderFactory.getDataProvider()

    return await dataProvider.getList("photos", searchData).then(resp => {
        return resp.data as Array<Photo>
    }, error => {
        NotificationManager.error(error.message, `Error Koneksi: ${error.name}`);
        return null
    })
}

export {
    getPhotoByAlbumIds
}
