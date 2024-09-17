import {
    Link,
    Td,
    Text,
    Tooltip,
    Tr,
    useColorModeValue,
} from '@chakra-ui/react'
import { FC } from 'react'
import { useTranslation } from 'react-i18next'
import { Link as Rlink } from 'react-router-dom'

const getDescription = (nameWithDescription: string): string | null => {
    return nameWithDescription?.includes(':')
        ? nameWithDescription.split(':')[1].trim()
        : null
}

const getName = (nameWithDescription: string): string => {
    return nameWithDescription?.includes(':')
        ? nameWithDescription.split(':')[0]
        : nameWithDescription
}

const MinimalPackageTableRow: FC<{ pkg: string; external: boolean }> = ({
    pkg,
    external,
}) => {
    const { t } = useTranslation()

    return (
        <Tr key={pkg}>
            <Td py='8px'>
                <Tooltip
                    openDelay={500}
                    label={
                        getDescription(pkg) ||
                        t('packageDetails.dependenciesModal.noDescription')
                    }
                >
                    <Text fontSize='md' fontWeight='500'>
                        {external ? (
                            <>{getName(pkg)}</>
                        ) : (
                            <Link
                                as={Rlink}
                                target='_blank'
                                color={useColorModeValue(
                                    'pink.600',
                                    'pink.400',
                                )}
                                to={`/packages/${getName(pkg)}`}
                            >
                                {getName(pkg)}
                            </Link>
                        )}
                    </Text>
                </Tooltip>
            </Td>
            <Td py='8px' textAlign='right'>
                <Text fontSize='sm'>
                    {external ? 'APT' : 'Pacstall Repository'}
                </Text>
            </Td>
        </Tr>
    )
}

export default MinimalPackageTableRow
