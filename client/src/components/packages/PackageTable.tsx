import {
    Table,
    Tbody,
    Th,
    Thead,
    Tr,
    useBreakpointValue,
} from '@chakra-ui/react'
import { FC } from 'react'
import PackageInfo from '../../types/package-info'
import PackageTableRow from './PackageTableRow'
import { useTranslation } from 'react-i18next'

const PackageTable: FC<{
    packages: PackageInfo[]
    linksDisabled?: boolean
}> = ({ packages, linksDisabled }) => {
    const { t } = useTranslation()
    return (
        <Table animation='ease-in 100ms'>
            <Thead>
                <Tr>
                    <Th>{t('packageSearch.table.name')}</Th>
                    <Th>{t('packageSearch.table.maintainer')}</Th>
                    <Th
                        display={useBreakpointValue({
                            base: 'none',
                            sm: 'table-cell',
                        })}
                    >
                        {t('packageSearch.table.version')}
                    </Th>
                    <Th
                        display={useBreakpointValue({
                            base: 'none',
                            md: 'table-cell',
                        })}
                    >
                        {t('packageSearch.table.install')}
                    </Th>
                </Tr>
            </Thead>
            <Tbody>
                {packages.map(pkg => (
                    <PackageTableRow
                        key={pkg.packageName}
                        disabled={linksDisabled}
                        pkg={pkg}
                    />
                ))}
            </Tbody>
        </Table>
    )
}

export default PackageTable
