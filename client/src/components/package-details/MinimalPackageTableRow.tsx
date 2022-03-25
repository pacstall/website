import { Link, Td, Text, Tr, useColorModeValue } from "@chakra-ui/react";
import { FC } from "react";
import { Link as Rlink } from "react-router-dom";

const MinimalPackageTableRow: FC<{ pkg: string, external: boolean }> = ({ pkg, external }) => (
    <Tr key={pkg}>
        <Td py='8px'>
            <Text fontSize='md' fontWeight='500'>
                {external
                    ? <>{pkg}</>
                    : <Link as={Rlink} target='_blank' color={useColorModeValue('pink.600', 'pink.400')} to={`/packages/${pkg}`} >{pkg}</Link>
                }
            </Text>
        </Td>
        <Td py='8px' textAlign='right'>
            <Text fontSize='sm'>
                {external ? 'External / APT' : 'Pacstall Repository'}
            </Text>
        </Td>
    </Tr>
)

export default MinimalPackageTableRow