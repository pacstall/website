import { Box, Heading } from "@chakra-ui/react";
import { FC } from "react";
import Navigation from "../components/Navigation";

const NotFound: FC = () => (
    <>
        <Navigation />
        <Box textAlign='center' mt='20vh'>
            <Heading>404 Page Not Found</Heading>
        </Box>
    </>
)

export default NotFound