import { Fade } from '@chakra-ui/react'
import { FC, ReactNode, useMemo } from 'react'
import browser from '../../util/browser'

const Empty: FC<any> = ({ children }) => <>{children}</>

const PageAnimation: FC<{ children: ReactNode }> = ({ children }) => {
    const Fader: typeof Fade = useMemo(
        () => (browser.isFirefox ? (Empty as any) : Fade),
        [],
    )

    return (
        <Fader
            in
            transition={{
                enter: { type: 'spring', damping: 5, stiffness: 50 },
            }}
        >
            {children}
        </Fader>
    )
}
export default PageAnimation
