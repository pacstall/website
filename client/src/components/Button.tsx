import { Button, ButtonProps, useColorModeValue } from "@chakra-ui/react";
import { FC } from "react";

export const PrimaryButton: FC<ButtonProps> = (props) => {

    return <Button
        bg={useColorModeValue('brand.200', 'brand.500')}
        _hover={{ background: useColorModeValue('brand.300', 'brand.400') }}
        {...props}
    >
        {props.children}
    </Button>
}