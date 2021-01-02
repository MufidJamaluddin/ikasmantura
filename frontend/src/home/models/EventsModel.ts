import DataProviderFactory from "../../dataprovider/DataProviderFactory";
import {NotificationManager} from 'react-notifications';
import authProvider from "../../dataprovider/authProvider";

async function getEventById(id: string | number) {
    let dataProvider = DataProviderFactory.getDataProvider()
    return await dataProvider.getOne('events', {id: id}).then(resp => {
        return resp.data
    }, error => {
        NotificationManager.error(error.message, `Error Koneksi: ${error.name}`);
        return null
    })
}

async function registerEventById(id: string | number) {
    let dataProvider = DataProviderFactory.getDataProvider()
    let isLogin = await authProvider.checkAuth().then(() => {
        return true
    }, () => {
        NotificationManager.error('Anda harus masuk terlebihdahulu', 'Anda Belum Login')
        return false
    })

    if (isLogin) {
        return await dataProvider.create(`eventregister/${id}`, {data: null})
            .then(_ => {
                NotificationManager.success('Anda telah terdaftar', 'Pendaftaran Sukses')
                return true
            }, error => {
                NotificationManager.error(error.message, `Pendaftaran Gagal: ${error.name}`)
                return false
            })
    }
    return false
}

function initEvents({ start, end, successCallback, failureCallback }) {
    const dataProvider = DataProviderFactory.getDataProvider()

    let params:any = {
        pagination: {
            page: 1,
            perPage: 1000,
        },
        sort: {
            field: '',
            order: '',
        },
        filter: {
            start_gte: start,
            start_lte: end
        },
    }

    try
    {
        dataProvider.getList('events', params).then(value => {
            successCallback(value.data);
        }, error => {
            NotificationManager.error(error.message, `Error Koneksi: ${error.name}`)
            failureCallback(error);
        })
    }
    catch (e)
    {
        NotificationManager.error(e.toString(), 'Kesalahan Teknis!');
    }
}

export {
    getEventById,
    registerEventById,
    initEvents
}
