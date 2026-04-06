import { createFileRoute } from '@tanstack/react-router'
import AudioStreamPlayer from "@/components/AudioStreamPlayer.tsx";

export const Route = createFileRoute('/streams/$stream_id')({
    component: RouteComponent,
})

function RouteComponent() {
    const { stream_id } = Route.useParams()

    return (
        <div className="p-2 border rounded-lg bg-neutral-50">
            <AudioStreamPlayer stream_id={stream_id} />
        </div>
    )
}
