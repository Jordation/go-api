"use client"
import { BarOptions } from './ChartConfigs'
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

ChartJS.register(
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
);

export default function GraphWithParams({query}: any) {
    const [Data, setData] = useState({})
    const handleClick = () => {
        fetch("http://localhost:8000/testing", {method: "POST", body: JSON.stringify({q: query})})}

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
            <Bar options={BarOptions} data={data} />
            <button onClick={handleClick}>click me</button>
            holding the graph</div>
        </>

    )
}