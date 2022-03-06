import axios from "axios"
import { useState, useEffect } from "react"
import { useQueryParams, NumberParam, StringParam } from "use-query-params"
import serverConfig from "../config/server"
import { PackageInfoPage } from "../types/package-info"

const usePackages = () => {
    const [data, setData] = useState<PackageInfoPage>(null as any)
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)
    const [queryParams, setQueryParams] = useQueryParams({
        page: {
            decode: val => +val,
            encode: val => val,
            equals: (v1, v2) => v1 == v2
        },
        size: NumberParam,
        sort: StringParam,
        sortBy: StringParam,
        filter: StringParam,
        filterBy: StringParam
    })

    useEffect(() => {
        const url = `/api/packages?page=${queryParams.page || 0}&size=${queryParams.size || 25}&sort=${queryParams.sort || ''}&sortBy=${queryParams.sortBy || 'default'}&filter=${queryParams.filter || ''}&filterBy=${queryParams.filterBy || 'name'}`
        setLoading(true)
        setError(false)
        axios.get<PackageInfoPage>(`${serverConfig.host}${url}`).then(res => {
            setData(res.data)
            setQueryParams({ page: queryParams.page || 0 })
            setLoading(false)
        }).catch(() => setError(true))
    }, [queryParams])

    const onSearch = (filter: string, filterBy: string) => {
        setLoading(true)
        setQueryParams({ filter, filterBy, page: 0 })
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