import SearchBar from "@/components/SearchBar";

export default function Home() {
    return (
        <>
            <div className="flex items-center justify-center h-full">
                <div className="flex flex-col items-center w-full">
                    <img className="mx-auto mt-28 my-4" src="logo.png" alt="Copolio Logo" style={{width: '250px', height: '250px'}} />
                    <SearchBar className="lg:w-2/5 w-full px-2" placeholder="Type your domain name ideas" />
                </div>
            </div>
        </>
    )
}
