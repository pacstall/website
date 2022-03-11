import axios from "axios"
import React from "react"
import { useState, useEffect } from "react"
import { useLocation, useNavigate } from "react-router-dom"
import serverConfig from "../config/server"
import { PackageInfoPage } from "../types/package-info"

const useQuery = (): [URLSearchParams, (query: URLSearchParams, replace?: boolean) => void] => {
    const { search, pathname } = useLocation();
    const navigate = useNavigate()

    const data = React.useMemo(() => new URLSearchParams(search), [search]);
    const setData = (query: URLSearchParams, replace = false) => navigate(`${pathname}?${query.toString()}`, { replace })

    return [data, setData]
}

const usePackages = () => {
    const [data, setData] = useState<PackageInfoPage>(null as any)
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)
    const [queryParams, setQueryParams] = useQuery()

    useEffect(() => {
        const url = `/api/packages?page=${queryParams.get('page') || 0}&size=${queryParams.get('size') || 25}&sort=${queryParams.get('sort') || ''}&sortBy=${queryParams.get('sortBy') || 'default'}&filter=${queryParams.get('filter') || ''}&filterBy=${queryParams.get('filterBy') || 'name'}`
        setLoading(true)
        setError(false)
        axios.get<PackageInfoPage>(`${serverConfig.host}${url}`).then(res => {
            setData(res.data)

            if (!queryParams.get('page')) {
                queryParams.set('page', '0')
                setQueryParams(queryParams, true)
            }
            setLoading(false)
        }).catch(() => setError(true))
    }, [queryParams])

    const onSearch = (filter: string, filterBy: string) => {
        setLoading(true)

        queryParams.set('page', '0')
        queryParams.set('filter', filter)
        queryParams.set('filterBy', filterBy)
        setQueryParams(queryParams)
    }

    return {
        data,
        loading,
        error,
        onSearch,
        loaded: !loading && !error
    }
}

export default usePackages