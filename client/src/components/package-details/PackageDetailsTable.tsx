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
import useNumericDisplay from '../../hooks/useNumericDisplay'

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

const UNKNOWN_DATE_SENTINEL = '0001-01-01T00:00:00Z'

const PackageDetailsTable: FC<{
    data: PackageInfo
    dependencyCount: number
    requiredByModal: UseDisclosureProps
    dependenciesModal: UseDisclosureProps
}> = ({ data, dependencyCount, requiredByModal, dependenciesModal }) => {
    const { t } = useTranslation()
    const displayNumber = useNumericDisplay()
    return (
        <Table mt='10'>
            <Tbody>
                <Entry header={t('packageDetails.table.name')}>
                    {data.packageName}
                </Entry>

                <Entry header={t('packageDetails.table.version')}>
                    <Text fontWeight='bold'>
                        <SemanticVersionColor
                            git={data.packageName.endsWith('-git')}
                            version={data.version}
                            status={data.updateStatus}
                        />
                    </Text>
                </Entry>

                <Entry header={t('packageDetails.table.maintainer')}>
                    <PackageDetailsMaintainer text={data.maintainer} />
                </Entry>

                <Entry header={t('packageDetails.table.lastUpdatedAt')}>
                    {data.lastUpdatedAt === UNKNOWN_DATE_SENTINEL
                        ? '-'
                        : new Intl.DateTimeFormat().format(
                              new Date(data.lastUpdatedAt),
                          )}
                </Entry>

                <Entry header={t('packageDetails.table.dependencies')}>
                    {dependencyCount
                        ? displayNumber(dependencyCount)
                        : t('packageDetails.noResults')}{' '}
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
                    {data.requiredBy.length
                        ? displayNumber(data.requiredBy.length)
                        : t('packageDetails.noResults')}{' '}
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
                        href={`https://github.com/pacstall/pacstall-programs/blob/master/packages/${data.packageName}/${data.packageName}.pacscript`}
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
