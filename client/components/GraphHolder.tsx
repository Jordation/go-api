"use client"
import {use, useState} from 'react';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import { Bar } from 'react-chartjs-2';
import { Props, ScriptProps } from 'next/script';

const options = {
    plugins: {
        title: {
            display: true,
            text: 'Chart.js Bar Chart - Stacked',
        },
    },
    responsive: true,
    interaction: {
        mode: 'index' as const,
        intersect: false,
    },
};

async function processQuery(query: object){
    let data = fetch("http://localhost:8000/testing", {method: "POST", body: JSON.stringify({q: query})})
    .then(res => res.json)
    .catch(error => console.error(error))
    return data
}

export default function GraphWithParams({query}: any) {

    //fetch("http://localhost:8000/testing", {method: "POST", body: JSON.stringify({q: query})})
    //.then(res => res.json)
    //.then(res => setTheData(res))

    let fetcheddata = use(processQuery(query))

    const [thedata, setTheData] = useState({})

    ChartJS.register(
        CategoryScale,
        LinearScale,
        BarElement,
        Title,
        Tooltip,
        Legend
    );

    const labels = ["aye", "two", "four"]
    const chartData = ["1", "2", "3"]
    const data = {
        labels,
        datasets: [
            {
                label: "label",
                data: chartData,
                backgroundColor: 'rgb(75, 192, 192)'
            }
        ]
    }
    return (
        <>
            <div>im the div
            <Bar options={options} data={data} />
            {JSON.stringify(fetcheddata)}
            {JSON.stringify(thedata)}
            holding the graph</div>
        </>

    )
}



// export default function GraphHolder({ props, params }){
    // let query = params.query
    // return (
        // <>
        {/*  */}
        {/* </> */}
    // )
// }