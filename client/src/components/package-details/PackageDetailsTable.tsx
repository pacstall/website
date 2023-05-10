import { ExternalLinkIcon } from '@chakra-ui/icons'
import {
    Table,
    Tbody,
    Link,
    Tr,
    Th,
    Td,
    Text,
    Icon,
    UseDisclosureProps,
} from '@chakra-ui/react'
import { FC, ReactNode } from 'react'
import { Link as Rlink } from 'react-router-dom'
import PackageInfo from '../../types/package-info'
import SemanticVersionColor from '../SemanticVersionColor'
import PackageDetailsMaintainer from './PackageDetailsMaintainer'
import { useTranslation } from 'react-i18next'

const Entry: FC<{
    header: string
    disabled?: boolean
    children: ReactNode
}> = ({ header, children, disabled }) => (
    <>
        {!disabled && (
            <Tr>
                <Th>{header}</Th>
                <Td>{children}</Td>
            </Tr>
        )}
    </>
)

const PackageDetailsTable: FC<{
    data: PackageInfo
    dependencyCount: number
    requiredByModal: UseDisclosureProps
    dependenciesModal: UseDisclosureProps
}> = ({ data, dependencyCount, requiredByModal, dependenciesModal }) => {
    const { t } = useTranslation()
    return (
        <Table mt='10'>
            <Tbody>
                <Entry header={t('packageDetails.table.name')}>
                    {data.name}
                </Entry>

                <Entry header={t('packageDetails.table.version')}>
                    <Text fontWeight='bold'>
                        <SemanticVersionColor
                            git={data.name.endsWith('-git')}
                            version={data.version}
                            status={data.updateStatus}
                        />
                    </Text>
                </Entry>

                <Entry header={t('packageDetails.table.maintainer')}>
                    <PackageDetailsMaintainer text={data.maintainer} />
                </Entry>

                <Entry header={t('packageDetails.table.dependencies')}>
                    {dependencyCount || t('packageDetails.noResults')}{' '}
                    {dependencyCount > 0 ? (
                        <Link
                            onClick={dependenciesModal.onOpen}
                            pl='2'
                            color='pink.400'
                            as={Rlink}
                            to={`#`}
                        >
                            {t('packageDetails.view')}
                        </Link>
                    ) : (
                        ''
                    )}
                </Entry>

                <Entry header={t('packageDetails.table.requiredBy')}>
                    {data.requiredBy.length || t('packageDetails.noResults')}{' '}
                    {data.requiredBy?.length || 0 > 0 ? (
                        <Link
                            onClick={requiredByModal.onOpen}
                            pl='2'
                            color='pink.400'
                            as={Rlink}
                            to={`#`}
                        >
                            {t('packageDetails.view')}
                        </Link>
                    ) : (
                        ''
                    )}
                </Entry>

                <Entry header='Pacscript'>
                    <Link
                        color='pink.400'
                        isExternal
                        href={`https://github.com/pacstall/pacstall-programs/blob/master/packages/${data.name}/${data.name}.pacscript`}
                    >
                        {t('packageDetails.openInGithub')}{' '}
                        <Icon
                            position='relative'
                            bottom='2px'
                            size='md'
                            ml='1px'
                            as={ExternalLinkIcon}
                        />
                    </Link>
                </Entry>
            </Tbody>
        </Table>
    )
}

export default PackageDetailsTable
