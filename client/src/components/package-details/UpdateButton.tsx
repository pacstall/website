import { Button, useColorModeValue } from '@chakra-ui/react'
import { FC } from 'react'

const UpdateButton: FC<{ disabled?: boolean }> = ({ disabled }) => (
    <Button
        color={useColorModeValue('black', 'white')}
        bg={useColorModeValue('blue.500', 'blue.200')}
        _hover={{ bg: useColorModeValue('brand.200', 'brand.300') }}
        _active={{ bg: useColorModeValue('brand.200', 'brand.300') }}
        disabled={disabled}
        size='lg'
    >
        Update
    </Button>
)

export default UpdateButton
