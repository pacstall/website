/** @type {import('fs/promises')} */
const { randomUUID } = require('crypto')
const fs = require('fs/promises')
const cp = require('child_process')

const useTempFile = async (cb) => {
    const uuid = randomUUID()
    await fs.writeFile(`./${uuid}.sh`, '')
    try {
        await cb(`./${uuid}.sh`)
    } catch (e) {
        throw e
    } finally {
        await fs.unlink(`./${uuid}.sh`)
    }
}

(async () => {


})()