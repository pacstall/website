import {
    Box,
    Center,
    Container,
    Spinner,
    Stack,
    useBreakpointValue,
    Image,
    Heading,
} from '@chakra-ui/react'
import { FC } from 'react'
import { Navigate } from 'react-router-dom'
import Search from '../components/packages/Search'
import Pagination from '../components/Pagination'
import usePackages from '../hooks/usePackages'
import useRandomPackage from '../hooks/useRandomPackage'
import { PackageInfoPage } from '../types/package-info'
import PageAnimation from '../components/animations/PageAnimation'
import PackageGrid from '../components/packages/PackageGrid'
// @ts-ignore:next-line
import PacstallLogo from '../../public/pacstall.svg'

const ComputedPackageList: FC<{ result: PackageInfoPage }> = ({ result }) => (
    <>
        <PackageGrid data={result?.data || []} />
        <Box m='10'>
            <Center>
                <Pagination current={result.page} last={result.lastPage} />
            </Center>
        </Box>
    </>
)

const Pacstore: FC = () => {
    const { data: result, error, loading, onSearch, loaded } = usePackages()
    const randomPackageName = useRandomPackage(result?.data || [])

    if (error) {
        return <Navigate to='/packages' />
    }

    return (
        <>
            <Center>
                <Stack
                    justify='space-between'
                    mt='10'
                    direction={useBreakpointValue({
                        base: 'column',
                        md: 'row',
                    })}
                >
                    <div>
                        <Heading size='2xl' pb='3' color='brand.400'>
                            Pacstall Store
                        </Heading>
                        <Heading size='lg'>Package Discovery for Pacstall</Heading>
                    </div>
                    <Image
                        src={PacstallLogo}
                        width='200px'
                        height='200px'
                        minW='10em'
                        alt='Pacstall logo'
                        mx='auto'
                        my='1em'
                        position='relative'
                        display={useBreakpointValue({
                            base: 'none',
                            md: 'initial',
                        })}
                        bottom='1.75em'
                        loading='lazy'
                    />
                </Stack>
            </Center>
            <PageAnimation>
                <Container maxW='60em'>
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

export default Pacstore
