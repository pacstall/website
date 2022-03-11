import { Box, Heading } from "@chakra-ui/react";
import { FC } from "react";
import { Helmet } from "react-helmet";
import Navigation from "../components/Navigation";

const NotFound: FC = () => (
    <>
        <Helmet>
            <title>Not Found - Pacstall</title>
        </Helmet>
        <Navigation />
        <Box textAlign='center' mt='20vh'>
            <Heading>404 Page Not Found</Heading>
        </Box>
    </>
)

export default NotFound