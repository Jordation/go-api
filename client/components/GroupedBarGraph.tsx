"use client"
import {use, useEffect, useState} from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
ChartJS.register(CategoryScale,LinearScale,BarElement,Title,Tooltip,Legend);
import { Bar } from 'react-chartjs-2';
import { useSearchParams } from 'next/navigation';


export default function GroupedBarGraph() {
    const SearchParams = useSearchParams();
    const [Data, setData] = useState({})

    useEffect(() => {
        console.log(SearchParams)
    }, [SearchParams]);

    return (
            <div className='graphArea'>
                {SearchParams && SearchParams}
            </div>
    )
}



// <button onClick={() => {
//     const params = Object.fromEntries(SearchParams.entries());
//     let newo = {filters:{}, filters_NOT:{}}
//     for(const [key, value] of Object.entries(params)) {
//         if (key.includes(".")) {
//             let frags = key.split(".")
//             newo[frags[0]][frags[1]] = value
//         } else {
//             newo[key] = value
//         }
//     }
//     console.log("fixed object", newo)
// }}>clickethme to show data</button>