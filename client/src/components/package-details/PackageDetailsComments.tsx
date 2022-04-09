import { Box, Heading } from '@chakra-ui/react'
import { FC } from 'react'
import { useFeatureFlag } from '../../state/feature-flags'

// @ts-ignore:next-line
const PackageDetailsComments: FC = () => {
    const commentsEnabled = useFeatureFlag(
        flags => flags.packageDetailsPage.comments,
    )

    return (
        commentsEnabled && (
            <Box mt='10'>
                <Heading size='lg'>Latest Comments</Heading>
            </Box>
        )
    )
}

export default PackageDetailsComments
