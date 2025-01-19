import {
    Container,
    Heading,
    Image,
    Link,
    Stack,
    Text,
    useBreakpointValue,
    useColorMode,
} from '@chakra-ui/react'
import { FC } from 'react'
import Card from '../components/Card'
import OneLineCodeSnippet from '../components/OneLineCodeSnippet'

import { Helmet } from 'react-helmet'
import AsciinemaFrame from '../components/AsciinemaFrame'
import PageAnimation from '../components/animations/PageAnimation'
import { Trans, useTranslation } from 'react-i18next'

const Home: FC = () => {
    const { t } = useTranslation()
    const { colorMode } = useColorMode()
    return (
        <>
            <Helmet>
                <title>Pacstall - The AUR for Ubuntu</title>
            </Helmet>

            <PageAnimation>
                <Container maxW='60em'>
                    <Stack
                        justify='space-between'
                        mt='7'
                        direction={useBreakpointValue({
                            base: 'column',
                            md: 'row',
                        })}
                    >
                        <div>
                            <Heading size='2xl' pb='3' color='brand.400'>
                                {t('home.title')}
                            </Heading>
                            <Heading size='lg'>{t('home.subtitle')}</Heading>
                        </div>
                        <Image
                            src='/pacstall.svg'
                            width='200px'
                            height='200px'
                            minW='10em'
                            alt='Pacstall logo'
                            mx='auto'
                            my='1em'
                            position='relative'
                            display={useBreakpointValue({
                                base: 'none',
                                md: 'initial',
                            })}
                            bottom='1.75em'
                            loading='lazy'
                        />
                    </Stack>

                    <Stack
                        direction={useBreakpointValue({
                            base: 'column',
                            md: 'row',
                        })}
                    >
                        <Stack maxW='2xl'>
                            <Card title={t('home.cards.whyDifferent.title')}>
                                <Text maxW='65ch'>
                                    {t('home.cards.whyDifferent.description')}
                                </Text>
                            </Card>
                        </Stack>

                        <Stack maxW='2xl'>
                            <Card title={t('home.cards.howItWorks.title')}>
                                <Text maxW='65ch'>
                                    <Trans i18nKey='home.cards.howItWorks.description'>
                                        <Link
                                            color='blue.400'
                                            href='https://github.com/pacstall/pacstall/wiki/Pacscript-101'
                                        >
                                            pacscripts
                                        </Link>
                                    </Trans>
                                </Text>
                            </Card>
                        </Stack>
                    </Stack>

                    <Heading size={'lg'} mb='3' mt='10'>
                        {t('home.installationInstructions')}
                    </Heading>
                    <OneLineCodeSnippet>
                        sudo bash -c "$(curl -fsSL
                        https://pacstall.dev/q/install)"
                    </OneLineCodeSnippet>
                    <Heading size={'lg'} mb='3' mt='10'>
                        {t('home.showcase.title')}
                    </Heading>
                    <AsciinemaFrame
                        autoPlay={true}
                        loop={true}
                        src='/showcase.cast'
                        preload={true}
                        terminalFontFamily='FiraCodeNF'
                        terminalfontSize='13'
                        theme={colorMode === 'light' ? 'glitter' : 'sparkle'}
                        rows='32'
                    />

                    <Heading size={'lg'} mb='3'>
                        {t('home.showcase.packageSearch')}
                    </Heading>
                    <AsciinemaFrame
                        autoPlay={true}
                        loop={true}
                        src='/search.cast'
                        preload={true}
                        terminalFontFamily='FiraCodeNF'
                        terminalfontSize='13'
                        theme={colorMode === 'light' ? 'glitter' : 'sparkle'}
                        rows='32'
                    />

                    <Stack justify='center'>
                        <Image
                            src='/pacstall.svg'
                            width='200px'
                            height='200px'
                            alt='Pacstall logo'
                            mx='auto'
                            mb='5'
                            display={useBreakpointValue({
                                base: 'initial',
                                md: 'none',
                            })}
                            loading='lazy'
                        />
                    </Stack>
                </Container>
            </PageAnimation>
        </>
    )
}

export default Home
