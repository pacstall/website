import {
    Box,
    Center,
    Stack,
    Text,
    Grid,
    GridItem,
    chakra,
    useBreakpointValue,
    Icon,
    useColorModeValue,
} from '@chakra-ui/react'
import { FC } from 'react'
import { VscPackage } from 'react-icons/vsc'
import { VscTerminalDebian } from 'react-icons/vsc'
import { VscSymbolNamespace } from 'react-icons/vsc'
import { VscFileBinary } from 'react-icons/vsc'
import { useNavigate } from 'react-router-dom'
import PackageInfo from '../../types/package-info'
import toTitleCase from '../../util/title-case'
import { CheckCircleIcon } from '@chakra-ui/icons'
import getPackageKind, { PackageKind } from '../../util/package-kind'

const packageKindIcon: Record<PackageKind, any> = {
    [PackageKind.App]: VscPackage,
    [PackageKind.DebFile]: VscTerminalDebian,
    [PackageKind.Source]: VscSymbolNamespace,
    [PackageKind.Binary]: VscFileBinary,
}

const PackageGridItem: FC<{ pkg: PackageInfo }> = ({ pkg }) => {
    const navigate = useNavigate()

    const onClick = () => {
        navigate('/packages/' + pkg.name)
    }

    const pkgKind = getPackageKind(pkg.name)
    const icon = packageKindIcon[pkgKind]
    const title = toTitleCase(pkg);

    return (
        <GridItem
            alignSelf='stretch'
            w='100%'
            position='relative'
            onClick={onClick}
            bgColor={useColorModeValue('inherit', 'gray.800')}
            borderWidth='1px'
            borderColor={useColorModeValue('gray.300', 'gray.700')}
            _hover={{
                bgColor: useColorModeValue('inherit','gray.700'),
                borderColor: useColorModeValue('pink.500','gray.700'),
            }}
            style={{
                transition: 'ease 0.2s',
            }}
            borderRadius='10px'
            minW='100%'
            h='115px'
            p='1em'
            overflow='hidden'
        >
            <Stack>
                <chakra.div color='pink.400' fontWeight='semibold' zIndex={1}>
                    {title.length > 30
                        ? title.substring(0, 27) + '...'
                        : title}
                </chakra.div>
                <chakra.div color={useColorModeValue('gray.600', 'gray.400')}>
                    {pkg.description.length > 55
                        ? pkg.description.substring(0, 52) + '...'
                        : pkg.description}
                </chakra.div>
            </Stack>
            {pkg.installed && (
                <CheckCircleIcon
                    position='absolute'
                    right='15px'
                    bottom='13px'
                    fontSize='md'
                    color='brand.400'
                />
            )}
            <Icon
                position='absolute'
                right='15px'
                top='15px'
                color={useColorModeValue('gray.600', 'gray.400')}
                fontSize='xl'
                as={icon}
            />
        </GridItem>
    )
}

const PackageGrid: FC<{ data: PackageInfo[] }> = ({ data }) => (
    <Box mt='10' mx='5'>
        <Grid
            justifyItems='center'
            gridTemplateColumns={useBreakpointValue({
                base: '1fr',
                sm: '1fr 1fr',
                md: '1fr 1fr 1fr',
            })}
            gap='30px'
            pt='1em'
        >
            {data.map(pkg => (
                <PackageGridItem pkg={pkg} />
            ))}
        </Grid>
        {data.length === 0 && (
            <Box mt='5'>
                <Center>
                    <Text>No packages found</Text>
                </Center>
            </Box>
        )}
    </Box>
)

export default PackageGrid
