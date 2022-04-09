import { Fade } from '@chakra-ui/react'
import { FC } from 'react'

const PageAnimation: FC = ({ children }) => (
    <Fade
        in
        transition={{ enter: { type: 'spring', damping: 5, stiffness: 50 } }}
    >
        {children}
    </Fade>
)
export default PageAnimation
