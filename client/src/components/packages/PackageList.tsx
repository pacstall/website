import { Box, Center, Text, chakra } from '@chakra-ui/react'
import { FC } from 'react'
import PackageInfo from '../../types/package-info'
import PackageTable from './PackageTable'
import { Trans, useTranslation } from 'react-i18next'

const PackageList: FC<{ data: PackageInfo[] }> = ({ data }) => {
    const { t } = useTranslation()

    return (
        <Box mt='10'>
            <PackageTable packages={data} />
            {data.length === 0 && (
                <Box my='5'>
                    <Center>
                        <Text>
                            <Trans i18nKey='packageSearch.noResults'>
                                <chakra.a
                                    href='https://github.com/pacstall/pacstall-programs/issues/new?assignees=&labels=package+request&template=PACKAGE-REQUEST.yml&title=PacReq%3A+%3Cname_of_the_package%3E'
                                    target='_blank'
                                    color='pink.400'
                                    textDecoration='underline'
                                    fontWeight='800'
                                >
                                    create-request
                                </chakra.a>
                            </Trans>
                        </Text>
                    </Center>
                </Box>
            )}
        </Box>
    )
}

export default PackageList
