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
    Heading,
    Box,
} from '@chakra-ui/react'
import { FC } from 'react'
import usePackageDependencies from '../../hooks/usePackageDependencies'
import MinimalPackageTable from './MinimalPackageTable'
import { useTranslation } from 'react-i18next'

const PackageDependenciesModal: FC<{ name: string } & UseDisclosureProps> = ({
    name,
    isOpen,
    onClose,
}) => {
    const { data, loaded } = usePackageDependencies(name)
    const { t } = useTranslation()

    return (
        <Modal
            scrollBehavior='inside'
            size='xl'
            isOpen={isOpen!}
            onClose={onClose!}
        >
            <ModalOverlay />
            <ModalContent maxH='50vh'>
                <ModalHeader>
                    {t('packageDetails.dependenciesModal.title')}
                </ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    {loaded && data && (
                        <>
                            {data!.runtimeDependencies.length > 0 && (
                                <Box mb='5'>
                                    <Heading mb='1' size='xs'>
                                        {t(
                                            'packageDetails.dependenciesModal.runtimeDependencies',
                                        )}
                                    </Heading>
                                    <MinimalPackageTable
                                        packages={
                                            data!.runtimeDependencies || []
                                        }
                                    />
                                </Box>
                            )}

                            {data!.buildDependencies.length > 0 && (
                                <Box mb='5'>
                                    <Heading mb='1' size='xs'>
                                        {t(
                                            'packageDetails.dependenciesModal.buildDependencies',
                                        )}
                                    </Heading>
                                    <MinimalPackageTable
                                        packages={data!.buildDependencies || []}
                                    />
                                </Box>
                            )}

                            {data!.optionalDependencies.length > 0 && (
                                <Box mb='5'>
                                    <Heading mb='1' size='xs'>
                                        {t(
                                            'packageDetails.dependenciesModal.optionalDependencies',
                                        )}
                                    </Heading>
                                    <MinimalPackageTable
                                        packages={
                                            data!.optionalDependencies || []
                                        }
                                    />
                                </Box>
                            )}

                            {data!.pacstallDependencies.length > 0 && (
                                <Box mb='5'>
                                    <Heading mb='1' size='xs'>
                                        {t(
                                            'packageDetails.dependenciesModal.pacstallDependencies',
                                        )}
                                    </Heading>
                                    <MinimalPackageTable
                                        packages={
                                            data!.pacstallDependencies || []
                                        }
                                    />
                                </Box>
                            )}
                        </>
                    )}
                </ModalBody>

                <ModalFooter>
                    <Button
                        colorScheme='blue'
                        variant='ghost'
                        mr={3}
                        onClick={onClose}
                    >
                        {t('packageDetails.dependenciesModal.close')}
                    </Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    )
}

export default PackageDependenciesModal
