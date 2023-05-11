import { Box, Center, Container, Spinner, Stack } from '@chakra-ui/react'
import { FC } from 'react'
import { Navigate } from 'react-router-dom'
import PackageList from '../components/packages/PackageList'
import Search from '../components/packages/Search'
import Pagination from '../components/Pagination'
import usePackages from '../hooks/usePackages'
import useRandomPackage from '../hooks/useRandomPackage'
import { PackageInfoPage } from '../types/package-info'
import Helmet from 'react-helmet'
import PageAnimation from '../components/animations/PageAnimation'

const ComputedPackageList: FC<{ result: PackageInfoPage }> = ({ result }) => (
    <>
        <PackageList data={result?.data || []} />
        <Box m='10px' mt='20px !important'>
            <Center>
                <Pagination current={result.page} last={result.lastPage} />
            </Center>
        </Box>
    </>
)

const Packages: FC = () => {
    const { data: result, error, loading, onSearch, loaded } = usePackages()
    const randomPackageName = useRandomPackage(result?.data || [])

    if (error) {
        return <Navigate to='/packages' />
    }

    return (
        <>
            <Helmet>
                <title>Packages - Pacstall</title>
            </Helmet>

            <PageAnimation>
                <Container maxW='70em'>
                    <Stack my='10'>
                        <Search
                            isLoading={loading}
                            placeholder={randomPackageName}
                            onSearch={onSearch}
                        />
                        {loaded ? (
                            <ComputedPackageList result={result} />
                        ) : (
                            <Box pt='10' textAlign='center'>
                                <Spinner size='lg' />
                            </Box>
                        )}
                    </Stack>
                </Container>
            </PageAnimation>
        </>
    )
}

export default Packages
