import axios from "axios";
import { useEffect } from "react";
import { atom, selector, useRecoilState, useRecoilValue } from "recoil";
import serverConfig from "../config/server";
import useNotification from "../hooks/useNotification";

export default interface FeatureFlags {
    oldSyntax: boolean;
    packageDetailsPage: {
        lastUpdated: boolean;
        votes: boolean;
        popularity: boolean;
        installProtocol: boolean;
        comments: boolean;
    }
}

interface FeatureFlagsStateSuccess {
    flags: FeatureFlags;
    loading: false;
    error: false;
}

interface FeatureFlagsStateLoading {
    flags: null;
    loading: true;
    error: false;
}

interface FeatureFlagsStateError {
    flags: null;
    loading: false;
    error: true;
}

export type FeatureFlagsState =
    | FeatureFlagsStateSuccess
    | FeatureFlagsStateLoading
    | FeatureFlagsStateError

export const featureFlagsState = atom<FeatureFlagsState>({
    key: 'featureFlagsState',
    default: axios.get(`${serverConfig.host}/api/feature-flags`).then(result => ({
        loading: false,
        error: false,
        flags: result.data
    })).catch(() => ({
        loading: false,
        error: true,
        flags: null
    })) as Promise<FeatureFlagsState>
});

const isLoadedSelector = selector<boolean>({
    key: 'isLoadedSelector',
    get: ({ get }) => !get(featureFlagsState).loading && !get(featureFlagsState).error,
})

const isFlagEnabled = selector<(select: (flag: FeatureFlags) => boolean) => boolean>({
    key: 'isFlagEnabled',
    get: ({ get }) => select => get(isLoadedSelector) ? select(get(featureFlagsState).flags!) : null!
})

export const useFeatureFlag = (select: (flag: FeatureFlags) => boolean): boolean =>
    useRecoilValue(isFlagEnabled)(select)

export const useFeatureFlags = () => {
    const [flags, setFlags] = useRecoilState(featureFlagsState)
    const loaded = useRecoilValue(isLoadedSelector)
    const notify = useNotification()

    useEffect(() => {
        setFlags({
            loading: true,
            error: false,
            flags: null
        })

        axios.get(`${serverConfig.host}/api/feature-flags`).then(result => {
            setFlags({
                loading: false,
                error: false,
                flags: result.data
            })
        }).catch(() => {
            setFlags({
                loading: false,
                error: true,
                flags: null
            })
        })

    }, [])

    useEffect(() => {
        if (flags.error) {
            notify({
                title: 'Whoops! Something went wrong.',
                text: 'It looks like our server did not respond successfully',
                type: 'error'
            });
        }
    }, [flags.error])

    return {
        ...flags,
        loaded
    }
}