import { ExternalLinkIcon } from '@chakra-ui/icons'
import {
    Button,
    Modal,
    Text,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader,
    ModalOverlay,
    useDisclosure,
    Link,
    Icon,
    useColorModeValue,
} from '@chakra-ui/react'
import { FC, useEffect } from 'react'
import { Trans, useTranslation } from 'react-i18next'
import { Link as Rlink, useLocation } from 'react-router-dom'
import useCookie from 'react-use-cookie'

const CookieBanner: FC = () => {
    const [cookie, setCookie] = useCookie('privacy-policy-accepted')
    const location = useLocation()
    const { isOpen, onOpen, onClose } = useDisclosure({
        isOpen: !cookie && !location.pathname.endsWith('/privacy'),
    })
    const { t } = useTranslation()

    const onCookieAccept = () => {
        setCookie('1')
        onClose()
    }

    useEffect(() => {
        if (!cookie && !isOpen && !location.pathname.endsWith('/privacy')) {
            onOpen()
        }
    }, [location])

    return (
        <Modal isCentered isOpen={isOpen} onClose={onClose}>
            <ModalOverlay bg='blackAlpha.300' backdropFilter='blur(10px)' />
            <ModalContent>
                <ModalHeader>{t('cookieConsent.title')}</ModalHeader>
                <ModalBody>
                    <Text>{t('cookieConsent.paragraphs.0')}</Text>
                    <Text mt='3'>{t('cookieConsent.paragraphs.1')}</Text>
                    <Text mt='3'>
                        <Trans i18nKey='cookieConsent.paragraphs.2'>
                            <Link
                                as={Rlink}
                                to={'/privacy'}
                                color='pink.400'
                                target='_blank'
                            >
                                here <Icon as={ExternalLinkIcon} mb='1' />
                            </Link>
                        </Trans>
                    </Text>
                    <Text mt='3'>
                        <Trans i18nKey='cookieConsent.paragraphs.3' />
                    </Text>
                </ModalBody>
                <ModalFooter justifyContent='center'>
                    <Button
                        color='white'
                        _hover={{
                            bg: useColorModeValue('brand.500', 'brand.300'),
                        }}
                        onClick={onCookieAccept}
                        bg='brand.400'
                        size='lg'
                    >
                        {t('cookieConsent.accept')}
                    </Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    )
}

export default CookieBanner
