import { Box, Center, Container, Stack } from "@chakra-ui/react";
import { FC } from "react";
import { Navigate } from "react-router-dom";
import ComponentLoader from "../components/ComponentLoader";
import Navigation from "../components/Navigation";
import PackageList from "../components/packages/PackageList";
import Search from "../components/packages/Search";
import Pagination from "../components/Pagination";
import usePackages from "../hooks/usePackages";
import useRandomPackage from "../hooks/useRandomPackage";

const Packages: FC = () => {
    const { data: result, error, loading, onSearch, loaded } = usePackages()

    if (error) {
        return <Navigate to="/packages" />
    }

    const ComputedPackageList = () => (
        <>
            <PackageList {...result} />
            <Box m='10'>
                <Center>
                    <Pagination current={result.page} last={result.lastPage} />
                </Center>
            </Box>
        </>
    )
    const randomPackageName = useRandomPackage(result?.data)

    return <>
        <Navigation />

        <Container maxW='1080px'>
            <Stack mt='10'>
                <Search isLoading={loading} placeholder={randomPackageName} onSearch={onSearch} />
                <ComponentLoader isLoading={loading} hasError={error} content={ComputedPackageList} />
            </Stack>
        </Container>
    </>
}

export default Packages