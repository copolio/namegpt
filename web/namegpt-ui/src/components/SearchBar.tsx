export default function SearchBar(props: {
    className: string,
    placeholder?: string
}) {
    return (
        <div className={props.className}>
            <div className="relative">
                <input
                    className="w-full h-12 px-4 pr-20 text-gray-700 bg-white-200 border border-gray-300 rounded-full focus:outline-none focus:ring-2 focus:ring-blue-500 resize-y resize-x-none"
                    type="text" placeholder={props.placeholder ?? "Search"}/>
                <button
                    className="absolute top-2 right-10 flex items-center justify-center w-8 h-8 text-gray-500 hover:text-gray-700 focus:outline-none">
                    <svg xmlns="http://www.w3.org/2000/svg" className="w-4 h-4" viewBox="0 0 24 24" fill="none"
                         stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <path d="M18 6L6 18M6 6l12 12"/>
                    </svg>
                </button>
                <button
                    className="absolute top-0 right-0 flex items-center justify-center w-12 h-12 text-gray-500 hover:text-gray-700 focus:outline-none">
                    <svg xmlns="http://www.w3.org/2000/svg" className="w-6 h-6" viewBox="0 0 24 24" fill="none"
                         stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                        <circle cx="11" cy="11" r="8"/>
                        <path d="M21 21l-4.35-4.35"/>
                    </svg>
                </button>
            </div>
        </div>
    );
}