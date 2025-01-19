import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalCloseButton,
    ModalBody,
    ModalFooter,
    Button,
    UseDisclosureProps,
} from '@chakra-ui/react'
import { FC } from 'react'
import usePackageRequiredBy from '../../hooks/usePackageRequiredBy'
import ComponentLoader from '../ComponentLoader'
import MinimalPackageTable from './MinimalPackageTable'
import { useTranslation } from 'react-i18next'

const PackageRequiredByModal: FC<{ name: string } & UseDisclosureProps> = ({
    name,
    isOpen,
    onClose,
}) => {
    const { data, loading, error } = usePackageRequiredBy(name)
    const { t } = useTranslation()

    const ComputedPackageList = () => (
        <MinimalPackageTable packages={data} type='requiredBy' />
    )

    return (
        <Modal scrollBehavior='inside' isOpen={isOpen!} onClose={onClose!}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader>
                    {t('packageDetails.requiredByModal.title')}
                </ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <ComponentLoader
                        isLoading={loading}
                        hasError={error}
                        content={ComputedPackageList}
                    />
                </ModalBody>

                <ModalFooter>
                    <Button
                        colorScheme='blue'
                        variant='ghost'
                        mr={3}
                        onClick={onClose}
                    >
                        {t('packageDetails.requiredByModal.close')}
                    </Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    )
}

export default PackageRequiredByModal
