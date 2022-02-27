import axios from "axios";
import { FC, useEffect, useState } from "react";
import { useParams, Navigate, Link as Rlink } from "react-router-dom";
import Navigation from "../components/Navigation";
import serverConfig from "../config/server";
import PackageInfo from "../types/package-info";
// @ts-ignore:next-line
import DefaultAppImg from "../../public/app.png";
import OneLineCodeSnippet, { SmartCodeSnippetInstall } from "../components/OneLineCodeSnippet";
import { featureFlagsState } from "../state/feature-flags";
import useNotification from "../hooks/useNotification";
import { Container, Table, Tbody, Td, Th, Tr, Link, Heading, HStack, Image, Text, Box, Button, useColorModeValue, Spinner, useMediaQuery, Stack } from "@chakra-ui/react";
import { useRecoilState } from "recoil";

const toTitle = (str: string): string => {
    const parts = str.split('-')

    if (['deb', 'git', 'app', 'bin'].includes(parts[parts.length - 1])) {
        parts.pop()
    }

    return parts
        .map(part => part[0].toUpperCase() + part.substring(1))
        .join(' ')

}

const Maintainer: FC<{ text: string }> = ({ text }) => {
    if (!text || text === '-' || text.toLowerCase() === 'orphan' || text.toLowerCase() === 'orphaned') {
        return <>Orphaned</>
    }

    if (!['<', '>', '@'].every(symbol => text.includes(symbol))) {
        return <>{text}</>
    }

    const shortenName = (name: string, splitBy: string): string =>
        name.split(splitBy).reduce((acc, part) => (acc + part).length > 14 ? acc : acc + splitBy + part, '')

    let name = text.split('<')[0].trim()
    if (name.length > 15) {
        if (name.includes(' ')) {
            name = shortenName(name, ' ')
        } else if (name.includes('-')) {
            name = shortenName(name, '-')
        } else {
            name = name.substring(0, 12) + '..'
        }
    }

    const fullEmail = text.split('<')[1].split('>')[0].trim()
    const shortEmail = fullEmail.split('@')[0].length > 15 ? fullEmail.split('@')[0].substring(0, 13) + '[..]@' + fullEmail.split('@')[1] : fullEmail

    return (
        <>
            <span>{name}, </span>
            <Link color='pink.400' href={"mailto: " + fullEmail}>
                {shortEmail}
            </Link>
        </>
    )
}

const InstallNowButton: FC<{ disabled?: boolean }> = ({ disabled }) => (
    <Button
        color={useColorModeValue('black', 'white')}
        bg={useColorModeValue('brand.200', 'brand.500')}
        _hover={{ bg: useColorModeValue('brand.200', 'brand.300') }}
        _active={{ bg: useColorModeValue('brand.200', 'brand.300') }}
        disabled={disabled}
        size='lg'>
        Install Now
    </Button>
)

const PackageDetails: FC = () => {
    const [featureFlags] = useRecoilState(featureFlagsState)
    const [pageDisabled, setPageDisabled] = useState<boolean>()
    const [isMobile] = useMediaQuery('only screen and (max-device-width: 480px)')

    useEffect(() => {
        setPageDisabled(
            featureFlags.loading
                ? undefined
                : featureFlags.error
                    ? false
                    : featureFlags.flags!.packageDetailsPageDisabled
        )
    }, [featureFlags])

    const notify = useNotification()

    const { name } = useParams() as { name: string }
    const [pkgInfo, setPkgInfo] = useState<PackageInfo>()
    const [loading, setLoading] = useState(true)

    useEffect(() => {
        setLoading(true)
        axios.get<PackageInfo>(serverConfig.host + `/api/packages/${name}`)
            .then(result => setPkgInfo(result.data))
            .catch(() => { })
            .then(() => setLoading(false))
    }, [name])

    if (pageDisabled) {
        notify({
            title: 'This feature is not ready yet.',
            text: 'You are being redirected back to the home page.',
            type: 'warning'
        })
        return <Navigate to="/" />
    }

    if (!loading && !pkgInfo) {
        return <Navigate to="/not-found" />
    }

    const ResponsiveStack = isMobile ? Stack : HStack;

    const allDependencies = [...pkgInfo?.buildDependencies || [], ...pkgInfo?.runtimeDependencies || [], ...pkgInfo?.optionalDependencies || [], ...pkgInfo?.pacstallDependencies || []]

    return (<>
        <Navigation />

        {!loading && (
            <Container maxW='900px' mt='10'>
                <HStack justify='space-between'>
                    <HStack>
                        <Image src={DefaultAppImg} maxW='64px' />
                        <Heading>{toTitle(name)}</Heading>
                    </HStack>
                    {!isMobile && <InstallNowButton />}
                </HStack>

                <Text mt='5'>
                    {pkgInfo?.description}
                </Text>

                <Table mt='10'>
                    <Tbody>
                        <Tr>
                            <Th>Name</Th>
                            <Td>{pkgInfo?.name}</Td>
                        </Tr>
                        <Tr>
                            <Th>Version</Th>
                            <Td>
                                <Text
                                    fontWeight='bold'
                                    color={useColorModeValue('black', 'white')}
                                    as='span'
                                    p='1'
                                    px='2'
                                    borderRadius='lg'
                                    bg={useColorModeValue('gray.100', 'gray.700')}>
                                    {pkgInfo?.version}
                                </Text>
                            </Td>
                        </Tr>
                        <Tr>
                            <Th>Maintainer</Th>
                            <Td><Maintainer text={pkgInfo!.maintainer} /></Td>
                        </Tr>
                        <Tr>
                            <Th>Last Updated</Th>
                            <Td>Today</Td>
                        </Tr>
                        <Tr>
                            <Th>Votes</Th>
                            <Td>{Math.floor(Math.random() * 1200) + 30}</Td>
                        </Tr>
                        <Tr>
                            <Th>Popularity</Th>
                            <Td>{Math.floor(Math.random() * 1000) / 10}%</Td>
                        </Tr>
                        <Tr>
                            <Th>Dependencies</Th>
                            <Td>{allDependencies.length || 'None'} {allDependencies.length > 0 ? <Link pl='2' color='pink.400' as={Rlink} to={`/packages/${name}/dependencies`}>View</Link> : ''}</Td>
                        </Tr>
                        <Tr>
                            <Th>Required By</Th>
                            <Td>{pkgInfo?.requiredBy.length || 'None'} {pkgInfo?.requiredBy?.length || 0 > 0 ? <Link color='pink.400' as={Rlink} to={`/packages/${name}/required-by`}>(see all)</Link> : ''}</Td>
                        </Tr>
                        <Tr>
                            <Th>Pacscript</Th>
                            <Td><Link color='pink.400' isExternal href={`https://github.com/pacstall/pacstall-programs/blob/master/packages/${name}/${name}.pacscript`}>View</Link></Td>
                        </Tr>
                    </Tbody>
                </Table>

                <Box mt='10'>
                    <Heading size='lg'>How to Install</Heading>

                    <ResponsiveStack justify='space-between'>
                        <Text fontWeight='semibold' m='3'>Step 1:  Install Pacstall</Text>
                        <Box>
                            <OneLineCodeSnippet size="sm">sudo bash -c "$(wget -q https://git.io/JsADh -O -)"</OneLineCodeSnippet>
                        </Box>
                    </ResponsiveStack>

                    <ResponsiveStack justify='space-between'>
                        <Text fontWeight='semibold' m='3'>Step 2a: Enable Browser Integration</Text>
                        <Box>
                            <OneLineCodeSnippet size="sm">sudo pacstall enable browser-integration</OneLineCodeSnippet>
                        </Box>
                    </ResponsiveStack>

                    <ResponsiveStack justify='space-between'>
                        <Text fontWeight='semibold' m='3'>Step 3:  Click on Install Now</Text>
                        <InstallNowButton disabled={isMobile} />
                    </ResponsiveStack>

                    <ResponsiveStack justify='space-between'>
                        <Text fontWeight='semibold' m='3'>Step 2b:  Alternatively, you use the Terminal</Text>
                        <Box>
                            <SmartCodeSnippetInstall size="sm" name={name} />
                        </Box>
                    </ResponsiveStack>
                </Box>

                <Box mt='10'>
                    <Heading size='lg'>Latest Comments</Heading>
                </Box>
            </Container>
        ) || (
                <Box pt='10' textAlign='center'>
                    <Spinner size='lg' />
                </Box>
            )}
    </>
    )
}

export default PackageDetails