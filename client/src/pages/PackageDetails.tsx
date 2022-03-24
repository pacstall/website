import { FC, useEffect, useMemo } from "react";
import { useParams, Navigate } from "react-router-dom";
import PackageDetailsPage from "../components/package-details/PackageDetailsPage";
import usePackageInfo from "../hooks/usePackageInfo";
import useDeviceType from "../hooks/useDeviceType";
import { Box, Spinner, useDisclosure } from "@chakra-ui/react";
import { useQueryParam } from "use-query-params";

type OpenPopup = 'required' | 'dependencies' | null

const PackageDetails: FC = () => {
    const { isMobile } = useDeviceType()
    const { name } = useParams() as { name: string }
    const { data, error, loading } = usePackageInfo(name)
    const requiredByModal = useDisclosure()
    const dependenciesModal = useDisclosure()
    const [openPopup, setOpenPopup] = useQueryParam<OpenPopup>('popup', {
        decode(value) {
            const val = value || ''
            return ['required', 'dependencies'].includes(val.toString()) ? val as OpenPopup : null
        },
        encode(value) {
            return value
        }
    })

    if (error) {
        return <Navigate to="/not-found" />
    }

    useEffect(() => {
        if (openPopup === 'dependencies') {
            dependenciesModal.onOpen()
        } else if (openPopup === 'required') {
            requiredByModal.onOpen()
        }
    }, [])

    useEffect(() => {
        const openRequired = requiredByModal.onOpen
        requiredByModal.onOpen = () => {
            setOpenPopup('required')
            openRequired()
        }

        const closeRequired = requiredByModal.onClose
        requiredByModal.onClose = () => {
            setOpenPopup(null)
            closeRequired()
        }

        const openDeps = dependenciesModal.onOpen
        dependenciesModal.onOpen = () => {
            setOpenPopup('required')
            openDeps()
        }

        const closeDeps = dependenciesModal.onClose
        dependenciesModal.onClose = () => {
            setOpenPopup(null)
            closeDeps()
        }
    }, [requiredByModal, dependenciesModal])

    const allDependencies = useMemo(() => [
        ...(data?.buildDependencies || []),
        ...(data?.runtimeDependencies || []),
        ...(data?.optionalDependencies || []),
        ...(data?.pacstallDependencies || [])
    ], [data]);

    return (
        !loading && !error && !!data && (
            <PackageDetailsPage
                allDependencies={allDependencies}
                isMobile={isMobile}
                data={data}
                requiredByModal={requiredByModal}
                dependenciesModal={dependenciesModal}
            />
        ) || (
            <Box pt='10' textAlign='center'>
                <Spinner size='lg' />
            </Box>
        )
    )

}

export default PackageDetails