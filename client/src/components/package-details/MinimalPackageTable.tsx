import { chakra, Table, Tbody, Th, Thead, Tr, useColorModeValue } from "@chakra-ui/react";
import { FC } from "react"
import PackageInfo from "../../types/package-info";
import MinimalPackageTableRow from "./MinimalPackageTableRow"

const MinimalPackageTable: FC<{ packages: (PackageInfo | string)[] }> = ({ packages }) => {
    return (
        <Table variant={'simple'}>
            <Thead>
                <Tr>
                    <Th>
                        Name
                    </Th>
                    <Th textAlign='right'>
                        Provider
                    </Th>
                </Tr>
            </Thead>
            <Tbody>
                {packages.map((pkg, i) => <MinimalPackageTableRow external={typeof pkg === 'string'} key={(typeof pkg === 'string' ? pkg : pkg.name) + i} pkg={typeof pkg === 'string' ? pkg : pkg.name} />)}
            </Tbody>
        </Table>
    )
}

export default MinimalPackageTable