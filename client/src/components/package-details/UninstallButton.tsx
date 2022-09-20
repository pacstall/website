import { Button, useColorModeValue } from '@chakra-ui/react'
import { FC } from 'react'

const UninstallButton: FC<{ disabled?: boolean }> = ({ disabled }) => (
    <Button
        color={useColorModeValue('black', 'white')}
        bg={useColorModeValue('red.500', 'red.800')}
        _hover={{ bg: useColorModeValue('red.400', 'red.700') }}
        _active={{ bg: useColorModeValue('red.400', 'red.700') }}
        disabled={disabled}
        size='lg'
    >
        Remove
    </Button>
)

export default UninstallButton
