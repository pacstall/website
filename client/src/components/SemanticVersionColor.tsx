import { chakra, useColorModeValue } from "@chakra-ui/react";
import { FC } from "react";
import { UpdateStatus } from "../types/package-info";

const SemanticVersionColor: FC<{ version: string; status: UpdateStatus }> = ({ version, status }) => {
    const versionColors: Record<UpdateStatus, string> = {
        [UpdateStatus.Unknown]: useColorModeValue('gray.100', 'gray.700'),
        [UpdateStatus.Latest]: useColorModeValue('green.200', 'green.500'),
        [UpdateStatus.Patch]: useColorModeValue('yellow.300', 'yellow.600'),
        [UpdateStatus.Minor]: useColorModeValue('orange.300', 'orange.600'),
        [UpdateStatus.Major]: useColorModeValue('red.400', 'red.600'),
    }

    const versionTooltip: Record<UpdateStatus, string> = {
        [UpdateStatus.Unknown]: 'This package is not in the repology registry',
        [UpdateStatus.Latest]: 'This package is the latest version',
        [UpdateStatus.Patch]: 'This package has a patch update available',
        [UpdateStatus.Minor]: 'This package has a minor update available',
        [UpdateStatus.Major]: 'This package has a major update available',
    }

    return <chakra.span
        bg={versionColors[status]}
        p='1'
        px='2'
        borderRadius='lg'
        title={versionTooltip[status]}
        color={useColorModeValue('black', 'white')}>
        {version}
    </chakra.span>
}

export default SemanticVersionColor