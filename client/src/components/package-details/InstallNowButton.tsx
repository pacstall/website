import { Button, useColorModeValue } from "@chakra-ui/react"
import { FC } from "react"

const InstallNowButton: FC<{ disabled?: boolean }> = ({ disabled }) => (
    <Button
        color={useColorModeValue('black', 'white')}
        bg={useColorModeValue('brand.200', 'brand.500')}
        _hover={{ bg: useColorModeValue('brand.200', 'brand.300') }}
        _active={{ bg: useColorModeValue('brand.200', 'brand.300') }}
        disabled={disabled}
        size='lg'>
        Install Now
    </Button>
)

export default InstallNowButton