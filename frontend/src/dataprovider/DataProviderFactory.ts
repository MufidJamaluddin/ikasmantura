import { DataProvider } from 'ra-core';
import { fetchUtils } from 'react-admin';

import jsonServerProvider from 'ra-data-json-server';

const httpClient = (url, options: fetchUtils.Options = {}) => {

    let token = localStorage.getItem("token")

    if (!options.headers) {
        options.headers = new Headers({
            Accept: 'application/json',
            "Content-Type": "application/json",
        });
    }

    if (token !== null) {
        options.headers["Authorization"] = `Basic ${token}`
    }

    // add your own headers here
    //options.headers.set('X-Custom-Header', 'foobar');
    return fetchUtils.fetchJson(url, options);
};

export default class DataProviderFactory
{
    private static dataProvider?: DataProvider = null;

    public static getDataProvider()
    {
        if(DataProviderFactory.dataProvider !== null)
            return DataProviderFactory.dataProvider

        if(
            process.env.NODE_ENV === "development"
            || process.env.NODE_ENV === "test"
        )
        {
            const fnFakeData = require('./fakedata').default
            const fnFakeDataProvider = require('ra-data-fakerest').default

            DataProviderFactory.dataProvider =
                fnFakeDataProvider(fnFakeData())
        }
        else
        {

            DataProviderFactory.dataProvider =
                jsonServerProvider(process.env.PUBLIC_URL + '/api/v1', httpClient);
        }

        return DataProviderFactory.dataProvider
    }
}