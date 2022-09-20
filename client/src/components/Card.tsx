import { FC, ReactNode } from 'react'
import { Stack, Text } from '@chakra-ui/react'

const Card: FC<{ title: string; children: ReactNode }> = ({ title, children }) => {
    return (
        <Stack p='4' boxShadow='lg' m='4' borderRadius='sm'>
            <Stack direction='row' alignItems='center'>
                <Text fontWeight='semibold'>{title}</Text>
            </Stack>

            {children}
        </Stack>
    )
}

export default Card
