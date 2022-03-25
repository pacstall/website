import { useMediaQuery } from "@chakra-ui/react"

const useDeviceType = () => {
    const [isMobile] = useMediaQuery('only screen and (max-device-width: 480px)')

    return {
        isMobile,
        isDesktop: !isMobile
    }
}

export default useDeviceType