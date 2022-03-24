import axios from "axios"
import { useEffect, useState } from "react"
import serverConfig from "../config/server"
import PackageRequiredBy from "../types/package-requiredby"

const usePackageRequiredBy = (name: string) => {
    const [data, setData] = useState<PackageRequiredBy>([])
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)

    useEffect(() => {
        setError(false)
        setLoading(true)
        axios.get<PackageRequiredBy>(serverConfig.host + `/api/packages/${name}/requiredBy`)
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
        loaded: !loading && !error
    }
}

export default usePackageRequiredBy