import { EmailIcon } from "@chakra-ui/icons"
import { Icon, Link } from "@chakra-ui/react"
import { FC } from "react"

const PackageDetailsMaintainer: FC<{ text: string }> = ({ text }) => {
    if (!text || text === '-' || text.toLowerCase() === 'orphan' || text.toLowerCase() === 'orphaned') {
        return <>Orphaned</>
    }

    if (!['<', '>', '@'].every(symbol => text.includes(symbol))) {
        return <>{text}</>
    }

    const shortenName = (name: string, splitBy: string): string =>
        name.split(splitBy).reduce((acc, part) => (acc + part).length > 14 ? acc : acc + splitBy + part, '')

    let name = text.split('<')[0].trim()
    if (name.length > 15) {
        if (name.includes(' ')) {
            name = shortenName(name, ' ')
        } else if (name.includes('-')) {
            name = shortenName(name, '-')
        } else {
            name = name.substring(0, 12) + '..'
        }
    }

    const fullEmail = text.split('<')[1].split('>')[0].trim()
    const shortEmail = fullEmail.split('@')[0].length > 15 ? fullEmail.split('@')[0].substring(0, 13) + '[..]@' + fullEmail.split('@')[1] : fullEmail

    return (
        <>
            <span>{name}, </span>
            <Link title={`Contact ${name} via email`} color='pink.400' href={"mailto: " + fullEmail}>
                {shortEmail} <Icon size='md' mx='2px' as={EmailIcon} />
            </Link>
        </>
    )
}

export default PackageDetailsMaintainer