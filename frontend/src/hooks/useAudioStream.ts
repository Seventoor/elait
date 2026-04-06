import {useEffect, useRef, useState, useCallback} from "react";
import {config} from "@/config.ts";

export function useAudioStream(language: string, stream_id: string) {
    const [audioQueue, setAudioQueue] = useState<ArrayBuffer[]>([])
    const [isPlaying, setIsPlaying] = useState<boolean>(false)
    const [isStarted, setIsStarted] = useState<boolean>(false)
    const [volumeOn, setVolumeOn] = useState<boolean>(true)

    const audioCtxRef = useRef<AudioContext | null>(null)
    const gainNodeRef = useRef<GainNode | null>(null)

    if (!audioCtxRef.current) {
        audioCtxRef.current = new AudioContext()
    }

    if (!gainNodeRef.current) {
        gainNodeRef.current = audioCtxRef.current.createGain()
        gainNodeRef.current.connect(audioCtxRef.current.destination)
    }

    const playNext = useCallback(async () => {
        if (isPlaying || audioQueue.length === 0) return

        setIsPlaying(true)
        const rawData = audioQueue[0]
        setAudioQueue(prev => prev.slice(1))

        try {
            const audioBuffer = await audioCtxRef.current!.decodeAudioData(rawData)
            const source = audioCtxRef.current!.createBufferSource()
            source.buffer = audioBuffer
            source.connect(gainNodeRef.current!)
            gainNodeRef.current!.connect(audioCtxRef.current!.destination)

            source.onended = () => {
                setIsPlaying(false)
            }

            source.start(0)
        } catch (e) {
            console.error("Fehler beim Dekodieren:", e)
            setIsPlaying(false)
            playNext()
        }
    }, [isPlaying, audioQueue])

    const handlePlayPause = () => {
        if (isStarted) {
            setIsStarted(false)
            setAudioQueue([])
            setIsPlaying(false)
        } else {
            setIsStarted(true)
        }
    }

    const handleVolume = () => {
        if (volumeOn) {
            setVolumeOn(false)
            gainNodeRef.current!.gain.value = 0
        } else {
            setVolumeOn(true)
            gainNodeRef.current!.gain.value = 1
        }
    }

    useEffect(() => {
        if (isStarted) {
            const ws = new WebSocket(`${config.wsUrl}/audio-ws?lang=${language}&stream_id=${stream_id}`)
            ws.binaryType = 'arraybuffer'

            ws.onmessage = (event) => {
                setAudioQueue(prev => [...prev, event.data])
            }

            ws.onclose = () => {
                setAudioQueue([])
            }

            return () => ws.close()
        }
    }, [language, stream_id, isStarted])

    useEffect(() => {
        if (isStarted && !isPlaying && audioQueue.length > 0) {
            playNext()
        }
    }, [audioQueue, isPlaying, playNext, isStarted])

    return {
        isStarted,
        isPlaying,
        volumeOn,
        handlePlayPause,
        handleVolume,
    }
}

