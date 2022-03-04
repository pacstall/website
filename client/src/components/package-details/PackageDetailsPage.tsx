import { Container, UseDisclosureProps } from "@chakra-ui/react"
import { FC } from "react"
import PackageInfo from "../../types/package-info"
import ComponentLoader from "../ComponentLoader"
import Navigation from "../Navigation"
import HowToInstall from "./HowToInstall"
import PackageDependenciesModal from "./PackageDependenciesModal"
import PackageDetailsComments from "./PackageDetailsComments"
import PackageDetailsHeader from "./PackageDetailsHeader"
import PackageDetailsTable from "./PackageDetailsTable"
import PackageRequiredByModal from "./PackageRequiredByModal"

type PackageDetailsPageProps = {
    data: PackageInfo;
    allDependencies: string[];
    isMobile: boolean;
    requiredByModal: UseDisclosureProps;
    dependenciesModal: UseDisclosureProps;
}

const PackageDetailsPage: FC<PackageDetailsPageProps> = ({ allDependencies, data, isMobile, dependenciesModal, requiredByModal }) => (
    <>
        <Navigation />
        <Container maxW='900px' mt='10'>
            <PackageDetailsHeader data={data} isMobile={isMobile} />
            <PackageDetailsTable data={data} dependencyCount={allDependencies.length} dependenciesModal={dependenciesModal} requiredByModal={requiredByModal} />
            <HowToInstall name={data.name} isMobile={isMobile} />
            <PackageDetailsComments />
        </Container>
        <PackageRequiredByModal name={data.name} {...requiredByModal} />
        <PackageDependenciesModal name={data.name} {...dependenciesModal} />
    </>
)

export default PackageDetailsPage