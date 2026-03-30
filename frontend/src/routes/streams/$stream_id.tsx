import { createFileRoute } from '@tanstack/react-router'
import AudioStreamPlayer from "@/components/AudioStreamPlayer.tsx";

export const Route = createFileRoute('/streams/$stream_id')({
    component: RouteComponent,
})

function RouteComponent() {
    const { stream_id } = Route.useParams()

    return (
        <>
            <AudioStreamPlayer stream_id={stream_id} />
        </>
    )
}
