import { atom } from "recoil";
import { PackageInfoPage } from "../types/package-info";

interface PackageInfoPageStateUnloaded {
    data: null;
    loading: false;
    error: false;
}

interface PackageInfoStateSuccess {
    data: PackageInfoPage;
    loading: false;
    error: false;
}

interface PackageInfoStateLoading {
    data: null;
    loading: true;
    error: false;
}

interface PackageInfoStateError {
    data: null;
    loading: false;
    error: true;
}

export type PackagePageState =
    | PackageInfoPageStateUnloaded
    | PackageInfoStateSuccess
    | PackageInfoStateLoading
    | PackageInfoStateError

export const packageInfoState = atom<PackagePageState>({
    key: "packageInfoState",
    default: {
        loading: false,
        error: false,
        data: null
    }
})