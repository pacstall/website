import { Table, Tbody, Th, Thead, Tr } from '@chakra-ui/react'
import { FC } from 'react'
import PackageInfo from '../../types/package-info'
import MinimalPackageTableRow from './MinimalPackageTableRow'
import { useTranslation } from 'react-i18next'

const MinimalPackageTable: FC<{ packages: (PackageInfo | string)[] }> = ({
    packages,
}) => {
    const { t } = useTranslation()
    return (
        <Table variant={'simple'}>
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
                        external={typeof pkg === 'string'}
                        key={
                            (typeof pkg === 'string' ? pkg : pkg.packageName) +
                            i
                        }
                        pkg={typeof pkg === 'string' ? pkg : pkg.packageName}
                    />
                ))}
            </Tbody>
        </Table>
    )
}

export default MinimalPackageTable
