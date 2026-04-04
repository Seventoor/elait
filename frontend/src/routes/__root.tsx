import { createRootRoute, Link, Outlet } from '@tanstack/react-router'

export const Route = createRootRoute({
    component: () => (
        <>
            <nav>
                <Link to="/">Home</Link>
            </nav>
            <main role="main">
                <div className="inset-0 dark:bg-gray-900 overflow-auto flex justify-center z-0 min-h-screen">
                    <div className="w-full max-w-3xl lg:max-w-4xl xl:max-w-4xl space-y-2 px-3 sm:px-4 py-3 sm:py-4">
                        <Outlet />
                    </div>
                </div>
            </main>
        </>
    ),
})