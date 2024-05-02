import { EmailIcon } from '@chakra-ui/icons'
import { Icon, Link, Tooltip } from '@chakra-ui/react'
import { FC, Fragment } from 'react'
import { useTranslation } from 'react-i18next'

const PackageDetailsMaintainer: FC<{ text: string[] }> = ({
    text: maintainers,
}) => {
    const { t } = useTranslation()

    if (
        !maintainers?.length ||
        maintainers[0] === '-' ||
        maintainers[0].toLowerCase() === 'orphan' ||
        maintainers[0].toLowerCase() === 'orphaned'
    ) {
        return <>{t('packageDetails.orphaned')}</>
    }

    if (!['<', '>', '@'].every(symbol => maintainers.join().includes(symbol))) {
        return <>{maintainers.join(', ')}</>
    }

    const shortenName = (name: string, splitBy: string): string =>
        name
            .split(splitBy)
            .reduce(
                (acc, part) =>
                    (acc + part).length > 14 ? acc : acc + splitBy + part,
                '',
            )

    const maintainerInfos = maintainers.map(maintainer => {
        let name = maintainer.split('<')[0].trim()
        if (name.length > 15) {
            if (name.includes(' ')) {
                name = shortenName(name, ' ')
            } else if (name.includes('-')) {
                name = shortenName(name, '-')
            } else {
                name = name.substring(0, 12) + '..'
            }
        }

        const fullEmail = maintainer.split('<')[1]?.split('>')[0]?.trim()
        return {
            name,
            fullEmail,
            shortEmail:
                fullEmail?.split('@')[0]?.length > 15
                    ? fullEmail.split('@')[0]?.substring(0, 13) +
                      '[..]@' +
                      fullEmail.split('@')[1]
                    : fullEmail,
        }
    })

    return (
        <>
            {maintainerInfos.map((maintainer, idx) => (
                <Fragment key={idx}>
                    <span>{maintainer.name} </span>
                    {maintainer.fullEmail && (
                        <Tooltip
                            openDelay={500}
                            label={`Contact ${name} via email`}
                        >
                            <Link
                                color='pink.400'
                                href={'mailto: ' + maintainer.fullEmail}
                            >
                                {maintainer.shortEmail}{' '}
                                <Icon size='md' mx='2px' as={EmailIcon} />{' '}
                            </Link>
                        </Tooltip>
                    )}
                    <br />
                </Fragment>
            ))}
        </>
    )
}

export default PackageDetailsMaintainer
