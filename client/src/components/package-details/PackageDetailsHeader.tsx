import { HStack, Heading, Image, Text } from "@chakra-ui/react";
import { FC } from "react";
import PackageInfo from "../../types/package-info";
import InstallNowButton from "./InstallNowButton";
// @ts-ignore:next-line
import DefaultAppImg from "../../../public/app.png";
import { useFeatureFlag } from "../../state/feature-flags";


const toTitle = (str: string): string => {
    const parts = str.split('-')

    if (['deb', 'git', 'app', 'bin'].includes(parts[parts.length - 1])) {
        parts.pop()
    }

    return parts
        .map(part => part[0].toUpperCase() + part.substring(1))
        .join(' ')
}

const PackageDetailsHeader: FC<{ data: PackageInfo, isMobile: boolean }> = ({ data: { name, description } }, isMobile) => {
    const installButtonEnabled = useFeatureFlag(flags => flags.packageDetailsPage.installProtocol)

    return (
        <>
            <HStack justify='space-between'>
                <HStack>
                    <Image src={DefaultAppImg} maxW='64px' />
                    <Heading>{toTitle(name)}</Heading>
                </HStack>
                {!isMobile && installButtonEnabled && <InstallNowButton />}
            </HStack>

            <Text mt='5'>{description}</Text>
        </>
    )
}

export default PackageDetailsHeader