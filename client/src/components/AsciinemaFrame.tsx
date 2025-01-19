import React, { useEffect, useRef, useState } from 'react'
import 'asciinema-player/dist/bundle/asciinema-player.css'
import '../../public/styles/asciicast.css'

type AsciinemaPlayerProps = {
    src: string
    // START asciinemaOptions
    cols?: string
    rows?: string
    controls?: boolean | string
    autoPlay?: boolean
    preload?: boolean
    loop?: boolean | number
    startAt?: number | string
    speed?: number
    idleTimeLimit?: number
    theme?: string
    poster?: string
    fit?: string
    terminalfontSize?: string
    terminalFontFamily?: string
    // END asciinemaOptions
}

function AsciinemaFrame({
    src,
    speed = 0.75,
    controls = false,
    ...asciinemaOptions
}: AsciinemaPlayerProps) {
    const ref = useRef<HTMLDivElement>(null)
    const [player, setPlayer] = useState<typeof import('asciinema-player')>()
    useEffect(() => {
        import('asciinema-player').then(p => {
            setPlayer(p)
        })
    }, [])
    useEffect(() => {
        const currentRef = ref.current
        const instance = player?.create(src, currentRef, asciinemaOptions)
        return () => {
            instance?.dispose()
        }
    }, [src, player, asciinemaOptions])

    return <div ref={ref} />
}

export default AsciinemaFrame
