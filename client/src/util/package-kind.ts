export enum PackageKind {
    DebFile,
    Source,
    App,
    Binary
}

export default function getPackageKind(pkgName: string): PackageKind {
    switch (true) {
        case pkgName.endsWith('-deb'): return PackageKind.DebFile
        case pkgName.endsWith('-app'): return PackageKind.App
        case pkgName.endsWith('-bin'): return PackageKind.Binary
        default: return PackageKind.Source
    }
}