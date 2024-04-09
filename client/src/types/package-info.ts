export default interface PackageInfo {
    version: string
    packageName: string
    maintainer: string
    description: string
    source: string[]
    runtimeDependencies: string[]
    buildDependencies: string[]
    optionalDependencies: string[]
    conflicts: string[]
    gives: string
    replaces: string[]
    hash?: string
    ppa: string[]
    pacstallDependencies: string[]
    patch: string[]
    requiredBy: string[]
    latestVersion?: string
    prettyName: string
    updateStatus: UpdateStatus
    lastUpdatedAt: string
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
