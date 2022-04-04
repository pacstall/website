import { chakra, useColorModeValue } from "@chakra-ui/react";
import { FC } from "react";
import { UpdateStatus } from "../types/package-info";

const SemanticVersionColor: FC<{ version: string; status: UpdateStatus, fill?: boolean, git?: boolean }> = ({ version, status, fill, git }) => {
    const versionColors: Record<UpdateStatus, string> = {
        [UpdateStatus.Unknown]: useColorModeValue('blue.100', 'blue.600'),
        [UpdateStatus.Latest]: useColorModeValue('green.200', 'green.500'),
        [UpdateStatus.Patch]: useColorModeValue('yellow.300', 'yellow.600'),
        [UpdateStatus.Minor]: useColorModeValue('orange.300', 'orange.600'),
        [UpdateStatus.Major]: useColorModeValue('red.300', 'red.400'),
    }

    const versionTooltip: Record<UpdateStatus, string> = {
        [UpdateStatus.Unknown]: 'This package is not in the repology registry',
        [UpdateStatus.Latest]: 'This package is the latest version',
        [UpdateStatus.Patch]: 'This package has a patch update available',
        [UpdateStatus.Minor]: 'This package has a minor update available',
        [UpdateStatus.Major]: 'This package has a major update available',
    }

    const tooltip = status !== UpdateStatus.Unknown ? versionTooltip[status] : (git ? 'This package is built from a specific Git commit' : versionTooltip[UpdateStatus.Unknown])

    return <chakra.span
        bg={versionColors[status]}
        p='1'
        px='2'
        borderRadius='lg'
        display={fill ? 'block' : 'inline-block'}
        minW={fill ? 'initial' : '4em'}
        m={0}
        textAlign='center'
        title={tooltip}
        fontWeight='700'
        color={useColorModeValue('black', 'white')}>
        {version}
    </chakra.span>
}

export default SemanticVersionColor