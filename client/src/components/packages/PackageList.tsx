import { Box, Center, Text, chakra } from '@chakra-ui/react'
import { FC } from 'react'
import PackageInfo from '../../types/package-info'
import PackageTable from './PackageTable'

const PackageList: FC<{ data: PackageInfo[] }> = ({ data }) => (
    <Box mt='10'>
        <PackageTable packages={data} />
        {data.length === 0 && (
            <Box my='5'>
                <Center>
                    <Text>
                        Not finding the App you wanted?{' '}
                        <chakra.a
                            href='https://github.com/pacstall/pacstall-programs/issues/new?assignees=&labels=package+request&template=PACKAGE-REQUEST.yml&title=PacReq%3A+%3Cname_of_the_package%3E'
                            target='_blank'
                            color='pink.400'
                            textDecoration='underline'
                            fontWeight='800'
                        >
                            Create a request!
                        </chakra.a>
                    </Text>
                </Center>
            </Box>
        )}
    </Box>
)

export default PackageList
