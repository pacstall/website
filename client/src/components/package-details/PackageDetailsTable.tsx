import { ExternalLinkIcon } from "@chakra-ui/icons";
import { Table, Tbody, Link, Tr, Th, Td, useColorModeValue, Text, Icon, UseDisclosureProps } from "@chakra-ui/react";
import { FC } from "react";
import { Link as Rlink } from 'react-router-dom'
import { useFeatureFlag } from "../../state/feature-flags";
import PackageInfo, { UpdateStatus } from "../../types/package-info";
import SemanticVersionColor from "../SemanticVersionColor";
import PackageDetailsMaintainer from "./PackageDetailsMaintainer";

const Entry: FC<{ header: string, disabled?: boolean }> = ({ header, children, disabled }) => (
    <>
        {!disabled && (
            <Tr>
                <Th>{header}</Th>
                <Td>{children}</Td>
            </Tr>
        )}
    </>
)

const PackageDetailsTable: FC<{ data: PackageInfo, dependencyCount: number, requiredByModal: UseDisclosureProps, dependenciesModal: UseDisclosureProps }> = ({ data, dependencyCount, requiredByModal, dependenciesModal }) => {
    const lastUpdatedDisabled = !useFeatureFlag(flag => flag.packageDetailsPage.lastUpdated)
    const popularityDisabled = !useFeatureFlag(flag => flag.packageDetailsPage.popularity)
    const votesDisabled = !useFeatureFlag(flag => flag.packageDetailsPage.votes)

    return (
        <Table mt='10'>
            <Tbody>
                <Entry header="Name">
                    {data.name}
                </Entry>

                <Entry header="Version">
                    <Text
                        fontWeight='bold'>
                        <SemanticVersionColor version={data.version} status={data.updateStatus} />
                    </Text>
                </Entry>

                <Entry header="Maintainer">
                    <PackageDetailsMaintainer text={data.maintainer} />
                </Entry>

                <Entry header="Last Updated" disabled={lastUpdatedDisabled}>
                    Today
                </Entry>

                <Entry header="Votes" disabled={votesDisabled}>
                    {Math.floor(Math.random() * 1200) + 30}
                </Entry>

                <Entry header="Popularity" disabled={popularityDisabled}>
                    {Math.floor(Math.random() * 1000) / 10}%
                </Entry>

                <Entry header="Dependencies">
                    {dependencyCount || 'None'} {dependencyCount > 0
                        ? <Link onClick={dependenciesModal.onOpen} pl='2' color='pink.400' as={Rlink} to={`#`}>View</Link>
                        : ''
                    }
                </Entry>

                <Entry header="Required By">
                    {data.requiredBy.length || 'None'} {data.requiredBy?.length || 0 > 0
                        ? <Link onClick={requiredByModal.onOpen} pl='2' color='pink.400' as={Rlink} to={`#`}>View</Link>
                        : ''
                    }
                </Entry>

                <Entry header="Pacscript">
                    <Link
                        color='pink.400'
                        isExternal
                        href={`https://github.com/pacstall/pacstall-programs/blob/master/packages/${data.name}/${data.name}.pacscript`}>
                        Open in GitHub <Icon position='relative' bottom='2px' size='md' ml='1px' as={ExternalLinkIcon} />
                    </Link>
                </Entry>
            </Tbody>
        </Table>
    )
}

export default PackageDetailsTable