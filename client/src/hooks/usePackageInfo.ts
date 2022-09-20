import axios from 'axios'
import { useEffect, useState } from 'react'
import serverConfig from '../config/server'
import PackageInfo from '../types/package-info'
import getPacstore from '../util/pacstore'

export const fetchPackageInfo = (name: string): Promise<PackageInfo> =>
    axios.get<PackageInfo>(serverConfig.host + `/api/packages/${name}`).then(res => res.data)

const usePackageInfo = (name: string) => {
    const [data, setData] = useState<PackageInfo>()
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)

    useEffect(() => {
        setLoading(true)
        setError(false)
        fetchPackageInfo(name)
            .then(async result => {
                if (serverConfig.isWeb) {
                    result.installed = false;
                } else {
                    const pacstore = await getPacstore();
                    const installed = await pacstore.getInstalledPackages()
    
                    result.installed = installed.includes(result.name);
                    if (result.installed) {
                        result.installedVersion = await pacstore.getPackageInstalledVersion(result.name);
                    }
                }

                setData(result)
            })
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
