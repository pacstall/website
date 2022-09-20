const cp = require('child_process');
const { promisify } = require('util')
const execAsync = promisify(cp.exec)
const { Notification } = require('electron')

const logChannelMessage = (method, args) => process.env['PACSTORE_DEBUG'] && console.log(method, args)

module.exports.getInstalledPackages = async () => {
    logChannelMessage('getInstalledPackages', {});
    const { stdout } = await execAsync('pacstall -L')
    return stdout.split('\n').map(it => it.trim()).filter(it => !!it)
}

module.exports.getPackageInstalledVersion = async (/** @type {string} */ pkgName) => {
    logChannelMessage('getPackageInstalledVersion', {
        pkgName,
    });

    const { stdout } = await execAsync('pacstall -Qi ' + pkgName)
    return stdout
        .split('\n')
        .map(it => it.trim())
        .find(it => it.includes('version'))
        .split('\x1B[1;32mversion\x1B(B\x1B[m: ')[1]
}

module.exports.sendNotification = async (/** type {string} */ title, /** type {string} */ message) => {
    logChannelMessage('sendNotification', {
        title,
        message
    });
    new Notification({
        title: `${title} - Pacstore`,
        body: message,
        urgency: 'normal'
    }).show()
}