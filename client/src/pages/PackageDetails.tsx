import axios from "axios";
import { FC, useEffect, useState } from "react";
import { useParams, Navigate } from "react-router-dom";
import Navigation from "../components/Navigation";
import serverConfig from "../config/server";
import PackageInfo from "../types/package-info";
import DefaultAppImg from "../../public/app.png";

const toTitle = (str: string): string => {
    const parts = str.split('-')

    if (['deb', 'git', 'app', 'bin'].includes(parts[parts.length - 1])) {
        parts.pop()
    }

    return parts
        .map(part => part[0].toUpperCase() + part.substring(1))
        .join(' ')

}

const Panel: FC<{ title: string, showIcon?: boolean, version?: string }> = ({ title, version, showIcon, children }) => (
    <div className="panel">
        {showIcon ? (
            <div className="panel-icon-title">
                <img width="64" height="64" src={DefaultAppImg} alt={`${title} logo`} />
                <h1>{title}</h1>
                <span>
                    <span className="px-2 inline-flex text-sm leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                        {version}
                    </span>
                </span>
            </div>
        ) : <h1>{title}</h1>}

        <div className="py-5 px-3">
            {children}
        </div>
    </div>
)

const Maintainer: FC<{ text: string }> = ({ text }) => {
    if (!text || text === '-' || text.toLowerCase() === 'orphan' || text.toLowerCase() === 'orphaned') {
        return <>Orphaned</>
    }

    if (!['<', '>', '@'].every(symbol => text.includes(symbol))) {
        return <>{text}</>
    }

    const shortenName = (name: string, splitBy: string): string =>
        name.split(splitBy).reduce((acc, part) => (acc + part).length > 14 ? acc : acc + splitBy + part, '')

    let name = text.split('<')[0].trim()
    if (name.length > 15) {
        if (name.includes(' ')) {
            name = shortenName(name, ' ')
        } else if (name.includes('-')) {
            name = shortenName(name, '-')
        } else {
            name = name.substring(0, 12) + '..'
        }
    }

    const fullEmail = text.split('<')[1].split('>')[0].trim()
    const shortEmail = fullEmail.split('@')[0].length > 15 ? fullEmail.split('@')[0].substring(0, 13) + '[..]@' + fullEmail.split('@')[1] : fullEmail

    return (
        <>
            <span>{name}, </span>
            <a className="uk-link" href={"mailto: " + fullEmail}>
                {shortEmail}
            </a>
        </>
    )
}

const PackageDetails: FC = () => {
    const { name } = useParams() as { name: string }
    const [pkgInfo, setPkgInfo] = useState<PackageInfo>()
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        setLoading(true)
        axios.get<PackageInfo>(serverConfig.host + `/api/packages/${name}`)
            .then(result => setPkgInfo(result.data))
            .catch(() => { })
            .then(() => setLoading(false))
    }, [name])

    if (!loading && !pkgInfo) {
        return <Navigate to="/not-found" />
    }

    if (loading) {
        return <>Loading</>
    }

    const allDependencies = [...pkgInfo!.buildDependencies, ...pkgInfo!.runtimeDependencies, ...pkgInfo!.optionalDependencies, ...pkgInfo!.pacstallDependencies]

    return (
        <>
            <Navigation />

            <div className="container py-8 panel-container">
                <Panel showIcon title={toTitle(name)} version={pkgInfo!.version}>
                    <div className="text-gray-700" style={{ marginBottom: '1.5em' }}>{pkgInfo!.description}</div>
                    <div className="panel-inverse-table">
                        <span>Name</span>
                        <span>{pkgInfo!.name}</span>

                        <span>Maintainer</span>
                        <span><Maintainer text={pkgInfo!.maintainer} /></span>

                        <span>Last Updated</span>
                        <span>Today</span>

                        <span>Votes</span>
                        <span>{Math.floor(Math.random() * 1200) + 30}</span>

                        <span>Popularity</span>
                        <span>{Math.floor(Math.random() * 1000) / 10}%</span>

                        <span>Dependencies</span>
                        <span>{allDependencies.length || 'None'} {allDependencies.length > 0 ? <a className="uk-link" href={`/packages/${name}/dependencies`}>(see all)</a> : ''}</span>

                        <span>Required By</span>
                        <span>{pkgInfo!.requiredBy.length || 'None'} {pkgInfo!.requiredBy.length > 0 ? <a className="uk-link" href={`/packages/${name}/required-by`}>(see all)</a> : ''}</span>
                    </div>
                </Panel>
                <Panel title="How to Install" />
                <Panel title="Pacscript Details" />
                <Panel title="Comments" />
            </div>
        </>
    )
}

export default PackageDetails