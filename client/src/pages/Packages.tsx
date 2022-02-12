import axios from "axios";
import { FC, useEffect, useState } from "react";
import { Navigate } from "react-router-dom";
import { useRecoilState } from "recoil";
import { NumberParam, StringParam, useQueryParams } from "use-query-params";
import Navigation from "../components/Navigation";
import PackageTable from "../components/packages/PackageTable";
import Search from "../components/packages/Search";
import Pagination from "../components/Pagination";
import serverConfig from "../config/server";
import useNotification from "../hooks/useNotification";
import useRandomPackage from "../hooks/useRandomPackage";
import { useFeatureFlags } from "../state/feature-flags";
import { PackageInfoPage } from "../types/package-info";

const Packages: FC = () => {
    const featureFlags = useFeatureFlags()
    const [pageDisabled, setPageDisabled] = useState<boolean>()
    useEffect(() => {
        setPageDisabled(
            featureFlags.loading
                ? undefined
                : featureFlags.error
                    ? false
                    : featureFlags.flags!.packageListPageDisabled
        )
    }, [featureFlags])

    const notify = useNotification()
    const [packagePage, setPackagePage] = useState<PackageInfoPage>()
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState(false)
    const randomPackageName = useRandomPackage(packagePage?.data)
    const [queryParams, setQueryParams] = useQueryParams({
        page: NumberParam,
        size: NumberParam,
        sort: StringParam,
        sortBy: StringParam,
        filter: StringParam,
        filterBy: StringParam
    })

    useEffect(() => {
        const url = `/api/packages?page=${queryParams.page || 0}&size=${queryParams.size || 25}&sort=${queryParams.sort || ''}&sortBy=${queryParams.sortBy || 'default'}&filter=${queryParams.filter || ''}&filterBy=${queryParams.filterBy || 'name'}`
        setLoading(true)
        axios.get<PackageInfoPage>(`${serverConfig.host}${url}`).then(res => {
            setPackagePage(res.data)
            setLoading(false)
        }).catch(() => setError(true))
    }, [queryParams])



    const onSearch = (filter: string, filterBy: string) => {
        setQueryParams({ filter, filterBy, page: 0 })
    }

    if (pageDisabled) {
        notify({
            title: 'This feature is not ready yet.',
            text: 'You are being redirected back to the home page.',
            type: 'info'
        })
        return <Navigate to="/" />
    }

    if (error) {
        return <Navigate to="/packages" />
    }

    return <>
        <Navigation />
        <div className="no-scrollbar py-8">
            <Search placeholder={randomPackageName} onSearch={onSearch} />

            {loading ? <h1 className="text-center">Loading...</h1> : (
                <>
                    <div className="flex justify-center">
                        <PackageTable linksDisabled={!!featureFlags.flags?.packageDetailsPageDisabled} packages={packagePage!.data} />
                    </div>

                    <div className="flex justify-center my-8">
                        <Pagination current={packagePage!.page} last={packagePage!.lastPage} />
                    </div>
                </>
            )}


        </div>
    </>
}

export default Packages