import { Box, Button, Flex, HStack, Input, Select, useColorModeValue } from "@chakra-ui/react"
import { FC, useRef, MutableRefObject, KeyboardEventHandler } from "react"

const Search: FC<{ isLoading: boolean, placeholder: string, onSearch: (filter: string, filterBy: string) => any }> = ({ placeholder, onSearch, isLoading }) => {
    const inputRef = useRef<HTMLInputElement>() as MutableRefObject<HTMLInputElement>
    const selectRef = useRef<HTMLSelectElement>() as MutableRefObject<HTMLSelectElement>

    const onPressEnterKey: KeyboardEventHandler = e => {
        if (e.key === 'Enter') {
            onSearch(inputRef.current.value, selectRef.current.value)
        }
    }

    return (
        <HStack justify='space-between'>
            <HStack>
                <Select
                    maxW='150px'
                    defaultValue="name"
                    ref={selectRef}
                    className="focus:outline-none appearance-none h-full rounded-l rounded-r border sm:rounded-r-none sm:border-r-0 border-r border-b block appearance-none w-full bg-white border-gray-400 text-gray-700 py-2 px-4 pr-8 leading-tight focus:outline-none focus:border-l focus:border-r focus:bg-white focus:border-gray-500">
                    <option value="name">Package</option>
                    <option value="maintainer">Maintainer</option>
                </Select>

                <Input minW={{ base: 'xs', xs: '0' }} onKeyUp={onPressEnterKey} placeholder={placeholder} ref={inputRef} />
            </HStack>

            <Button
                isLoading={isLoading}
                loadingText="Searching"
                onClick={() => onSearch(inputRef.current.value, selectRef.current.value)}
                bgColor={useColorModeValue('brand.200', 'brand.500')}
                _hover={{ bg: useColorModeValue('brand.300', 'brand.400') }}
                px='10'
            >
                Search
            </Button>
        </HStack>
    )
}

export default Search