import { Link, Td, Text, Tr, useColorModeValue } from "@chakra-ui/react";
import { FC } from "react";
import { Link as Rlink } from "react-router-dom";
import PackageInfo from "../../types/package-info";
import { SmartCodeSnippetInstall } from "../OneLineCodeSnippet";

const PackageTableRow: FC<{ pkg: PackageInfo, disabled?: boolean, isMobile: boolean }> = ({ pkg, disabled, isMobile }) => (
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
        {!isMobile && (<>
            <Td>
                <Text fontSize='sm'>
                    {pkg.version.substring(0, 14)}
                </Text>

            </Td>
            <Td>
                <Text fontSize='sm'>
                    <SmartCodeSnippetInstall size="sm" name={pkg.name} />
                </Text>
            </Td>
        </>)}
    </Tr>
)

export default PackageTableRow