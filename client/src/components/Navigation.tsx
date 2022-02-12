import { Link, useLocation } from "react-router-dom";
import { FC, useEffect, useState } from "react";
import { useFeatureFlags } from "../state/feature-flags";

const Navigation: FC = () => {
    const { pathname } = useLocation()
    const [downloadHover, setDownloadHover] = useState(false)
    const featureFlags = useFeatureFlags()

    const active = (path: string) => path === pathname ? 'uk-active' : ''

    return (
        <nav className="uk-navbar-container" uk-navbar={'true'}>
            <div className="sel-target: .uk-navbar-container; cls-active: uk-navbar-sticky">

                <ul className="uk-navbar-nav">
                    <li className={active("/")}><Link to="/">Home</Link></li>
                    <li onMouseEnter={() => setDownloadHover(true)}>
                        <Link to="#">Download</Link>
                        <div className={'uk-navbar-dropdown ' + (downloadHover ? 'uk-open' : '')} onMouseLeave={() => setDownloadHover(false)}>
                            <ul className="uk-nav uk-navbar-dropdown-nav">
                                <li><a href="https://github.com/pacstall/pacstall#installing" target="_blank">Install Guide</a></li>
                                <li><a href="https://github.com/pacstall/pacstall/releases/latest" target="_blank">Releases</a></li>
                                <li><a href="https://github.com/pacstall/pacstall-programs" target="_blank">Packages</a></li>
                            </ul>
                        </div>
                    </li>
                    <li className={active("/showcase")}><Link to="/showcase">Showcase</Link></li>
                    <li><a href="https://discord.com/invite/sWB6YtKyvW" target="_blank">Discord</a></li>
                    <li><a href="https://matrix.to/#/#pacstall:matrix.org" target="_blank">Matrix</a></li>
                    {!featureFlags?.flags?.packageListPageDisabled && <li className={active("/packages")}><Link to="/packages">Packages</Link></li>}
                </ul>

            </div>
        </nav>
    )
}

export default Navigation