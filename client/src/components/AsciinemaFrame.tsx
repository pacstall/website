import { FC } from 'react'

const AsciinemaFrame: FC<{
    id: string
    autoplay?: boolean
    loop?: boolean
}> = ({ id, autoplay, loop }) => (
    <iframe
        src={`https://asciinema.org/a/${id}/iframe?theme=monokai&autoplay=${
            autoplay ? 1 : 0
        }&loop=${loop ? 1 : 0}&speed=0.75`}
        id={`asciicast-iframe-${id}`}
        name={`asciicast-iframe-${id}`}
        scrolling='no'
        data-allowfullscreen='true'
        style={{
            overflow: 'hidden',
            margin: '0px',
            border: '0px',
            display: 'inline-block',
            width: '100%',
            float: 'none',
            visibility: 'visible',
            aspectRatio: '16 / 9',
        }}
    ></iframe>
)

export default AsciinemaFrame
