"use client"

import SearchBar from "@/components/SearchBar";
import {useState} from "react";
import {
    Configuration,
    GithubComCopolioNamegptPkgDtoGenerateDomainNameResult,
    V1ApiFactory
} from "@/utils/clients/namegpt";
import LoadingSpinner from "@/components/LoadingSpinner";

export default function Home() {
    const v1Api = V1ApiFactory(new Configuration({basePath: "http://localhost:8080"}));
    const [searchResult, setSearchResult] =
        useState<GithubComCopolioNamegptPkgDtoGenerateDomainNameResult[]>([]);
    let [isLoading, setIsLoading] = useState(false);
    const search = (value: string) => {
        setIsLoading(true);
        v1Api.apiV1DomainsSimilarNamesPost({keyword: value})
            .then(res => {
                setSearchResult(res.data)
            })
            .finally(() => setIsLoading(false))
    }
    // @ts-ignore
    return (
        <>
            <div className="flex items-center justify-center h-full">
                <LoadingSpinner isLoading={isLoading} label={"Generating domains...."} />
                <div className="flex flex-col items-center w-full">
                    <img className="mx-auto mt-28 my-4" src="logo.png" alt="Copolio Logo"
                         style={{width: '250px', height: '250px'}}/>
                    <SearchBar className="lg:w-2/5 w-full px-2" placeholder="Type your domain name ideas"
                               onSearch={search}/>
                    {searchResult.length > 0 && (<div className="m-10 relative overflow-x-auto shadow-md sm:rounded-lg">
                        <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                            <thead
                                className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                            <tr>
                                <th scope="col" className="px-6 py-3">Name</th>
                                {searchResult.length > 0 ? searchResult[0].info?.map((item, index) => (
                                    <th scope="col" className="px-6 py-3" key={index}>{item.suffix}</th>
                                )) : ""}
                            </tr>
                            </thead>
                            <tbody>
                            {searchResult.map((item, index) => (
                                <tr key={index}
                                    className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600">
                                    <td className="px-6 py-4">{item.domainName}</td>
                                    {item.info?.map((info, infoIndex) => (
                                        <td className="px-6 py-4" key={infoIndex}>
                                            <div
                                                dangerouslySetInnerHTML={{__html: (info.available ? info.price as string : "구매 불가")}}/>
                                        </td>
                                    ))}
                                </tr>
                            ))
                            }
                            </tbody>
                        </table>
                    </div>)}
                </div>
            </div>
        </>
    )
}
