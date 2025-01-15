import axios from 'axios'
import { useEffect, useState } from 'react'
import serverConfig from '../config/server'
import PackageDependencies from '../types/package-dependencies'

const usePackageDependencies = (name: string) => {
    const [data, setData] = useState<PackageDependencies>()
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)

    useEffect(() => {
        const fetchDependencies = async () => {
            setError(false)
            setLoading(true)

            try {
                const result = await axios.get<PackageDependencies>(
                    `${serverConfig.host}/api/packages/${name}/dependencies`,
                )

                const pacdeps = result.data.pacstallDependencies || []

                const descriptionDeps = await Promise.all(
                    pacdeps.map(async dep => {
                        try {
                            const depResult = await axios.get(
                                `${serverConfig.host}/api/packages/${dep.value}`,
                            )
                            return {
                                ...dep,
                                description: depResult.data.description,
                            }
                        } catch (e) {
                            console.error(
                                `Failed to fetch description for ${dep.value}`,
                                e,
                            )
                            return { ...dep, description: null }
                        }
                    }),
                )

                setData({
                    ...result.data,
                    pacstallDependencies: descriptionDeps,
                })
                setLoading(false)
            } catch (e) {
                console.error('Error fetching package dependencies:', e)
                setError(true)
                setLoading(false)
            }
        }

        fetchDependencies()
    }, [name])

    return {
        data,
        loading,
        error,
        loaded: !loading && !error && data,
    }
}

export default usePackageDependencies
