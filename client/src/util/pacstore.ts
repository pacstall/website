declare global {
    interface Window {
        pacChan: {
            getInstalledPackages(): Promise<string[]>
            getPackageInstalledVersion(pkgName: string): Promise<string>
            sendNotification(title: string, message: string): Promise<string>
        }
    }
}

const sleep = (ms: number) => new Promise(r => setTimeout(r, ms))
const res = new Promise<typeof window.pacChan>(async resolve => {
    while (!window.pacChan) {
        await sleep(50)
    }

    console.log('Loaded pacstore native channel.')
    resolve(window.pacChan)
});
const getPacstore = () => res;

export default getPacstore