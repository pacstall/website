import axios from "axios";
import { useEffect } from "react";
import { atom, useRecoilState } from "recoil";
import serverConfig from "../config/server";
import useNotification from "../hooks/useNotification";

export default interface FeatureFlags {
    packageListPageDisabled: boolean;
    packageDetailsPageDisabled: boolean;
    oldSyntax: boolean;
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

const featureFlagsState = atom<FeatureFlagsState>({
    key: 'featureFlagsState',
    default: {
        flags: null,
        loading: true,
        error: false
    }
});

export const useFeatureFlags = () => {
    const [flags, setFlags] = useRecoilState(featureFlagsState)
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

    return flags
}