import {
    Container,
    Flex,
    Heading,
    Image,
    Link,
    Stack,
    Text,
    useBreakpointValue,
} from '@chakra-ui/react'
import { FC } from 'react'
import Card from '../components/Card'
import OneLineCodeSnippet, {
    SmartCodeSnippetInstall,
} from '../components/OneLineCodeSnippet'

// @ts-ignore:next-line
import PacstallLogo from '../../public/pacstall.svg'
import { Helmet } from 'react-helmet'
import AsciinemaFrame from '../components/AsciinemaFrame'
import { PrimaryButton } from '../components/Button'
import PageAnimation from '../components/animations/PageAnimation'

const Home: FC = () => {
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
                                Pacstall
                            </Heading>
                            <Heading size='lg'>The AUR for Ubuntu</Heading>
                        </div>
                        <Image
                            src={PacstallLogo}
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
                            <Card title='Why is this any different than any other package manager?'>
                                <Text maxW='65ch'>
                                    Pacstall uses the stable base of Ubuntu but
                                    allows you to use bleeding edge software
                                    with little to no compromises, so you don't
                                    have to worry about security patches or new
                                    features.
                                </Text>
                            </Card>
                        </Stack>

                        <Stack maxW='2xl'>
                            <Card title='How does it work then?'>
                                <Text maxW='65ch'>
                                    Pacstall takes in files known as{' '}
                                    <Link
                                        color='blue.400'
                                        href='https://github.com/pacstall/pacstall/wiki/Pacscript-101'
                                    >
                                        pacscripts
                                    </Link>{' '}
                                    (similar to PKGBUILD's) that contain the
                                    necessary contents to build packages, and
                                    builds them into executables on your system.
                                </Text>
                            </Card>
                        </Stack>
                    </Stack>

                    <Heading size={'lg'} mb='3' mt='10'>
                        Installation Instructions
                    </Heading>
                    <OneLineCodeSnippet>
                        sudo bash -c "$(curl -fsSL https://git.io/JsADh || wget
                        -q https://git.io/JsADh -O -)"
                    </OneLineCodeSnippet>
                    <Heading size={'lg'} mb='3' mt='10'>
                        Showcase
                    </Heading>
                    <AsciinemaFrame autoplay loop id='459473' />

                    <Heading size={'lg'} mb='3'>
                        Package search
                    </Heading>
                    <AsciinemaFrame id='459474' />

                    <Stack justify='center'>
                        <Image
                            src={PacstallLogo}
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
