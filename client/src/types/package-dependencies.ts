import PackageInfo from "./package-info";

export default interface PackageDependencies {
    runtimeDependencies: (PackageInfo | string)[];
    buildDependencies: (PackageInfo | string)[];
    optionalDependencies: (PackageInfo | string)[];
    pacstallDependencies: (PackageInfo | string)[];
}