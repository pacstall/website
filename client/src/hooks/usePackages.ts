import Joi from 'joi'
import { useEffect, useState } from 'react'
import PackageInfo, { PackageInfoPage, Page } from '../types/package-info'
import useCache from './useCache'
import useQuery from './useQuery'
import { useFetcher, UseFetcherResult } from './useFetcher'
import getPacstore from '../util/pacstore'
import serverConfig from '../config/server'

const useFetchPackages = (page: Page): UseFetcherResult<PackageInfoPage> => {
    const cache = useCache<Page, PackageInfoPage>('packages')
    return useFetcher<PackageInfoPage>(`/api/packages`, { params: page, cache })
}

const mapToPacstorePackages = async (pkgs: PackageInfo[]): Promise<PackageInfo[]> => {
    const pacstore = await getPacstore();
    const installed = await pacstore.getInstalledPackages();
    const mappedData = pkgs.map(pkg => {
        if (installed.includes(pkg.name)) {
            return {
                ...pkg,
                installed: true,
            }
        }

        return {
            ...pkg,
            installed: false
        }
    });

    return mappedData;
}

const usePackages = () => {
    const [queryParams, setQueryParams] = useQuery<Page>(
        Joi.object({
            page: Joi.alt().try(Joi.number().min(0)).default(0),
            size: Joi.alt().try(Joi.number().integer().min(1)).default(18),
            sortBy: Joi.alt()
                .try(Joi.string().valid('default'))
                .default('name'),
            sort: Joi.alt()
                .try(Joi.string().valid('asc', 'desc'))
                .default('asc'),
            filter: Joi.alt().try(Joi.string().allow('')).default(''),
            filterBy: Joi.alt()
                .try(Joi.string().valid('name', 'maintainer'))
                .default('name'),
        }),
        {
            page: 0,
            size: 18,
            sortBy: 'default',
            sort: 'asc',
            filter: '',
            filterBy: 'name',
        },
    );

    const [{ error, loading, data }, setState] = useFetchPackages(queryParams)

    useEffect(() => {
        if (!error && !loading && queryParams.page > data.lastPage) {
            setQueryParams(
                {
                    ...queryParams,
                    page: data.lastPage,
                },
                true,
            )
        }
    }, [loading])

    useEffect(() => {
        if (serverConfig.isPacstore && !error && !loading && queryParams.page <= data.lastPage) {
            mapToPacstorePackages(data.data).then(pkgs => {
                setState({
                    error,
                    loading,
                    data: {
                        ...data,
                        data: pkgs
                    }
                });
            });
        }
    }, [loading])

    const onSearch = (filter: string, filterBy: string) => {
        setQueryParams({
            ...queryParams,
            page: 0,
            filter,
            filterBy,
        })
    }

    return {
        data,
        loading,
        error,
        onSearch,
        loaded: !loading && !error,
    }
}

export default usePackages
