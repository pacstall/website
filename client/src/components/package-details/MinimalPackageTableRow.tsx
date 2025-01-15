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

const MinimalPackageTableRow: FC<{
    pkg: string
    description: string
    external: boolean
}> = ({ pkg, description, external }) => {
    const { t } = useTranslation()

    return (
        <Tr key={pkg}>
            <Td py='8px'>
                <Tooltip
                    openDelay={500}
                    label={
                        description ||
                        t('packageDetails.dependenciesModal.noDescription')
                    }
                >
                    <Text fontSize='md' fontWeight='500'>
                        {external ? (
                            <>{pkg}</>
                        ) : (
                            <Link
                                as={Rlink}
                                target='_blank'
                                color={useColorModeValue(
                                    'pink.600',
                                    'pink.400',
                                )}
                                to={`/packages/${pkg}`}
                            >
                                {pkg}
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
