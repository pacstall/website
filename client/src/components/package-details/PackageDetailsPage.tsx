import { Container, UseDisclosureProps } from '@chakra-ui/react'
import { FC } from 'react'
import { Helmet } from 'react-helmet'
import PackageInfo, { ArchDistroString } from '../../types/package-info'
import HowToInstall from './HowToInstall'
import PackageDependenciesModal from './PackageDependenciesModal'
import PackageDetailsHeader from './PackageDetailsHeader'
import PackageDetailsTable from './PackageDetailsTable'
import PackageRequiredByModal from './PackageRequiredByModal'

type PackageDetailsPageProps = {
    data: PackageInfo
    allDependencies: ArchDistroString[]
    isMobile: boolean
    requiredByModal: UseDisclosureProps
    dependenciesModal: UseDisclosureProps
}

const PackageDetailsPage: FC<PackageDetailsPageProps> = ({
    allDependencies,
    data,
    isMobile,
    dependenciesModal,
    requiredByModal,
}) => (
    <>
        <Helmet>
            <title>{data.packageName} - Pacstall</title>
            <meta
                name='keywords'
                content={
                    data.packageName +
                    ',' +
                    data.packageName.split('-').join(',')
                }
            />
            <meta name='description' content={data.description} />

            <meta name='twitter:card' content='summary' />
            <meta property='og:title' content={data.packageName} />
            <meta property='og:type' content='article' />
            <meta property='og:url' content={location.href} />
            <meta property='og:image' content='/public/app.png' />
            <meta property='og:description' content={data.description} />
        </Helmet>
        <Container maxW='60em' mt='10'>
            <PackageDetailsHeader data={data} isMobile={isMobile} />
            <PackageDetailsTable
                data={data}
                dependencyCount={allDependencies.length}
                dependenciesModal={dependenciesModal}
                requiredByModal={requiredByModal}
            />
            <HowToInstall
                name={
                    data.baseTotal > 1
                        ? data.baseIndex === 0
                            ? `${data.packageBase}:pkgbase`
                            : `${data.packageBase}:${data.packageName}`
                        : data.packageName
                }
                prettyName={data.packageName}
                isMobile={isMobile}
            />
        </Container>
        <PackageRequiredByModal name={data.packageName} {...requiredByModal} />
        <PackageDependenciesModal
            name={data.packageName}
            {...dependenciesModal}
        />
    </>
)

export default PackageDetailsPage
