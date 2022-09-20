import { Center, Heading, Progress } from '@chakra-ui/react'
import { FC, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'

const LoadingOverlay: FC<{ timeout: number; path: string }> = ({
    timeout,
    path,
}) => {
    const navigate = useNavigate()

    useEffect(() => {
        setTimeout(() => navigate(path), timeout)
    }, [])

    return (
        <>
            <Center>
                <Heading>Loading...</Heading>
                <br />
                <Progress colorScheme='pink' isIndeterminate />
            </Center>
        </>
    )
}

export default LoadingOverlay
