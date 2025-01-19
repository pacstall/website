export default interface PackageInfo {
    packageName: string
    prettyName: string
    packageBase: string
    baseIndex: int
    baseTotal: int
    baseChildren: string[]
    description: string
    version: string
    sourceVersion: string
    release: string
    epoch: string
    latestVersion?: string
    homepage: string
    priority: string
    architectures: string[]
    license: string[]
    gives: ArchDistroString
    runtimeDependencies: ArchDistroString[]
    checkDependencies: ArchDistroString[]
    buildDependencies: ArchDistroString[]
    optionalDependencies: ArchDistroString[]
    pacstallDependencies: ArchDistroString[]
    checkConflicts: ArchDistroString[]
    buildConflicts: ArchDistroString[]
    conflicts: ArchDistroString[]
    provides: ArchDistroString[]
    breaks: ArchDistroString[]
    replaces: ArchDistroString[]
    enhances: ArchDistroString[]
    recommends: ArchDistroString[]
    suggests: ArchDistroString[]
    mask: string[]
    compatible: string[]
    incompatible: string[]
    maintainers: string[]
    source: ArchDistroString[]
    noExtract: string[]
    noSubmodules: string[]
    md5sums: ArchDistroString[]
    sha1sums: ArchDistroString[]
    sha224sums: ArchDistroString[]
    sha256sums: ArchDistroString[]
    sha384sums: ArchDistroString[]
    sha512sums: ArchDistroString[]
    backup: string[]
    repology: string[]
    requiredBy: string[]
    updateStatus: UpdateStatus
    lastUpdatedAt: string
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
