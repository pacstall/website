import { Table, Tbody, Th, Thead, Tr, useBreakpointValue } from "@chakra-ui/react";
import { FC } from "react"
import PackageInfo from "../../types/package-info"
import PackageTableRow from "./PackageTableRow"

const PackageTable: FC<{ packages: PackageInfo[], linksDisabled?: boolean }> = ({ packages, linksDisabled }) => {
    return (
        <Table>
            <Thead>
                <Tr>
                    <Th>
                        Name
                    </Th>
                    <Th>
                        Maintainer
                    </Th>
                    <Th display={useBreakpointValue({ base: 'none', sm: 'table-cell' })}>
                        Version
                    </Th>
                    <Th display={useBreakpointValue({ base: 'none', md: 'table-cell' })}>
                        Install
                    </Th>
                </Tr>
            </Thead>
            <Tbody>
                {packages.map(pkg => <PackageTableRow key={pkg.name} disabled={linksDisabled} pkg={pkg} />)}
            </Tbody>
        </Table>
    )
}

export default PackageTable