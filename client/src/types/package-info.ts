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

export interface PackageInfoPage {
    page: number;
    size: number;
    filter: string;
    filterBy: string;
    sort: string;
    sortBy: string;
    lastPage: number;
    total: number;
    data: PackageInfo[];
} 