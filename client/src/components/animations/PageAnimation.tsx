import { Fade } from '@chakra-ui/react'
import { FC, ReactNode, useMemo } from 'react'
import { ANIMATIONS_DISABLED } from '../../util/animation'

const Empty: FC<any> = ({ children }) => <>{children}</>

const PageAnimation: FC<{ children: ReactNode }> = ({ children }) => {
    const Fader: typeof Fade = useMemo(
        () => (ANIMATIONS_DISABLED ? (Empty as any) : Fade),
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
