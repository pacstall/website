import axios from "axios"
import { useEffect, useState } from "react"
import serverConfig from "../config/server"
import PackageDependencies from "../types/package-dependencies"

const usePackageDependencies = (name: string) => {
    const [data, setData] = useState<PackageDependencies>()
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)

    useEffect(() => {
        setError(false)
        setLoading(true)
        axios.get<PackageDependencies>(serverConfig.host + `/api/packages/${name}/dependencies`)
            .then(result => {
                setLoading(false)
                setData(result.data)
            }).catch(() => {
                setError(true)
                setLoading(false)
            })
    }, [name])

    return {
        data,
        loading,
        error,
        loaded: !loading && !error && data
    }
}

export default usePackageDependencies