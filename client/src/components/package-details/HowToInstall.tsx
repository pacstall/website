import { Box, Heading, HStack, Stack, Text } from '@chakra-ui/react'
import { FC } from 'react'
import { useFeatureFlag } from '../../state/feature-flags'
import OneLineCodeSnippet, {
    SmartCodeSnippetInstall,
} from '../OneLineCodeSnippet'
import InstallNowButton from './InstallNowButton'

const HowToInstallFull: FC<{ name: string; isMobile: boolean }> = ({
    name,
    isMobile,
}) => {
    const ResponsiveStack = isMobile ? Stack : HStack

    return (
        <Box mt='10'>
            <Heading size='lg'>How to Install</Heading>

            <ResponsiveStack justify='space-between'>
                <Text fontWeight='semibold' m='3'>
                    Step 1: Install Pacstall
                </Text>
                <Box>
                    <OneLineCodeSnippet size='sm'>
                        sudo bash -c "$(wget -q https://git.io/JsADh -O -)"
                    </OneLineCodeSnippet>
                </Box>
            </ResponsiveStack>

            <ResponsiveStack justify='space-between'>
                <Text fontWeight='semibold' m='3'>
                    Step 2a: Enable Browser Integration
                </Text>
                <Box>
                    <OneLineCodeSnippet size='sm'>
                        sudo pacstall enable browser-integration
                    </OneLineCodeSnippet>
                </Box>
            </ResponsiveStack>

            <ResponsiveStack justify='space-between'>
                <Text fontWeight='semibold' m='3'>
                    Step 3: Click on Install Now
                </Text>
                <InstallNowButton disabled={isMobile} />
            </ResponsiveStack>

            <ResponsiveStack justify='space-between'>
                <Text fontWeight='semibold' m='3'>
                    Step 2b: Alternatively, you use the Terminal
                </Text>
                <Box>
                    <SmartCodeSnippetInstall size='sm' name={name} />
                </Box>
            </ResponsiveStack>
        </Box>
    )
}

const HowToInstallViaTerminal: FC<{
    name: string
    prettyName: string
    isMobile: boolean
}> = ({ name, isMobile, prettyName }) => {
    const ResponsiveStack = isMobile ? Stack : HStack

    return (
        <Box mt='10'>
            <Heading size='lg'>How to Install</Heading>

            <ResponsiveStack justify='space-between'>
                <Text fontWeight='semibold' m='3'>
                    Step 1: Setup Pacstall
                </Text>
                <Box>
                    <OneLineCodeSnippet size='sm'>
                        sudo bash -c "$(wget -q https://git.io/JsADh -O -)"
                    </OneLineCodeSnippet>
                </Box>
            </ResponsiveStack>

            <ResponsiveStack justify='space-between'>
                <Text fontWeight='semibold' m='3'>
                    Step 2: Install {prettyName}
                </Text>
                <Box>
                    <SmartCodeSnippetInstall size='sm' name={name} />
                </Box>
            </ResponsiveStack>
        </Box>
    )
}

const HowToInstall: FC<{
    name: string
    prettyName: string
    isMobile: boolean
}> = props => {
    const installProtocolEnabled = useFeatureFlag(
        flags => flags.packageDetailsPage.installProtocol,
    )
    return installProtocolEnabled ? (
        <HowToInstallFull {...props} />
    ) : (
        <HowToInstallViaTerminal {...props} />
    )
}

export default HowToInstall
