import { HStack, Heading, Text } from '@chakra-ui/react'
import { FC } from 'react'
import PackageInfo from '../../types/package-info'
import InstallNowButton from './InstallNowButton'
import { useFeatureFlag } from '../../state/feature-flags'
import toTitleCase from '../../util/title-case'

const PackageDetailsHeader: FC<{ data: PackageInfo; isMobile: boolean }> = (
    { data },
    isMobile,
) => {
    const installButtonEnabled = useFeatureFlag(
        flags => flags.packageDetailsPage.installProtocol,
    )

    return (
        <>
            <HStack justify='space-between'>
                <HStack>
                    <Heading>{data.prettyName}</Heading>
                </HStack>
                {!isMobile && installButtonEnabled && <InstallNowButton />}
            </HStack>

            <Text mt='5'>{data.description}</Text>
        </>
    )
}

export default PackageDetailsHeader
