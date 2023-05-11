import { Button, HStack, LinkBox, useColorModeValue } from '@chakra-ui/react'
import { FC } from 'react'
import { useTranslation } from 'react-i18next';
import { Link as Rlink } from 'react-router-dom'

const PageLink: FC<{ page: number; active?: boolean; disabled?: boolean }> = ({
    page,
    active,
    disabled,
}) => (
    <LinkBox
        as={Rlink}
        to={
            disabled
                ? '#'
                : location.search.replace(/page=[0-9]*/, `page=${page}`)
        }
    >
        <Button
            bg={
                active
                    ? useColorModeValue('gray.400', 'gray.500')
                    : useColorModeValue('gray.200', 'gray.700')
            }
            disabled={disabled}
        >
            {page || '0'}
        </Button>
    </LinkBox>
)

const PageSequentialLink: FC<{
    nextPage: number | string
    text: string
    disabled?: boolean
}> = ({ nextPage, text, disabled }) => (
    <LinkBox
        as={Rlink}
        to={
            disabled
                ? '#'
                : location.search.replace(/page=[0-9]*/, `page=${nextPage}`)
        }
    >
        <Button disabled={disabled}>{text}</Button>
    </LinkBox>
)

const Dots: FC = () => (
    <span
        style={{
            transform: 'translateY(-10%)',
            fontWeight: '800',
            fontSize: '1.5em',
        }}
    >
        ...
    </span>
)

const Pagination: FC<{ last: number; current: number }> = ({
    last,
    current,
}) => {
    const { t } = useTranslation()
    return (
        <HStack>
            <PageSequentialLink
                text={t('packageSearch.pagination.previous')}
                nextPage={current - 1}
                disabled={current === 0}
            />

            {current <= 1 ? (
                <>
                    {new Array(current + 1).fill(null).map((_, i) => (
                        <PageLink key={i} page={i} active={i === current} />
                    ))}
                </>
            ) : (
                <>
                    <PageLink page={0} />
                    <Dots />
                    <PageLink page={current - 1} />
                    <PageLink page={current} active />
                </>
            )}

            {current + 1 < last && <PageLink page={current + 1} />}
            {current + 2 < last ? (
                <>
                    <Dots />
                    <PageLink page={last} />
                </>
            ) : (
                current < last && <PageLink page={last} />
            )}

            <PageSequentialLink
                text={t('packageSearch.pagination.next')}
                nextPage={current + 1}
                disabled={current === last}
            />
        </HStack>
    )
}

export default Pagination
