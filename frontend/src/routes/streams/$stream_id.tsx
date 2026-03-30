import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/streams/$stream_id')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/public/stream"!</div>
}
