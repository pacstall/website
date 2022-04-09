import isEqual from 'lodash/isEqual'
import { useMemo } from 'react'

export class Cache<K, V> {
    private cache: Record<string, V> = {}
    constructor(public readonly id: string) { }

    async use(key: K, getter: () => V | Promise<V>): Promise<V> {
        for (const [k, v] of Object.entries(this.cache)) {
            if (isEqual(k, JSON.stringify(key))) {
                return v
            }
        }

        const value = await getter()
        this.cache[JSON.stringify(key)] = value

        return value
    }
}

const cache = new Map<string, Cache<any, any>>();

export default function useCache<K, V>(cacheId: string): Cache<K, V> {
    return useMemo(() => {
        if (cache.has(cacheId)) {
            return cache.get(cacheId)!;
        }

        const newCache = new Cache<K, V>(cacheId);
        cache.set(cacheId, newCache);
        return newCache
    }, [cacheId]);
}