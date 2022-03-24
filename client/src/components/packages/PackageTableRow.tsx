import { Link, Td, Text, Tr, useBreakpointValue, useColorModeValue } from "@chakra-ui/react";
import { FC } from "react";
import { Link as Rlink } from "react-router-dom";
import PackageInfo from "../../types/package-info";
import { SmartCodeSnippetInstall } from "../OneLineCodeSnippet";

const PackageTableRow: FC<{ pkg: PackageInfo, disabled?: boolean }> = ({ pkg, disabled }) => (
    <Tr key={pkg.name}>
        <Td>
            <Text fontSize='md' fontWeight='500'>
                {disabled === true ? <span>{pkg.name}</span> : <Link as={Rlink} color={useColorModeValue('pink.600', 'pink.400')} to={`/packages/${pkg.name}`} >{pkg.name}</Link>}
            </Text>
        </Td>
        <Td>
            <Text fontSize='sm'>
                {(pkg.maintainer || 'Orphaned').split('<')[0].trim()}
            </Text>

        </Td>
        <Td display={useBreakpointValue({ base: 'none', sm: 'table-cell' })}>
            <Text fontSize='sm'>
                {pkg.version.substring(0, 14)}
            </Text>

        </Td>
        <Td display={useBreakpointValue({ base: 'none', md: 'table-cell' })}>
            <Text fontSize='sm'>
                <SmartCodeSnippetInstall size="sm" name={pkg.name} />
            </Text>
        </Td>
    </Tr>
)

export default PackageTableRow