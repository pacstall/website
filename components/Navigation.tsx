import Link from "next/link";
import { useRouter } from "next/router";
import { FC, useState } from "react";

const Navigation: FC = () => {
    const { pathname } = useRouter()
    const [downloadHover, setDownloadHover] = useState(false)

    const active = (path: string) => path === pathname ? 'uk-active' : ''

    return (
        <nav className="uk-navbar-container" uk-navbar={'true'}>
            <div className="sel-target: .uk-navbar-container; cls-active: uk-navbar-sticky">

                <ul className="uk-navbar-nav">
                    <li className={active("/")}><Link href="/">Home</Link></li>
                    <li onMouseEnter={() => setDownloadHover(true)}>
                        <Link href="#">Download</Link>
                        <div className={'uk-navbar-dropdown ' + (downloadHover ? 'uk-open' : '')} onMouseLeave={() => setDownloadHover(false)}>
                            <ul className="uk-nav uk-navbar-dropdown-nav">
                                <li><Link href="https://github.com/pacstall/pacstall#installing">Install Guide</Link></li>
                                <li><Link href="https://github.com/pacstall/pacstall/releases/latest">Releases</Link></li>
                                <li><Link href="https://github.com/pacstall/pacstall-programs">Packages</Link></li>
                            </ul>
                        </div>
                    </li>
                    <li className={active("/showcase")}><Link href="/showcase">Showcase</Link></li>
                    <li><Link href="https://discord.com/invite/sWB6YtKyvW">Discord</Link></li>
                    <li className={active("/packages")}><Link href="/packages">Packages</Link></li>
                </ul>

            </div>
        </nav>
    )
}

export default Navigation