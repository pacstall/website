export default interface PackageInfo {
    name: string
    version: string
    packageName: string
    maintainer: string
    description: string
    url: string
    runtimeDependencies: string[]
    buildDependencies: string[]
    optionalDependencies: string[]
    breaks: string[]
    gives: string
    replace: string[]
    hash?: string
    ppa: string[]
    pacstallDependencies: string[]
    patch: string[]
    requiredBy: string[]
    latestVersion?: string
    prettyName: string
    updateStatus: UpdateStatus
    installed: boolean
    installedVersion: string
}

export enum UpdateStatus {
    Unknown = -1,
    Latest = 0,
    Major = 3,
    Minor = 2,
    Patch = 1,
}

export interface PackageInfoPage {
    page: number
    size: number
    filter: string
    filterBy: string
    sort: string
    sortBy: string
    lastPage: number
    total: number
    data: PackageInfo[]
}

export interface Page {
    page: number
    size: number
    filter: string
    filterBy: string
    sort: string
    sortBy: string
}
