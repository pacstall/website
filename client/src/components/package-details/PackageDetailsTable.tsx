import { ExternalLinkIcon } from '@chakra-ui/icons'
import {
    Table,
    Tbody,
    Link,
    Tr,
    Th,
    Td,
    Text,
    Icon,
    UseDisclosureProps,
} from '@chakra-ui/react'
import { FC, ReactNode } from 'react'
import { Link as Rlink } from 'react-router-dom'
import PackageInfo from '../../types/package-info'
import SemanticVersionColor from '../SemanticVersionColor'
import PackageDetailsMaintainer from './PackageDetailsMaintainer'

const Entry: FC<{
    header: string
    disabled?: boolean
    children: ReactNode
}> = ({ header, children, disabled }) => (
    <>
        {!disabled && (
            <Tr>
                <Th>{header}</Th>
                <Td>{children}</Td>
            </Tr>
        )}
    </>
)

const PackageDetailsTable: FC<{
    data: PackageInfo
    dependencyCount: number
    requiredByModal: UseDisclosureProps
    dependenciesModal: UseDisclosureProps
}> = ({ data, dependencyCount, requiredByModal, dependenciesModal }) => {
    return (
        <Table mt='10'>
            <Tbody>
                <Entry header='Name'>{data.name}</Entry>

                <Entry header='Version'>
                    <Text fontWeight='bold'>
                        <SemanticVersionColor
                            git={data.name.endsWith('-git')}
                            version={data.version}
                            status={data.updateStatus}
                        />
                    </Text>
                </Entry>

                <Entry header='Maintainer'>
                    <PackageDetailsMaintainer text={data.maintainer} />
                </Entry>

                <Entry header='Dependencies'>
                    {dependencyCount || 'None'}{' '}
                    {dependencyCount > 0 ? (
                        <Link
                            onClick={dependenciesModal.onOpen}
                            pl='2'
                            color='pink.400'
                            as={Rlink}
                            to={`#`}
                        >
                            View
                        </Link>
                    ) : (
                        ''
                    )}
                </Entry>

                <Entry header='Required By'>
                    {data.requiredBy.length || 'None'}{' '}
                    {data.requiredBy?.length || 0 > 0 ? (
                        <Link
                            onClick={requiredByModal.onOpen}
                            pl='2'
                            color='pink.400'
                            as={Rlink}
                            to={`#`}
                        >
                            View
                        </Link>
                    ) : (
                        ''
                    )}
                </Entry>

                <Entry header='Pacscript'>
                    <Link
                        color='pink.400'
                        isExternal
                        href={`https://github.com/pacstall/pacstall-programs/blob/master/packages/${data.name}/${data.name}.pacscript`}
                    >
                        Open in GitHub{' '}
                        <Icon
                            position='relative'
                            bottom='2px'
                            size='md'
                            ml='1px'
                            as={ExternalLinkIcon}
                        />
                    </Link>
                </Entry>
            </Tbody>
        </Table>
    )
}

export default PackageDetailsTable
