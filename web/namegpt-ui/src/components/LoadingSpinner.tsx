import React from 'react';

export default function LoadingSpinner(props: {
    isLoading: boolean,
    label: string
}) {
    return props.isLoading && (
        <div
            className="fixed top-0 right-0 bottom-0 left-0 flex items-center justify-center bg-gray-900 bg-opacity-75 z-50">
            <div className="relative">
                <div className="absolute top-0 right-0 bottom-0 left-0 flex items-center justify-center">
                    <span className="text-white text-sm">{props.label}</span>
                </div>
            </div>
        </div>


    );
}
