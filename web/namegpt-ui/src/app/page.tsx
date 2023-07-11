"use client"

import SearchBar from "@/components/SearchBar";
import {useState} from "react";
import {
    Configuration,
    GithubComCopolioNamegptPkgDtoGenerateDomainNameResult,
    V1ApiFactory
} from "@/utils/clients/namegpt";
import LoadingSpinner from "@/components/LoadingSpinner";
import axios from "axios";

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

    const openDomainInGabia = async (postData: { new_domain: string; search_gubun?: "domain_index" }) => {
        const form = document.createElement('form');
        form.method = 'POST';
        form.action = "https://domain.gabia.com/regist/regist_step1.php";
        form.target = '_blank';

        for (const key in postData) {
            if (postData.hasOwnProperty(key)) {
                const input = document.createElement('input');
                input.type = 'hidden';
                input.name = key;
                // @ts-ignore
                input.value = postData[key];
                form.appendChild(input);
            }
        }

        document.body.appendChild(form);
        form.submit();
        document.body.removeChild(form);
    }
    // @ts-ignore
    return (
        <>
            <div className="flex items-center justify-center h-full w-full">
                <LoadingSpinner isLoading={isLoading} label={"도메인 생성중...."} />
                <div className="flex flex-col items-center w-full">
                    <img className="mx-auto mt-28 my-4" src="logo.png" alt="Copolio Logo"
                         style={{width: '250px', height: '250px'}}/>
                    <SearchBar className="max-w-screen-lg md:w-4/5 sm:w-4/5 w-full px-2" placeholder="원하는 도메인에 대한 설명을 입력하세요."
                               onSearch={search}/>
                    {searchResult.length > 0 && (<div className="m-10 relative lg:w-2/5 overflow-x-auto shadow-md sm:rounded-lg">
                        <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                            <thead
                                className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                            <tr>
                                <th scope="col" className="px-6 py-3"></th>
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
                                    <td className="px-6 py-4">
                                        <button onClick={() => openDomainInGabia({new_domain: item.domainName || ""})}>
                                        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                                            <path strokeLinecap="round" strokeLinejoin="round" d="M2.25 3h1.386c.51 0 .955.343 1.087.835l.383 1.437M7.5 14.25a3 3 0 00-3 3h15.75m-12.75-3h11.218c1.121-2.3 2.1-4.684 2.924-7.138a60.114 60.114 0 00-16.536-1.84M7.5 14.25L5.106 5.272M6 20.25a.75.75 0 11-1.5 0 .75.75 0 011.5 0zm12.75 0a.75.75 0 11-1.5 0 .75.75 0 011.5 0z" />
                                        </svg>
                                        </button>
                                    </td>
                                    <td className="px-6 py-4">
                                        {item.domainName}
                                    </td>
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
