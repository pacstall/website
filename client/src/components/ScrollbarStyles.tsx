import { useColorModeValue } from '@chakra-ui/react'
import { FC } from 'react'

const ScrollbarStyles: FC = () => {
    const scrollbarStyle = `
    ::-webkit-scrollbar {
        width: 12px;
    }

    /* Track */
    ::-webkit-scrollbar-track {
        box-shadow: inset 0 0 1px ${useColorModeValue('var(--chakra-colors-gray-200)', 'var(--chakra-colors-gray-700)')}; 
        border-radius: 2px;
    }

    /* Handle */
    ::-webkit-scrollbar-thumb {
        background: ${useColorModeValue('var(--chakra-colors-gray-200)', 'var(--chakra-colors-gray-700)')}; 
        border-radius: 7px;
    }

    /* Handle on hover */
    ::-webkit-scrollbar-thumb:hover {
        background: ${useColorModeValue('var(--chakra-colors-gray-300)', 'var(--chakra-colors-gray-600)')}; 
    }
    `

    return <style dangerouslySetInnerHTML={{ __html: scrollbarStyle }}></style>
}

export default ScrollbarStyles
