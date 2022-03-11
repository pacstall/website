import { FC } from "react";
import Navigation from "../components/Navigation";
import { Container, Heading } from "@chakra-ui/react";
import { Helmet } from "react-helmet";

const AsciinemaFrame: FC<{ id: string }> = ({ id }) => (
    <iframe src={`https://asciinema.org/a/${id}/iframe?theme=monokai`}
        id={`asciicast-iframe-${id}`}
        name={`asciicast-iframe-${id}`}
        scrolling="no"
        data-allowfullscreen="true"
        style={{
            overflow: 'hidden',
            margin: '0px',
            border: '0px',
            display: 'inline-block',
            width: '100%',
            float: 'none',
            visibility: 'visible',
            height: 'min(75vw, 650px)'
        }}>

    </iframe>
)

const Showcase: FC = () => (
    <>
        <Helmet>
            <title>Showcase - Pacstall</title>
        </Helmet>
        <Navigation />
        <Container maxW='900px' mt='10'>
            <Heading mb='3'>Showcase</Heading>
            <AsciinemaFrame id="459473" />

            <Heading mb='3'>Package search</Heading>
            <AsciinemaFrame id="459474" />
        </Container>

    </>
)

export default Showcase