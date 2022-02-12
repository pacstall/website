import { FC } from "react";
import { Link } from "react-router-dom";
import PackageInfo from "../../types/package-info";
import { SmartCodeSnippetInstall } from "../OneLineCodeSnippet";

const PackageTableRow: FC<{ pkg: PackageInfo, disabled?: boolean }> = ({ pkg, disabled }) => (
    <tr key={pkg.name}>
        <td className="px-6 py-4 whitespace-nowrap">
            <div className="flex items-center">
                <div className="ml-4">
                    <div title={pkg.name} className="text-xs font-medium text-gray-900 text-ellipsis overflow-hidden">
                        {disabled === true ? <span>{pkg.name}</span> : <Link to={`/packages/${pkg.name}`} >{pkg.name}</Link>}
                    </div>
                </div>
            </div>
        </td>
        <td className="px-6 py-4">
            <div title={pkg.maintainer || 'This package is not being maintained.'} className="text-xs whitespace-nowrap text-gray-900">{(pkg.maintainer || 'Orphaned').split('<')[0].trim()}</div>
        </td>
        <td className="px-6 py-4">
            <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                <span title={pkg.version} className="text-ellipsis overflow-hidden whitespace-nowrap" style={{ maxWidth: '144px' }}>{pkg.version.substring(0, 14)}</span>
            </span>
        </td>
        <td className="px-6 py-4 whitespace-nowrap">
            <SmartCodeSnippetInstall size="sm" name={pkg.name} />
        </td>
    </tr>
)

export default PackageTableRow