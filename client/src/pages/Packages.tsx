import { Box, Center, Container, Spinner, Stack } from "@chakra-ui/react";
import { FC } from "react";
import { Navigate } from "react-router-dom";
import Navigation from "../components/Navigation";
import PackageList from "../components/packages/PackageList";
import Search from "../components/packages/Search";
import Pagination from "../components/Pagination";
import usePackages from "../hooks/usePackages";
import useRandomPackage from "../hooks/useRandomPackage";
import { PackageInfoPage } from "../types/package-info";
import Helmet from 'react-helmet'

const ComputedPackageList: FC<{ result: PackageInfoPage }> = ({ result }) => (
    <>
        <PackageList {...result} />
        <Box m='10'>
            <Center>
                <Pagination current={result.page} last={result.lastPage} />
            </Center>
        </Box>
    </>
)

const Packages: FC = () => {
    const { data: result, error, loading, onSearch, loaded } = usePackages()

    if (error) {
        return <Navigate to="/packages" />
    }


    const randomPackageName = useRandomPackage(result?.data)

    return <>
        <Helmet>
            <title>Packages - Pacstall</title>
        </Helmet>
        <Navigation />

        <Container maxW='1080px'>
            <Stack mt='10'>
                <Search isLoading={loading} placeholder={randomPackageName} onSearch={onSearch} />
                {loaded ? <ComputedPackageList result={result} />
                    : <Box pt='10' textAlign='center'>
                        <Spinner size='lg' />
                    </Box>}
            </Stack>
        </Container>
    </>
}

export default Packages