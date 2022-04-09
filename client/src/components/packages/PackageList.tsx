import { Box, Center, Text } from '@chakra-ui/react'
import { FC } from 'react'
import PackageInfo from '../../types/package-info'
import PackageTable from './PackageTable'

const PackageList: FC<{ data: PackageInfo[] }> = ({ data }) => (
    <Box mt='10'>
        <PackageTable packages={data} />
        {data.length === 0 && (
            <Box mt='5'>
                <Center>
                    <Text>No packages found</Text>
                </Center>
            </Box>
        )}
    </Box>
)

export default PackageList
