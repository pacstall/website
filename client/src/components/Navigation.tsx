import {
    Box,
    Flex,
    Text,
    IconButton,
    Button,
    Stack,
    Collapse,
    Icon,
    Link,
    Popover,
    PopoverTrigger,
    PopoverContent,
    useColorModeValue,
    useBreakpointValue,
    useDisclosure,
    useColorMode,
    Select,
} from '@chakra-ui/react'
import {
    HamburgerIcon,
    CloseIcon,
    ChevronDownIcon,
    ChevronRightIcon,
    MoonIcon,
    SunIcon,
    ExternalLinkIcon,
} from '@chakra-ui/icons'
import { Link as RLink, useNavigate } from 'react-router-dom'
import { PrimaryButton } from './Button'
import { useTranslation } from 'react-i18next'
import { useState } from 'react'
import { localeEntries } from '../locale/locale'

export function Navigation() {
    const { isOpen, onToggle } = useDisclosure()
    const { colorMode, toggleColorMode } = useColorMode()
    const navigate = useNavigate()
    const { t, i18n } = useTranslation()
    const [language, setLanguage] = useState(i18n.language)

    const handleLanguageChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const newLang = e.target.value
        i18n.changeLanguage(newLang)
        setLanguage(newLang)
    }

    return (
        <Box>
            <Flex
                bg={useColorModeValue('white', 'gray.800')}
                color={useColorModeValue('gray.600', 'white')}
                minH={'60px'}
                py={{ base: 2 }}
                px={{ base: 4 }}
                borderBottom={1}
                borderStyle={'solid'}
                borderColor={useColorModeValue('gray.200', 'gray.900')}
                align={'center'}
            >
                <Flex
                    flex={{ base: 1, lg: 'auto' }}
                    ml={{ base: -2 }}
                    display={{ base: 'flex', lg: 'none' }}
                >
                    <IconButton
                        onClick={onToggle}
                        icon={
                            isOpen ? (
                                <CloseIcon w={3} h={3} />
                            ) : (
                                <HamburgerIcon w={5} h={5} />
                            )
                        }
                        variant={'ghost'}
                        aria-label={'Toggle Navigation'}
                    />
                </Flex>
                <Flex
                    flex={{ base: 1 }}
                    justify={{ base: 'center', lg: 'start' }}
                >
                    <Text
                        textAlign='left'
                        cursor='pointer'
                        position='relative'
                        bottom='2px'
                        fontSize='xl'
                        display={useBreakpointValue({
                            base: 'none',
                            lg: 'inherit',
                        })}
                        onClick={() => navigate('/')}
                        color={useColorModeValue('brand.800', 'white')}
                    >
                        {t('navbar.title')}
                    </Text>
                    <Text
                        textAlign='left'
                        cursor='pointer'
                        position='absolute'
                        left='55px'
                        top='15px'
                        fontSize='xl'
                        display={useBreakpointValue({
                            base: 'inherit',
                            lg: 'none',
                        })}
                        onClick={() => navigate('/')}
                        color={useColorModeValue('brand.800', 'white')}
                    >
                        {t('navbar.title')}
                    </Text>

                    <Flex display={{ base: 'none', lg: 'flex' }} ml={10}>
                        <DesktopNav />
                    </Flex>
                </Flex>

                <Link
                    as={RLink}
                    to={'/privacy'}
                    px='7'
                    fontSize={'md'}
                    fontWeight={500}
                    color={useColorModeValue('brand.800', 'white')}
                    display={useBreakpointValue({
                        base: 'none',
                        sm: 'initial',
                    })}
                    pb='1'
                    _hover={{
                        textDecoration: 'none',
                        color: 'brand.400',
                    }}
                >
                    {t('navbar.privacy.title')}
                </Link>

                <Link
                    href='https://github.com/pacstall/pacstall#installing'
                    target='_blank'
                    mr='7'
                >
                    <PrimaryButton px='10'>{t('navbar.install')}</PrimaryButton>
                </Link>

                <Box mr='5' textAlign='center'>
                    <Select
                        variant='filled'
                        value={language}
                        onChange={handleLanguageChange}
                    >
                        {localeEntries.map(([key, value]) => (
                            <option key={key} value={key}>
                                {value}
                            </option>
                        ))}
                    </Select>
                </Box>

                <Button onClick={toggleColorMode}>
                    {colorMode === 'light' ? <MoonIcon /> : <SunIcon />}
                </Button>
            </Flex>

            <Collapse in={isOpen} animateOpacity>
                <MobileNav />
            </Collapse>
        </Box>
    )
}

const DesktopNav = () => {
    const linkColor = useColorModeValue('gray.600', 'gray.200')
    const linkHoverColor = useColorModeValue('teal.400', 'teal.400')
    const popoverContentBgColor = useColorModeValue('white', 'gray.800')
    const { t } = useTranslation()

    return (
        <Stack direction={'row'} spacing={4}>
            {NAV_ITEMS.filter(it => !it.smOnly).map(navItem => (
                <Box key={navItem.label}>
                    <Popover trigger={'hover'} placement={'bottom-start'}>
                        <PopoverTrigger>
                            <Link
                                p={2}
                                to={navItem.href ?? '#'}
                                href={navItem.href ?? '#'}
                                fontSize={'md'}
                                fontWeight={500}
                                color={linkColor}
                                as={
                                    navItem.href?.startsWith('/') ? RLink : Link
                                }
                                isExternal={
                                    !!navItem.href?.startsWith('https://')
                                }
                                _hover={{
                                    textDecoration: 'none',
                                    color: linkHoverColor,
                                }}
                            >
                                {t(navItem.label)}{' '}
                                {navItem.children && (
                                    <Icon
                                        position='relative'
                                        bottom='2px'
                                        color={'brand.400'}
                                        w={7}
                                        h={7}
                                        as={ChevronDownIcon}
                                    />
                                )}
                            </Link>
                        </PopoverTrigger>

                        {navItem.children && (
                            <PopoverContent
                                border={0}
                                boxShadow={'xl'}
                                bg={popoverContentBgColor}
                                p={4}
                                rounded={'xl'}
                                minW={'sm'}
                            >
                                <Stack>
                                    {navItem.children.map(child => (
                                        <DesktopSubNav
                                            key={child.label}
                                            {...child}
                                        />
                                    ))}
                                </Stack>
                            </PopoverContent>
                        )}
                    </Popover>
                </Box>
            ))}
        </Stack>
    )
}

const DesktopSubNav = ({ label, href, subLabel }: NavItem) => {
    const { t } = useTranslation()
    return (
        <Link
            href={href}
            role={'group'}
            display={'block'}
            p={2}
            rounded={'md'}
            _hover={{ bg: useColorModeValue('brand.50', 'gray.900') }}
        >
            <Stack direction={'row'} align={'center'}>
                <Box>
                    <Text
                        transition={'all .3s ease'}
                        _groupHover={{ color: 'brand.400' }}
                        fontWeight={500}
                    >
                        {t(label)}
                    </Text>
                    <Text fontSize={'sm'}>{t(subLabel)}</Text>
                </Box>
                <Flex
                    transition={'all .3s ease'}
                    transform={'translateX(-10px)'}
                    opacity={0}
                    _groupHover={{
                        opacity: '100%',
                        transform: 'translateX(0)',
                    }}
                    justify={'flex-end'}
                    align={'center'}
                    flex={1}
                >
                    {href?.startsWith('https://') ? (
                        <Icon
                            color={'brand.400'}
                            w={5}
                            h={5}
                            as={ExternalLinkIcon}
                        />
                    ) : (
                        <Icon
                            color={'brand.400'}
                            w={5}
                            h={5}
                            as={ChevronRightIcon}
                        />
                    )}
                </Flex>
            </Stack>
        </Link>
    )
}

const MobileNav = () => {
    return (
        <Stack
            bg={useColorModeValue('white', 'gray.800')}
            p={4}
            display={{ lg: 'none' }}
        >
            {NAV_ITEMS.map(navItem => (
                <MobileNavItem key={navItem.label} {...navItem} />
            ))}
        </Stack>
    )
}

const MobileNavItem = ({ label, children, href }: NavItem) => {
    const { isOpen, onToggle } = useDisclosure()

    return (
        <Stack spacing={4} onClick={children && onToggle}>
            <Flex
                py={2}
                as={Link}
                href={href ?? '#'}
                justify={'space-between'}
                align={'center'}
                _hover={{
                    textDecoration: 'none',
                }}
            >
                <Text
                    fontWeight={600}
                    color={useColorModeValue('gray.600', 'gray.200')}
                >
                    {label}
                </Text>
                {children && (
                    <Icon
                        as={ChevronDownIcon}
                        transition={'all .25s ease-in-out'}
                        transform={isOpen ? 'rotate(180deg)' : ''}
                        w={6}
                        h={6}
                    />
                )}
            </Flex>

            <Collapse
                in={isOpen}
                animateOpacity
                style={{ marginTop: '0!important' }}
            >
                <Stack
                    mt={2}
                    pl={4}
                    borderLeft={1}
                    borderStyle={'solid'}
                    borderColor={useColorModeValue('gray.200', 'gray.700')}
                    align={'start'}
                >
                    {children &&
                        children.map(child => (
                            <Link key={child.label} py={2} href={child.href}>
                                {child.label}
                            </Link>
                        ))}
                </Stack>
            </Collapse>
        </Stack>
    )
}

interface NavItem {
    label: string
    subLabel?: string
    children?: Array<NavItem>
    href?: string
    smOnly?: boolean
}

const NAV_ITEMS: Array<NavItem> = [
    {
        label: 'navbar.contribute.title',
        children: [
            {
                label: 'navbar.contribute.workOnFeatures',
                href: 'https://github.com/pacstall/pacstall/wiki/How-to-contribute',
            },
            {
                label: 'navbar.contribute.becomeAMaintainer',
                href: 'https://github.com/pacstall/pacstall/wiki/Pacscript-101',
            },
            {
                label: 'navbar.contribute.helpTranslate',
                href: '#',
            },
        ],
    },
    {
        label: 'navbar.social.title',
        children: [
            {
                label: 'navbar.social.discord',
                href: 'https://discord.com/invite/sWB6YtKyvW',
            },
            {
                label: 'navbar.social.matrix',
                href: 'https://matrix.to/#/#pacstall:matrix.org',
            },
            {
                label: 'navbar.social.reddit',
                href: 'https://www.reddit.com/r/pacstall',
            },
            {
                label: 'navbar.social.mastodon',
                href: 'https://social.linux.pizza/@pacstall',
            },
        ],
    },
    {
        label: 'navbar.browse.title',
        href: '/packages',
    },
    {
        label: 'navbar.privacy.title',
        href: '/privacy',
        smOnly: true,
    },
]

export default Navigation
