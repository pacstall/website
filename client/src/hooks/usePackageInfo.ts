import axios from 'axios'
import { useEffect, useState } from 'react'
import serverConfig from '../config/server'
import PackageInfo from '../types/package-info'

const usePackageInfo = (name: string) => {
    const [data, setData] = useState<PackageInfo>()
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)

    useEffect(() => {
        setLoading(true)
        setError(false)
        axios
            .get<PackageInfo>(serverConfig.host + `/api/packages/${name}`)
            .then(result => setData(result.data))
            .catch(() => setError(true))
            .then(() => setLoading(false))
    }, [name])

    return {
        data: data as PackageInfo,
        loading,
        error,
    }
}

export default usePackageInfo
