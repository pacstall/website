import axios, { AxiosRequestConfig } from 'axios'
import { useEffect, useState } from 'react'
import serverConfig from '../config/server'
import { Cache } from './useCache'


export type UseFetcherState<T> = {
    data?: T
    loading: boolean
    error: boolean
}
export type UseFetcherResult<T> = [UseFetcherState<T>, (state: UseFetcherState<T>) => any]

type UseFetcherOptions<T> = { cache?: Cache<any, T> } & AxiosRequestConfig

export function useFetcher<T>(
    url: string,
    options: UseFetcherOptions<T> = {},
): UseFetcherResult<T> {
    const [state, setState] = useState<UseFetcherState<T>>({
        data: null,
        loading: true,
        error: false,
    })

    useEffect(() => {
        ;(async () => {
            setState({ data: null, loading: true, error: false })
            try {
                const fetcher = async () =>
                    await axios
                        .get<T>(`${serverConfig.host}${url}`, options)
                        .then(res => res.data)
                let data: T
                if (options.cache) {
                    data = await options.cache.use(
                        { ...options, cache: undefined },
                        fetcher,
                    )
                } else {
                    data = await fetcher()
                }

                setState({ data, loading: false, error: false })
            } catch (error) {
                setState({ data: null, loading: false, error: true })
            }
        })()
    }, [url, options.params])

    return [state, setState]
}
