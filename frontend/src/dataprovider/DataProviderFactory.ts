import { DataProvider } from 'ra-core';
import {CreateParams, fetchUtils, UpdateParams } from 'react-admin';

import jsonServerProvider from 'ra-data-json-server';

const httpClient = (url, options: fetchUtils.Options = {}) => {
    return fetchUtils.fetchJson(url, options);
};

export default class DataProviderFactory
{
    private static dataProvider?: DataProvider = null;

    public static getDataProvider()
    {
        if(DataProviderFactory.dataProvider !== null)
            return DataProviderFactory.dataProvider

        /*
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
         */
            const apiUrl = process.env.PUBLIC_URL

            DataProviderFactory.dataProvider =
                jsonServerProvider(apiUrl + '/api/v1', httpClient);

            DataProviderFactory.dataProvider.create = (resource: string, params: CreateParams) =>
                httpClient(`${apiUrl}/${resource}`, {
                    method: 'POST',
                    body: (
                        params.data instanceof FormData ?
                            params.data : JSON.stringify(params.data)
                    ),
                }).then(({ json }) => ({
                    data: { ...params.data, id: json.id },
                }))

            DataProviderFactory.dataProvider.update = (resource: string, params: UpdateParams) =>
                httpClient(`${apiUrl}/${resource}/${params.id}`, {
                    method: 'PUT',
                    body: (
                        params.data instanceof FormData ?
                            params.data : JSON.stringify(params.data)
                    ),
                }).then(({ json }) => ({ data: json }))
        //}

        return DataProviderFactory.dataProvider
    }
}
