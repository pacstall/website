import PackageInfo, { ArchDistroString } from './package-info'

export default interface PackageDependencies {
    runtimeDependencies: (ArchDistroString | string)[]
    buildDependencies: (ArchDistroString | string)[]
    optionalDependencies: (ArchDistroString | string)[]
    pacstallDependencies: (ArchDistroString | string)[]
}
