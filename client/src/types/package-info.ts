export default interface PackageInfo {
    name: string;
    version: string;
    packageName: string;
    maintainer: string;
    description: string;
    url: string;
    runtimeDependencies: string[];
    buildDependencies: string[];
    optionalDependencies: string[];
    breaks: string[];
    gives: string;
    replace: string[];
    hash: string;
    ppa: string[];
    pacstallDependencies: string[];
    patch: string[];
    requiredBy: string[];
}