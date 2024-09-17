export default interface PackageInfo {
    architectures: string[]
    version: string
    packageName: string
    maintainers: string[]
    description: string
    source: ArchDistroString[]
    runtimeDependencies: ArchDistroString[]
    buildDependencies: ArchDistroString[]
    optionalDependencies: ArchDistroString[]
    checkDependencies: ArchDistroString[]
    pacstallDependencies: ArchDistroString[]
    conflicts: ArchDistroString[]
    gives: ArchDistroString
    replaces: ArchDistroString[]
    sha1sums: ArchDistroString[]
    sha224sums: ArchDistroString[]
    sha256sums: ArchDistroString[]
    sha384sums: ArchDistroString[]
    sha512sums: ArchDistroString[]
    md5sums: ArchDistroString[]
    priority: ArchDistroString[]
    requiredBy: string[]
    suggests: ArchDistroString[]
    recommends: ArchDistroString[]
    latestVersion?: string
    prettyName: string
    updateStatus: UpdateStatus
    lastUpdatedAt: string
    enhances: ArchDistroString[]
    changelog: string
    backup: string[]
    compatible: string[]
    incompatible: string[]
    epoch: string
    install: string
    license: string[]
    mask: string[]
    noExtract: string[]
    validPgpKeys: string[]
    groups: string[]
}

export interface ArchDistroString {
    arch?: string
    distro?: string
    value: string
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
