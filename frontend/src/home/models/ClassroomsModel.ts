import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';

export interface Classroom {
    id: string|number
    level: string
    major: string
    seq: string
}

async function getClassrooms() {
    let searchData = {
        pagination: {
            page: 1,
            perPage: 100,
        },
        sort: {
            field: 'id',
            order: 'DESC'
        },
        filter: {
        },
    }

    const dataProvider = DataProviderFactory.getDataProvider()

    return await dataProvider.getList("classrooms", searchData).then(resp => {
        return resp.data as Array<Classroom>
    }, error => {
        NotificationManager.error(error.message, `Error Koneksi: ${error.name}`);
        return null
    })
}

export {
    getClassrooms
}
