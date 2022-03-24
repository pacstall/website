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
} from "@chakra-ui/react";
import { FC } from "react";
import usePackageDependencies from "../../hooks/usePackageDependencies";
import MinimalPackageTable from "./MinimalPackageTable";

const PackageDependenciesModal: FC<{ name: string } & UseDisclosureProps> = ({ name, isOpen, onClose }) => {
    const { data, loaded } = usePackageDependencies(name)

    return (
        <Modal scrollBehavior="inside" size='xl' isOpen={isOpen!} onClose={onClose!}>
            <ModalOverlay />
            <ModalContent maxH='50vh'>
                <ModalHeader>Dependencies</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    {loaded && data && (
                        <>
                            {data!.runtimeDependencies.length > 0 && (
                                <Box mb='5'>
                                    <Heading mb='1' size="xs">Runtime Dependencies</Heading>
                                    <MinimalPackageTable packages={data!.runtimeDependencies || []} />
                                </Box>
                            )}

                            {data!.buildDependencies.length > 0 && (
                                <Box mb='5'>
                                    <Heading mb='1' size="xs">Build Dependencies</Heading>
                                    <MinimalPackageTable packages={data!.buildDependencies || []} />
                                </Box>
                            )}

                            {data!.optionalDependencies.length > 0 && (
                                <Box mb='5'>
                                    <Heading mb='1' size="xs">Optional Dependencies</Heading>
                                    <MinimalPackageTable packages={data!.optionalDependencies || []} />
                                </Box>
                            )}

                            {data!.pacstallDependencies.length > 0 && (
                                <Box mb='5'>
                                    <Heading mb='1' size="xs">Pacstall Dependencies</Heading>
                                    <MinimalPackageTable packages={data!.pacstallDependencies || []} />
                                </Box>
                            )}
                        </>
                    )}
                </ModalBody>

                <ModalFooter>
                    <Button colorScheme='blue' variant='ghost' mr={3} onClick={onClose}>
                        Close
                    </Button>
                </ModalFooter>
            </ModalContent>
        </Modal>
    )
}

export default PackageDependenciesModal