import { Table, Tbody, Th, Thead, Tr } from '@chakra-ui/react'
import { FC } from 'react'
import { ArchDistroString } from '../../types/package-info'
import MinimalPackageTableRow from './MinimalPackageTableRow'
import { useTranslation } from 'react-i18next'

const MinimalPackageTable: FC<{
    packages: (ArchDistroString | string)[]
    type: string
}> = ({ packages, type }) => {
    const { t } = useTranslation()
    return (
        <Table variant='simple'>
            <Thead>
                <Tr>
                    <Th>{t('packageDetails.requiredByModal.name')}</Th>
                    <Th textAlign='right'>
                        {t('packageDetails.requiredByModal.provider')}
                    </Th>
                </Tr>
            </Thead>
            <Tbody>
                {packages.map((pkg, i) => (
                    <MinimalPackageTableRow
                        external={
                            type !== 'requiredBy' &&
                            type !== 'pacstallDependencies'
                        }
                        key={(pkg.value || pkg.packageName) + i}
                        pkg={pkg.value || pkg.packageName}
                        description={pkg.description || null}
                    />
                ))}
            </Tbody>
        </Table>
    )
}

export default MinimalPackageTable
