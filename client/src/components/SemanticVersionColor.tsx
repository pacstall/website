import { chakra, Tooltip, useColorModeValue } from '@chakra-ui/react'
import { FC } from 'react'
import { UpdateStatus } from '../types/package-info'
import { useTranslation } from 'react-i18next'

const SemanticVersionColor: FC<{
    version: string
    status: UpdateStatus
    fill?: boolean
    git?: boolean
}> = ({ version, status, fill, git }) => {
    const { t } = useTranslation()

    const versionColors: Record<UpdateStatus, string> = {
        [UpdateStatus.Unknown]: useColorModeValue('blue.100', 'blue.600'),
        [UpdateStatus.Latest]: useColorModeValue('green.200', 'green.500'),
        [UpdateStatus.Patch]: useColorModeValue('yellow.300', 'yellow.600'),
        [UpdateStatus.Minor]: useColorModeValue('orange.300', 'orange.600'),
        [UpdateStatus.Major]: useColorModeValue('red.300', 'red.400'),
    }

    const versionTooltip: Record<UpdateStatus, string> = {
        [UpdateStatus.Unknown]: 'packageSearch.versionTooltip.notInRegistry',
        [UpdateStatus.Latest]: 'packageSearch.versionTooltip.latest',
        [UpdateStatus.Patch]: 'packageSearch.versionTooltip.patch',
        [UpdateStatus.Minor]: 'packageSearch.versionTooltip.minor',
        [UpdateStatus.Major]: 'packageSearch.versionTooltip.major',
    }

    const tooltip =
        status !== UpdateStatus.Unknown
            ? versionTooltip[status]
            : git
              ? 'packageSearch.versionTooltip.isGit'
              : versionTooltip[UpdateStatus.Unknown]

    return (
        <Tooltip openDelay={500} label={t(tooltip)}>
            <chakra.span
                bg={versionColors[status]}
                p='1'
                px='2'
                borderRadius='lg'
                display={fill ? 'block' : 'inline-block'}
                minW={fill ? 'initial' : '4em'}
                m={0}
                textAlign='center'
                fontWeight='700'
                color={useColorModeValue('black', 'white')}
            >
                {version}
            </chakra.span>
        </Tooltip>
    )
}

export default SemanticVersionColor
