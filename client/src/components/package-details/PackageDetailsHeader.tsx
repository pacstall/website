import { HStack, Heading, Text, Box, Flex } from '@chakra-ui/react'
import { FC, useEffect } from 'react'
import PackageInfo from '../../types/package-info'
import InstallNowButton from './InstallNowButton'
import { useFeatureFlag } from '../../state/feature-flags'
import serverConfig from '../../config/server'
import UninstallButton from './UninstallButton'
import UpdateButton from './UpdateButton'
import toTitleCase from '../../util/title-case'

const PackageDetailsHeader: FC<{ data: PackageInfo; isMobile: boolean }> = (
    { data },
    isMobile,
) => {
    const installButtonEnabled = useFeatureFlag(
        flags => flags.packageDetailsPage.installProtocol,
    )

    useEffect(() => {
        console.log(data)
    }, [data])

    return (
        <>
            <HStack justify='space-between'>
                <HStack>
                    <Heading>{toTitleCase(data)}</Heading>
                </HStack>
                {(serverConfig.isPacstore ||
                    (!isMobile && installButtonEnabled)) &&
                    (data.installed ? (
                        <Flex gap='15px' justifyContent='space-between'>
                            <UpdateButton disabled={data.installedVersion === data.version} />
                            <UninstallButton />
                        </Flex>
                    ) : (
                        <InstallNowButton />
                    ))}
            </HStack>

            <Text mt='5'>{data.description}</Text>
        </>
    )
}

export default PackageDetailsHeader
