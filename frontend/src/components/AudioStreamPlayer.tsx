import { config } from "@/config"
import {useCallback, useEffect, useRef, useState} from "react";
import { Button } from "@/components/ui/button.tsx";
import {
    Select,
    SelectContent,
    SelectGroup,
    SelectItem,
    SelectLabel,
    SelectTrigger,
    SelectValue,
} from "@/components/ui/select"
import { Play, Pause } from "lucide-react"

interface AudioStreamPlayerProps {
    className?: string
    stream_id: string
}

export default function AudioStreamPlayer({
  className,
  stream_id,
}: AudioStreamPlayerProps) {
    const [language, setLanguage] = useState<string>('ru-RU')
    const [audioQueue, setAudioQueue] = useState<ArrayBuffer[]>([])
    const [isPlaying, setIsPlaying] = useState<boolean>(false)
    const [isStarted, setIsStarted] = useState<boolean>(false)

    // useRef damit audioCtx nicht bei jedem Re-render neu erstellt wird
    const audioCtxRef = useRef<AudioContext | null>(null)
    if (!audioCtxRef.current) {
        audioCtxRef.current = new AudioContext()
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
            source.connect(audioCtxRef.current!.destination)

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

    return (
        <div className={`flex items-center gap-2 ${className}`}>
            <Select onValueChange={setLanguage} disabled={isStarted} >
                <SelectTrigger className="w-full max-w-48">
                    <SelectValue placeholder="Sprache / Language" />
                </SelectTrigger>
                <SelectContent>
                    <SelectGroup>
                        <SelectLabel>Sprachen</SelectLabel>
                        <SelectItem value="en-US">English</SelectItem>
                        <SelectItem value="ru-RU">Русский</SelectItem>
                    </SelectGroup>
                </SelectContent>
            </Select>
            <Button variant="default" size="icon-lg" className="cursor-pointer" onClick={handlePlayPause}>
                {isStarted ? <Pause /> : <Play />}
            </Button>
        </div>
    )
}