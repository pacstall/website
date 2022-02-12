import { FC } from "react"
import PackageInfo from "../../types/package-info"
import PackageTableRow from "./PackageTableRow"

const PackageTable: FC<{ packages: PackageInfo[], linksDisabled?: boolean }> = ({ packages, linksDisabled }) => {
    return (
        (
            <div className="flex flex-col">
                <div className="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
                    <div className="py-2 align-middle sm:px-6 lg:px-8 block" style={{ margin: 'auto' }}>
                        <div className="container overflow-hidden border-b border-gray-200 sm:rounded-lg">
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
                                    {packages.map(pkg => <PackageTableRow key={pkg.name} disabled={linksDisabled} pkg={pkg} />)}
                                </tbody>
                            </table>
                        </div>
                    </div>
                </div>
            </div>
        )
    )
}

export default PackageTable