import { HStack, Heading, Text } from '@chakra-ui/react'
import { FC } from 'react'
import PackageInfo from '../../types/package-info'

const PackageDetailsHeader: FC<{ data: PackageInfo; isMobile: boolean }> = (
    { data },
    isMobile,
) => {
    return (
        <>
            <HStack justify='space-between'>
                <HStack>
                    <Heading>{data.packageName}</Heading>
                </HStack>
            </HStack>

            <Text mt='5'>{data.description}</Text>
        </>
    )
}

export default PackageDetailsHeader
