import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';

async function getOrganizations() {
    let dataProvider = DataProviderFactory.getDataProvider()

    let request = {
        pagination: {
            page: 1,
            perPage: 1000,
        },
        sort: {
            field: 'id',
            order: 'ASC'
        },
        filter: {

        },
    }

    return await dataProvider.getList("departments", request).then(resp => {
        if(resp.total === 0) {
            NotificationManager.warning('Belum ada struktur organisasi!');
        }
        return {
            data: resp.data,
            total: resp.total,
        }
    }, error => {
        NotificationManager.error(error.message, `Error Koneksi: ${error.name}`);
        return null
    })
}

export {
    getOrganizations
}
