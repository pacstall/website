import { Task } from 'true-parallel/dist/src'

import type { NextApiRequest, NextApiResponse } from 'next'
import { a } from './test'

export type PackageDto = {
    name: string;
    packageName?: string;
    maintainer?: string;
    description?: string;
    url: string;
    runtimeDependencies: string[]
    buildDependencies: string[]
    optionalDependencies: {
        name: string;
        description?: string;
    }[];
    breaks: string[];
    version: string;
    hash?: string;
    ppa?: string;
    pacstallDependencies: string[];
    patch?: string;
}

const LocalRepositoryPath = '/home/saenai/Projects/pacstall-programs'

const PacstallVars = {
    name: 'name',
    packageName: 'pkgname',
    maintainer: 'maintainer',
    description: 'description',
    url: 'url',
    gives: 'gives',
    version: 'version',
    hash: 'hash',
}

const PacstallArrays = {
    buildDependencies: 'build_depends',
    runtimeDependencies: 'depends',
    replace: 'replace',
    breaks: 'breaks',
}

const PacstallMaps = {
    patch: 'patch',
    optionalDependencies: 'optdepends',
    pacstallDependencies: 'pacdeps',
    ppa: 'ppa',
}

const parsePackage = async (name: string): Promise<PackageDto | null> => {

    const task = Task.of(async ({ LocalRepositoryPath, PacstallArrays, PacstallVars, PacstallMaps, name }): Promise<PackageDto | null> => {
        const get = eval('require')
        const fs = get('fs/promises')
        const path = get('path')
        const crypto = get('crypto')
        const cp = get('child_process')

        const workingDirPath = path.join(LocalRepositoryPath, 'packages', name)
        const tempFilePath = path.join(workingDirPath, `${crypto.randomUUID()}.sh`)
        const psPath = path.join(workingDirPath, `${name}.pacscript`)

        let result: PackageDto | null = null
        try {
            const psContent = (await fs.readFile(psPath)).toString()
            const tempFileContent = `
            ${psContent}
            echo {
                ${Object.entries({ ...PacstallVars, ...PacstallArrays })
                    .map(([key, value]) => `echo "    \\"${key}\\": \\"$${value.trim()}\\""`)
                    .join(',\n')}

                echo ","

                ${Object.entries(PacstallMaps)
                    .map(([key, value], idx, arr) => `

                        echo -n "    \\"${key}\\": ["
                        for val in "\${${value}[@]}"
                        do
                            echo -n " \\"$val\\","
                        done
                        echo " ]${idx < arr.length - 1 ? ',' : ''}"
                        `).join("\n")
                }

                echo -e "\\\\r "
            echo }
            
            
        `

            await fs.writeFile(tempFilePath, tempFileContent)
            cp.execSync(`chmod +x ${tempFilePath}`)
            const output = cp.execSync(`bash ${tempFilePath}`).toString().replace(/, \]/g, " ]")

            const rawPackageInfo = JSON.parse(output)

            for (const key of Object.keys(PacstallArrays)) {
                rawPackageInfo[key] = rawPackageInfo[key].split(' ')
                if (rawPackageInfo[key][0] === '') {
                    rawPackageInfo[key] = []
                }
            }

            for (const key of Object.keys(PacstallMaps)) {
                rawPackageInfo[key] = rawPackageInfo[key].map((line: string) => {
                    let colonIdx = line.lastIndexOf(': ')
                    if (colonIdx <= 0) {
                        return line.trim()
                    }

                    const mapKey = line.substring(0, colonIdx)
                    let mapDescription: string | null = null
                    if (colonIdx < line.length - 1) {
                        mapDescription = line.substring(colonIdx + 1)
                    }

                    const out = {
                        name: mapKey.trim(),
                        description: mapDescription?.trim()
                    }

                    return out
                })
            }

            result = rawPackageInfo

        } catch (e) {
            console.error(`Could not parse file '${psPath}'`, e)
        } finally {
            fs.unlink(tempFilePath).catch(() => {/* no-op */ });
        }

        return result

    }, { PacstallVars, PacstallArrays, PacstallMaps, LocalRepositoryPath, name });

    const result = await task.execute()
    return result
}


const queuedTasks: Array<() => Promise<PackageDto | null>> = []

const parseAllPackages = async (): Promise<PackageDto[]> => {
    require('child_process').execSync(`cd "${LocalRepositoryPath}" && git reset --hard HEAD && git fetch && git pull`)

    const maxTasks = +require('child_process').execSync('nproc --all') || 4
    const fs = require('fs/promises')
    const packageList = (
        await fs.readFile(`${LocalRepositoryPath}/packagelist`)
    ).toString().trim().split('\n') as string[]

    for (const packageName of packageList) {
        queuedTasks.push(() => parsePackage(packageName))
    }

    let out: PackageDto[] = []
    while (queuedTasks.length > 0) {
        const queued = queuedTasks.length
        console.log(`Remaining: ${queued}`, { maxTasks })

        const results = (await Promise.all(
            queuedTasks
                .splice(0, Math.min(maxTasks, queued))
                .map(parser => parser())
        )).filter(dto => dto !== null) as PackageDto[]

        out = [...out, ...results];
    }
    console.log('Done')

    return out
}

export let parsedPackages = parseAllPackages()

setInterval(async () => {
    const packages = await parseAllPackages()
    parsedPackages = Promise.resolve(packages)
}, 5 * 60 * 1000)

export default async function handler(
    req: NextApiRequest,
    res: NextApiResponse<PackageDto[]>
) {
    if (req.method?.toUpperCase() !== "GET") {
        return res.status(405).send('' as any)
    }



    const packages = await parsedPackages
    res.json(packages)
}
