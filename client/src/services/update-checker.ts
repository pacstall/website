import { fetchPackageInfo } from "../hooks/usePackageInfo";
import PackageInfo from "../types/package-info";
import getPacstore from "../util/pacstore";
import toTitleCase from "../util/title-case";

const CHECK_TIMER = 60 * 60 * 1000 // 1h

async function sendUpdatesAvailableNotification(pkgs: PackageInfo[]) {
    const pacstore = await getPacstore()

    const notificationMessage = pkgs.length === 1
        ? `Package ${toTitleCase(pkgs[0])} has an update available.`
        : `Found ${pkgs.length} updates available.`

    await pacstore.sendNotification('Updates Available', notificationMessage);
}

async function updateCheck() {
    const pacstore = await getPacstore()
    const installedPkgs = await pacstore.getInstalledPackages();
    const installedPkgVersions = await Promise.all(installedPkgs.map(async it => {
        const installedVersion = await pacstore.getPackageInstalledVersion(it);
        return {
            name: it,
            installedVersion
        }
    }));

    const remotePkgs = await Promise.all(
        installedPkgs.map(fetchPackageInfo)
    )

    const updatablePackages = remotePkgs
        .filter(remote =>
            installedPkgVersions.find(installed => remote.name === installed.name).installedVersion !== remote.version)

    if (updatablePackages.length === 0) {
        return
    }

    await sendUpdatesAvailableNotification(updatablePackages);
}

export async function registerUpdateChecker() {
    await updateCheck()
    setInterval(updateCheck, CHECK_TIMER);
}
