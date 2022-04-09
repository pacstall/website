import { Box, Spinner } from '@chakra-ui/react'
import { FC } from 'react'

const ComponentLoader: FC<{
    isLoading: boolean
    hasError?: boolean
    content: () => JSX.Element
}> = ({ isLoading, hasError, content: Content }) => {
    if (isLoading) {
        return (
            <Box pt='10' textAlign='center'>
                <Spinner size='lg' />
            </Box>
        )
    }

    if (hasError === true) {
        return <>Couldn't load this section. Please report this incident.</>
    }

    return <Content />
}

export default ComponentLoader
