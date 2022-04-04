import { Link, Td, Text, Tr, useBreakpointValue, useColorModeValue } from "@chakra-ui/react";
import { FC } from "react";
import { Link as Rlink } from "react-router-dom";
import PackageInfo from "../../types/package-info";
import { SmartCodeSnippetInstall } from "../OneLineCodeSnippet";
import SemanticVersionColor from "../SemanticVersionColor";

const PackageTableRow: FC<{ pkg: PackageInfo, disabled?: boolean }> = ({ pkg, disabled }) => (
    <Tr key={pkg.name}>
        <Td>
            <Text fontSize='md' fontWeight={useColorModeValue('700', '500')} title={pkg.description}>
                {disabled === true ? <span>{pkg.name}</span> : <Link as={Rlink} color={useColorModeValue('pink.600', 'pink.400')} to={`/packages/${pkg.name}`} >{pkg.name}</Link>}
            </Text>
        </Td>
        <Td>
            <Text fontSize='sm' title={pkg.maintainer ? `This package is being maintained by ${pkg.maintainer.split('<')[0].trim()}` : 'This package does not have a maintainer'}>
                {(pkg.maintainer || 'Orphaned').split('<')[0].trim()}
            </Text>

        </Td>
        <Td display={useBreakpointValue({ base: 'none', sm: 'table-cell' })}>
            <Text fontSize='sm'>
                <SemanticVersionColor git={pkg.name.endsWith("-git")} fill version={pkg.version.substring(0, 14)} status={pkg.updateStatus} />
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