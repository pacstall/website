import React, { useEffect, useRef, useState } from 'react'
import * as AsciinemaPlayerLibrary from 'asciinema-player'

type AsciinemaPlayerProps = {
    src: string
    // START asciinemaOptions
    cols?: string
    rows?: string
    autoPlay?: boolean
    preload?: boolean
    loop?: boolean | number
    startAt?: number | string
    speed?: number
    idleTimeLimit?: number
    theme?: string
    poster?: string
    fit?: string
    fontSize?: string
    // END asciinemaOptions
}

function AsciinemaFrame({ src, ...asciinemaOptions }: AsciinemaPlayerProps) {
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
