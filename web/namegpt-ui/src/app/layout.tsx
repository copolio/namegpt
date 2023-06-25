import './globals.css'

export const metadata = {
    title: 'NameGPT',
    description: 'Returns domain name recommendations given user input using ChatGPT',
}

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
        <body>
        <div className="flex flex-col min-h-screen">
            <header className="bg-white shadow py-4">
                <div className="max-w-7xl px-4">
                    <nav className="flex items-center justify-between">
                        <a className="text-lg font-bold" href="/">
                            NameGPT
                        </a>
                    </nav>
                </div>
            </header>
            <main className="flex-grow">{children}</main>
            <footer className="py-4">
                <div className="max-w-7xl mx-auto px-4">
                    <p className="text-gray-600 text-center">
                        &copy; {new Date().getFullYear()} Copolio. All rights reserved.
                    </p>
                </div>
            </footer>
        </div>
        </body>
        </html>
    )
}
