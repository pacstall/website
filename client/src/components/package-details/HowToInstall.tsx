import { Box, Heading, HStack, Stack, Text } from '@chakra-ui/react'
import { FC } from 'react'
import OneLineCodeSnippet, {
    SmartCodeSnippetInstall,
} from '../OneLineCodeSnippet'
import InstallNowButton from './InstallNowButton'
import { useTranslation } from 'react-i18next'

const HowToInstallFull: FC<{ name: string; isMobile: boolean }> = ({
    name,
    isMobile,
}) => {
    const ResponsiveStack = isMobile ? Stack : HStack
    const { t } = useTranslation()

    return (
        <Box mt='10'>
            <Heading size='lg'>
                {t('packageDetails.howToInstall.title')}
            </Heading>

            <ResponsiveStack justify='space-between'>
                <Text fontWeight='semibold' m='3'>
                    {t('packageDetails.howToInstall.step1')}
                </Text>
                <Box>
                    <OneLineCodeSnippet size='sm'>
                        sudo bash -c "$(wget -q https://pacstall.dev/q/install
                        -O -)"
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
    const { t } = useTranslation()

    return (
        <Box mt='10'>
            <Heading size='lg'>
                {t('packageDetails.howToInstall.title')}
            </Heading>

            <ResponsiveStack justify='space-between'>
                <Text fontWeight='semibold' m='3'>
                    {t('packageDetails.howToInstall.step1')}
                </Text>
                <Box>
                    <OneLineCodeSnippet size='sm'>
                        sudo bash -c "$(wget -q https://pacstall.dev/q/install
                        -O -)"
                    </OneLineCodeSnippet>
                </Box>
            </ResponsiveStack>

            <ResponsiveStack justify='space-between'>
                <Text fontWeight='semibold' m='3'>
                    {t('packageDetails.howToInstall.step2', {
                        name: prettyName,
                    })}
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
    return <HowToInstallViaTerminal {...props} />
}

export default HowToInstall
