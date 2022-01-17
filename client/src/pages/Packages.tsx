import axios from "axios";
import { FC, MutableRefObject, useEffect, useRef, useState } from "react";
import { Link } from "react-router-dom";
import Navigation from "../components/Navigation";
import serverConfig from "../config/server";
import type PackageInfo from "../types/package-info";

const Search: FC<{ placeholder: string, onSearch: (filter: string, filterBy: string) => any }> = ({ placeholder, onSearch }) => {
    const inputRef = useRef<HTMLInputElement>() as MutableRefObject<HTMLInputElement>
    const selectRef = useRef<HTMLSelectElement>() as MutableRefObject<HTMLSelectElement>

    return (
        <div className="packages-wrapper ml-auto">
            <div className="my-2 flex sm:flex-row flex-col">
                <div className="flex flex-row mb-1 sm:mb-0">
                    <div className="relative">
                        <select
                            ref={selectRef}
                            className="focus:outline-none appearance-none h-full rounded-l rounded-r border sm:rounded-r-none sm:border-r-0 border-r border-b block appearance-none w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:border-l focus:border-r focus:bg-white focus:border-gray-500">
                            <option selected value="name">Package</option>
                            <option value="maintainer">Maintainer</option>
                        </select>
                        <div
                            className="pointer-events-none absolute inset-y-0 right-0 flex items-center px-2 text-gray-700">
                            <svg className="fill-current h-4 w-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20">
                                <path d="M9.293 12.95l.707.707L15.657 8l-1.414-1.414L10 10.828 5.757 6.586 4.343 8z" />
                            </svg>
                        </div>
                    </div>
                </div>
                <div className="block relative">
                    <span className="h-full absolute inset-y-0 left-0 flex items-center pl-2">
                        <svg viewBox="0 0 24 24" className="h-4 w-4 fill-current text-gray-500">
                            <path
                                d="M10 4a6 6 0 100 12 6 6 0 000-12zm-8 6a8 8 0 1114.32 4.906l5.387 5.387a1 1 0 01-1.414 1.414l-5.387-5.387A8 8 0 012 10z">
                            </path>
                        </svg>
                    </span>
                    <input placeholder={placeholder} ref={inputRef}
                        className="appearance-none rounded-r rounded-l sm:rounded-l-none border border-gray-400 border-b block pl-8 pr-6 py-2 w-full bg-white text-sm placeholder-gray-400 text-gray-700 focus:bg-white focus:placeholder-gray-600 focus:text-gray-700 focus:outline-none" />
                </div>
            </div>

            <button
                onClick={() => onSearch(inputRef.current.value, selectRef.current.value)}
                className="search text-sm bg-gray-300 hover:bg-gray-400 text-gray-800 font-semibold py-2 px-4 rounded">
                Search
            </button>
        </div>
    )
}


const PackageTableRow: FC<{ pkg: PackageInfo, onCommandCopy: (what: string) => any }> = ({ pkg, onCommandCopy }) => (
    <tr key={pkg.name}>
        <td className="px-6 py-4 whitespace-nowrap">
            <div className="flex items-center">
                <div className="ml-4">
                    <div title={pkg.name} className="text-xs font-medium text-gray-900 text-ellipsis overflow-hidden">
                        <Link to={`/packages/${pkg.name}`} >{pkg.name}</Link>
                    </div>
                </div>
            </div>
        </td>
        <td className="px-6 py-4">
            <div title={pkg.maintainer || 'This package is not being maintained.'} className="text-xs whitespace-nowrap text-gray-900">{(pkg.maintainer || '-').split('<')[0].trim()}</div>
        </td>
        <td className="px-6 py-4">
            <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                <span title={pkg.version} className="text-ellipsis overflow-hidden whitespace-nowrap" style={{ maxWidth: '144px' }}>{pkg.version}</span>
            </span>
        </td>
        <td className="px-6 py-4 whitespace-nowrap">
            <span className="px-2 inline-flex text-xs leading-5 rounded bg-gray-600 text-white" style={{ cursor: "copy", display: 'block' }} onClick={() => onCommandCopy(pkg.name)}>
                <td className="whitespace-nowrap">sudo pacstall install {pkg.name}</td>
            </span>
            <input id={pkg.name} type="text" style={{ display: "none", maxWidth: "1px" }} readOnly defaultValue={`sudo pacstall install ${pkg.name}`} />
        </td>
    </tr>
)

const PackageTable: FC<{ packages: PackageInfo[] }> = ({ packages }) => {
    const onCommandCopy = (name: string) => {
        const input: HTMLInputElement | null = document.querySelector(`#${name}`)
        if (!input) {
            return
        }

        input.style.display = "inline"
        input.focus()
        input.select()
        document.execCommand('copy')
        input.style.display = "none"
    }

    return (
        (
            <div className="flex flex-col">
                <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                    <div className="py-2 align-middle sm:px-6 lg:px-8 block" style={{ margin: 'auto' }}>
                        <div className="packages-wrapper overflow-hidden border-b border-gray-200 sm:rounded-lg">
                            <table className="shadow divide-y divide-gray-200">
                                <thead className="bg-gray-50">
                                    <tr>
                                        <th
                                            scope="col"
                                            className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                                        >
                                            Name
                                        </th>
                                        <th
                                            scope="col"
                                            className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                                        >
                                            Maintainer
                                        </th>
                                        <th
                                            scope="col"
                                            className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                                        >
                                            Version
                                        </th>
                                        <th
                                            scope="col"
                                            className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                                        >
                                            Install
                                        </th>
                                    </tr>
                                </thead>
                                <tbody className="bg-white divide-y divide-gray-200">
                                    {packages.map(pkg => <PackageTableRow key={pkg.name} pkg={pkg} onCommandCopy={onCommandCopy} />)}
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        )
    )
}


const Packages: FC = () => {
    const [packages, setPackages] = useState<PackageInfo[]>([])
    const [loading, setLoading] = useState(true)
    const [loadNext, setLoadNext] = useState(false)
    const [keepLoadingPages, setKeepLoadingPages] = useState(true)
    const [page, setPage] = useState(0)
    const [filter, setFilter] = useState("")
    const [filterBy, setFilterBy] = useState("name")
    const [sort, setSort] = useState("")
    const [sortBy, setSortBy] = useState("")
    const [updatePage, setUpdatePage] = useState(false)
    const ref = useRef<HTMLDivElement>()

    useEffect(() => {
        if (page == 0) {
            setLoading(true)
        }

        let url = `/api/packages?page=${page}`
        if (filter.length > 1) {
            url += `&filter=${filter}&filterBy=${filterBy}`
        }

        if (sortBy != "") {
            url += `&sort=${sort}&sortBy=${sortBy}`
        }

        axios.get<PackageInfo[]>(`${serverConfig.host}${url}`).then(pkgs => {

            if (page === 0) {
                setLoading(false)
                setPackages([...pkgs.data])

            } else {
                setPackages([...packages, ...pkgs.data])
            }

            setLoadNext(false)
            if (pkgs.data.length === 0) {
                setKeepLoadingPages(false)
                return
            }
        })
    }, [page, updatePage])

    useEffect(() => {
        if (loadNext) {
            setKeepLoadingPages(true)
            setPage(page + 1)
        }
    }, [loadNext, setLoadNext])

    const onScroll = () => {
        if (packages.length === 0) {
            return
        }

        if (!keepLoadingPages) {
            return
        }

        if (ref.current) {
            const { scrollTop, scrollHeight, clientHeight } = ref.current;
            if (scrollTop + clientHeight > scrollHeight * 0.8) {
                setLoadNext(true)
            }
        }
    }

    const pickRandomPackage = () => {
        if (packages.length === 0) {
            return ''
        }
        console.log(packages)
        const idx = Math.floor(Math.random() * packages.length)
        return packages[idx].packageName || packages[idx].name.split('-').slice(0, -1).join('-') || packages[idx].name
    }
    const [randomPackage, setRandomPackage] = useState(pickRandomPackage())

    useEffect(() => {
        const intervalId = setInterval(() => setRandomPackage(pickRandomPackage), 2500)

        return () => {
            clearInterval(intervalId)
        }
    }, [])

    const onSearch = (filter: string, filterBy: string) => {
        setFilter(filter)
        setFilterBy(filterBy)
        setPage(0)
        setUpdatePage(!updatePage)
    }

    return <>
        <Navigation />
        <div className="no-scrollbar py-8">
            <Search placeholder={randomPackage} onSearch={onSearch} />

            <div className="px-4 overflow-y-auto" style={{ marginLeft: 'auto', marginRight: 'auto', height: 'calc(100vh - 80px)' }} onScroll={() => onScroll()} ref={ref as any}>
                {loading ? <h1 className="text-center">Loading...</h1> : (
                    <>
                        <PackageTable packages={packages} />
                    </>
                )}
            </div>
        </div>
    </>
}

export default Packages