import { FC } from "react";
import Navigation from "../components/Navigation";
import * as ShowcaseStyles from '../../public/styles/showcase.module.css'

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
        <hr className="uk-divider-icon"
        />
        <div className={ShowcaseStyles.scriptContainer}>
            <AsciinemaFrame id="459473" />
            <AsciinemaFrame id="459474" />
        </div>

    </>
)

export default Showcase