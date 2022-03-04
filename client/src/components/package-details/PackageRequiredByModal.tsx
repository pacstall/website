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
} from "@chakra-ui/react";
import { FC } from "react";
import usePackageRequiredBy from "../../hooks/usePackageRequiredBy";
import ComponentLoader from "../ComponentLoader";
import MinimalPackageTable from "./MinimalPackageTable";

const PackageRequiredByModal: FC<{ name: string } & UseDisclosureProps> = ({ name, isOpen, onClose }) => {
    const { data, loading, error } = usePackageRequiredBy(name)

    const ComputedPackageList = () => <MinimalPackageTable packages={data} />

    return (
        <Modal scrollBehavior="inside" isOpen={isOpen!} onClose={onClose!}>
            <ModalOverlay />
            <ModalContent>
                <ModalHeader>Required by</ModalHeader>
                <ModalCloseButton />
                <ModalBody>
                    <ComponentLoader isLoading={loading} hasError={error} content={ComputedPackageList} />
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

export default PackageRequiredByModal