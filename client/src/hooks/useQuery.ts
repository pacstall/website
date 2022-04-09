import { Schema } from "joi";
import { useEffect, useMemo, useState } from "react";
import { useLocation, useNavigate, useSearchParams } from "react-router-dom";

const useQuery = <T>(schema: Schema<T>, defaults: T): [T, (query: T, replace?: boolean) => void] => {
    const [params, setParams] = useSearchParams();
    const [trigger, setTrigger] = useState(false);
    const [validParams, setValidParams] = useState<T>(defaults);
    const navigate = useNavigate()
    const location = useLocation()
    const objParams = useMemo(() => [...params.entries()].reduce((acc, [key, value]) => ({ ...acc, [key]: value }), {}), [params]);


    useEffect(() => {
        navigate(location.pathname + '?' + new URLSearchParams({ ...defaults, ...objParams } as any).toString(), { replace: true })
    }, [])

    useEffect(() => {
        const result = schema.validate(objParams, { cache: true, abortEarly: false, stripUnknown: true, convert: true });
        if (result.error) {
            navigate(location.pathname + '?' + new URLSearchParams(defaults as any).toString(), { replace: true })
        } else {
            setValidParams(result.value);
        }

    }, [params, trigger])

    const setQuery = (query: T, replace?: boolean) => {
        const newParams = Object.entries(query).reduce((acc: Record<string, string>, [key, value]: [string, any]) => ({
            ...acc,
            [key]: value === null ? '' : value.toString()
        }), {});

        navigate(location.pathname + '?' + new URLSearchParams(newParams).toString(), { replace });
        setParams(newParams, { replace });
        setTrigger(!trigger);
    }

    return [validParams as T, setQuery]
}


export default useQuery