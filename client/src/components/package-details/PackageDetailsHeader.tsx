import { HStack, Heading, Text } from '@chakra-ui/react'
import { FC } from 'react'
import PackageInfo from '../../types/package-info'
import InstallNowButton from './InstallNowButton'

const PackageDetailsHeader: FC<{ data: PackageInfo; isMobile: boolean }> = (
    { data },
    isMobile,
) => {
    return (
        <>
            <HStack justify='space-between'>
                <HStack>
                    <Heading>{data.name}</Heading>
                </HStack>
                {!isMobile && <InstallNowButton />}
            </HStack>

            <Text mt='5'>{data.description}</Text>
        </>
    )
}

export default PackageDetailsHeader
