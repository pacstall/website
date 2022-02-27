import { FC } from "react";
import Navigation from "../components/Navigation";
import ShowcaseStyles from '../../public/styles/showcase.module.css'
import { Container, Heading } from "@chakra-ui/react";

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
            height: '649px'
        }}>

    </iframe>
)

const Showcase: FC = () => (
    <>
        <Navigation />
        <Container maxW='900px' mt='10'>
            <div className={ShowcaseStyles.scriptContainer}>
                <Heading mb='3'>Showcase</Heading>
                <AsciinemaFrame id="459473" />

                <Heading mb='3'>Package search</Heading>
                <AsciinemaFrame id="459474" />
            </div>
        </Container>

    </>
)

export default Showcase