import { ExternalLinkIcon } from "@chakra-ui/icons";
import { Button, Modal, Text, ModalBody, ModalCloseButton, ModalContent, ModalFooter, ModalHeader, ModalOverlay, useDisclosure, Link, Icon, useColorModeValue } from "@chakra-ui/react";
import { FC, useEffect } from "react";
import { Link as Rlink } from 'react-router-dom'
import useCookie from 'react-use-cookie'

const CookieBanner: FC = () => {
    const [cookie, setCookie] = useCookie('privacy-policy-accepted')
    const { isOpen, onOpen, onClose } = useDisclosure({ isOpen: !cookie })

    const onCookieAccept = () => {
        setCookie('1')
        onClose()
    }

    return (
        <Modal isCentered isOpen={isOpen} onClose={onClose}>
            <ModalOverlay
                bg='blackAlpha.300'
                backdropFilter='blur(10px)'
            />
            <ModalContent>
                <ModalHeader>Cookie notice TL;DR</ModalHeader>
                <ModalBody>
                    <Text>Hi, yes, we use cookies.</Text>
                    <Text mt='3'>We don't like to give you misleading or confusing information. We only use cookies for essential features such as theme settings and authentication.</Text>
                    <Text mt='3'>You can read the full privacy policy <Link as={Rlink} to={'/privacy'} color='pink.400' target='_blank'>here <Icon as={ExternalLinkIcon} mb='1' /></Link>.</Text>
                    <Text mt='3'>By continuing to use this website, you <strong>give your agreement to the privacy policy.</strong></Text>
                </ModalBody>
                <ModalFooter justifyContent='center'>
                    <Button color='white' _hover={{ bg: useColorModeValue('brand.500', 'brand.300') }} onClick={onCookieAccept} bg='brand.400' size='lg'>ok, nice</Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    )
}

export default CookieBanner