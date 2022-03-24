import { useEffect, useState } from "react";
import PackageInfo from "../types/package-info";

const pickRandomPackage = (packages: PackageInfo[]) => {
    if (packages.length === 0) {
        return ''
    }
    const idx = Math.floor(Math.random() * packages.length)
    return packages[idx].packageName || packages[idx].name.split('-').slice(0, -1).join('-') || packages[idx].name
}

export default function useRandomPackage(packages?: PackageInfo[]): string {
    const [name, setName] = useState('')

    useEffect(() => {
        const intervalId = setInterval(() => setName(pickRandomPackage(packages || [])), 2500)

        return () => {
            clearInterval(intervalId)
        }
    }, [packages])

    return name
}