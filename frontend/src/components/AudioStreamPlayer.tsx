import { useState } from "react";
import { Button } from "@/components/ui/button.tsx";
import {
    Combobox,
    ComboboxContent,
    ComboboxEmpty,
    ComboboxTrigger,
    ComboboxValue,
    ComboboxItem,
    ComboboxList, ComboboxInput,
} from "@/components/ui/combobox"
import { Play, Pause, ChevronDown, Volume2, VolumeX } from "lucide-react"
import FlagIcon from "@/components/FlagIcon.tsx"
import {useAudioStream} from "@/hooks/useAudioStream.ts";
import {useLanguages} from "@/hooks/useLanguages.ts";

interface AudioStreamPlayerProps {
    className?: string
    stream_id: string
}

export default function AudioStreamPlayer({
  className,
  stream_id,
}: AudioStreamPlayerProps) {

    const { isPending, isError, data, error } = useLanguages()

    const [language, setLanguage] = useState<string>("en-US")

    const {
        isStarted,
        volumeOn,
        handlePlayPause,
        handleVolume
    } = useAudioStream(language, stream_id)

    const getLanguageName = (code: string) => {
        return data?.find((lang) => lang.code === code)?.name
    }

    if (isPending) return (<div>Loading...</div>)
    if (isError) return (<div>Error: {error.message}</div>)

    return (
        <div className={`flex items-center gap-2 ${className}`}>
            <Combobox items={data} value={getLanguageName(language)} onValueChange={(val) => { if (val) setLanguage(val); }}>
                <ComboboxTrigger render={
                    <Button variant="outline" className="flex-1 min-w-0 md:flex-0 md:min-w-64 h-9.5 justify-between font-normal text-sm cursor-pointer bg-white" disabled={isStarted}>
                        <div className="flex items-center gap-2">
                            <FlagIcon code={language} />
                            <ComboboxValue />
                        </div>
                        <ChevronDown />
                    </Button>
                } />
                <ComboboxContent>
                    <ComboboxInput showTrigger={false} placeholder="Search" />
                    <ComboboxEmpty>No items found.</ComboboxEmpty>
                    <ComboboxList>
                        {(item) => (
                            <ComboboxItem
                                key={item.code}
                                value={item.code}
                                className="cursor-pointer"
                            >
                                <FlagIcon code={item.code} />
                                {item.name}
                            </ComboboxItem>
                        )}
                    </ComboboxList>
                </ComboboxContent>
            </Combobox>

            <Button variant="default" size="icon-xl" className="cursor-pointer" onClick={handlePlayPause}>
                {isStarted ? <Pause /> : <Play />}
            </Button>

            <Button variant="outline" size="icon-xl" className="cursor-pointer" onClick={handleVolume}>
                {volumeOn ? <Volume2 /> : <VolumeX />}
            </Button>
        </div>
    )
}