import PackageInfo from "../types/package-info"

export default function toTitleCase(pkg: PackageInfo): string {
    if (pkg.prettyName && !pkg.prettyName.includes('-') && pkg.prettyName !== pkg.prettyName.toLowerCase()) {
        return pkg.prettyName
    }


    const parts = pkg.name.split('-')

    if (['deb', 'git', 'app', 'bin'].includes(parts[parts.length - 1])) {
        parts.pop()
    }

    return parts
        .map(part => part[0].toUpperCase() + part.substring(1))
        .join(' ')
}