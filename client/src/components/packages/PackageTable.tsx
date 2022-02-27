import { Table, Tbody, Th, Thead, Tr, useMediaQuery } from "@chakra-ui/react"
import { FC } from "react"
import PackageInfo from "../../types/package-info"
import PackageTableRow from "./PackageTableRow"

const PackageTable: FC<{ packages: PackageInfo[], linksDisabled?: boolean }> = ({ packages, linksDisabled }) => {
    const [isMobile] = useMediaQuery('only screen and (max-device-width: 480px)')
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
                    {!isMobile && (<>
                        <Th>
                            Version
                        </Th>
                        <Th>
                            Install
                        </Th>
                    </>)}
                </Tr>
            </Thead>
            <Tbody>
                {packages.map(pkg => <PackageTableRow isMobile={isMobile} key={pkg.name} disabled={linksDisabled} pkg={pkg} />)}
            </Tbody>
        </Table>
    )
}

export default PackageTable