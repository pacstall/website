import { HStack, Heading, Text } from "@chakra-ui/react";
import { FC } from "react";
import PackageInfo from "../../types/package-info";
import InstallNowButton from "./InstallNowButton";
import { useFeatureFlag } from "../../state/feature-flags";


const getTitle = (pkg: PackageInfo): string => {
    if (pkg.prettyName && !pkg.prettyName.includes('-') && pkg.prettyName !== pkg.prettyName.toLowerCase()) {
        return pkg.prettyName
    }


    const parts = pkg.name.split('-')

    if (['deb', 'git', 'app', 'bin'].includes(parts[parts.length - 1])) {
        parts.pop()
    }

    return parts
        .map(part => part[0].toUpperCase() + part.substring(1))
        .join(' ')
}

const PackageDetailsHeader: FC<{ data: PackageInfo, isMobile: boolean }> = ({ data }, isMobile) => {
    const installButtonEnabled = useFeatureFlag(flags => flags.packageDetailsPage.installProtocol)

    return (
        <>
            <HStack justify='space-between'>
                <HStack>
                    <Heading>{getTitle(data)}</Heading>
                </HStack>
                {!isMobile && installButtonEnabled && <InstallNowButton />}
            </HStack>

            <Text mt='5'>{data.description}</Text>
        </>
    )
}

export default PackageDetailsHeader