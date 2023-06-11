"use client"

import SearchBar from "@/components/SearchBar";
import {
    Configuration,
    GithubComCopolioNamegptInternalNamegptEntityDomainName,
    V1Api,
    V1ApiFactory
} from "@/utils/clients";
import {useState} from "react";

export default function Home() {
    const v1Api = V1ApiFactory(new Configuration({basePath: "http://localhost:8080/api"}));
    const [searchResult, setSearchResult] =
        useState<GithubComCopolioNamegptInternalNamegptEntityDomainName[]>([]);
    const search = (value: string) => {
        v1Api.v1SearchGet(value)
            .then(res => {
                setSearchResult(res.data)
            })
    }
    return (
        <>
            <div className="flex items-center justify-center h-full">
                <div className="flex flex-col items-center w-full">
                    <img className="mx-auto mt-28 my-4" src="logo.png" alt="Copolio Logo"
                         style={{width: '250px', height: '250px'}}/>
                    <SearchBar className="lg:w-2/5 w-full px-2" placeholder="Type your domain name ideas"
                               onSearch={search}/>
                    <table className="table-auto m-20 w-full text-sm text-left text-gray-500 dark:text-gray-400">
                        <tbody>
                        {searchResult.map((item, index) => (
                            <tr key={index} className="bg-white border-b
                            {/*dark:bg-gray-800 dark:border-gray-700*/}
                            ">
                                <td>{item.name}</td>
                            </tr>
                        ))
                        }
                        </tbody>
                    </table>
                </div>
            </div>
        </>
    )
}
